package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/game-core/gocrafter/internal/changes"
	"gopkg.in/yaml.v3"
)

type Handler struct{}

func NewHandler() *Handler {
	return &Handler{}
}

const handlerTemplate = `package {{.Package}}

import (
	"context"

	"github.com/game-core/gocrafter/api/game/presentation/server/{{.Package}}"
	accountUsecase "github.com/game-core/gocrafter/api/game/usecase/{{.Package}}"
	"github.com/game-core/gocrafter/internal/errors"
	"github.com/game-core/gocrafter/internal/tokens"
)

type {{.Name}}Handler interface {
	{{.Package}}.{{.Name}}Server
}

type {{.Package}}Handler struct {
	{{.Package}}.Unimplemented{{.Name}}Server
	{{.Package}}Usecase {{.Package}}Usecase.{{.Name}}Usecase
}

func New{{.Name}}Handler(
	{{.Package}}Usecase {{.Package}}Usecase.{{.Name}}Usecase,
) {{.Name}}Handler {
	return &{{.Package}}Handler{
		{{.Package}}Usecase: {{.Package}}Usecase,
	}
}

{{.Script}}
`

// generate 生成する
func (s *Handler) generate(file string, base string) error {
	importCode = ""

	yamlStruct, err := s.getYamlStruct(file)
	if err != nil {
		return err
	}

	outputDir := filepath.Join(base, yamlStruct.Package)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}
	fileName := s.getOutputFileName(outputDir, filepath.Base(file[:len(file)-len(filepath.Ext(file))]))

	switch {
	case strings.Contains(fileName, "_request"):
		return nil
	case strings.Contains(fileName, "_response"):
		return nil
	case strings.Contains(fileName, "_structure"):
		return nil
	case strings.Contains(fileName, "_enum"):
		return nil
	default:
		if err := s.createOutputFile(yamlStruct, fileName); err != nil {
			return err
		}
		return nil
	}
}

// getYamlStruct yaml構造体を取得する
func (s *Handler) getYamlStruct(file string) (*YamlStruct, error) {
	yamlData, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var yamlStruct YamlStruct
	if err := yaml.Unmarshal(yamlData, &yamlStruct); err != nil {
		return nil, err
	}

	return &yamlStruct, nil
}

// getOutputFileName ファイル名を取得する
func (s *Handler) getOutputFileName(dir, name string) string {
	return filepath.Join(dir, fmt.Sprintf("%s_handler.gen.go", changes.UpperCamelToSnake(name)))
}

// createOutputFile ファイルを作成する
func (s *Handler) createOutputFile(yamlStruct *YamlStruct, outputFileName string) error {
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}

	if err := s.createTemplate(yamlStruct, outputFile); err != nil {
		return err
	}

	return nil
}

// createTemplate テンプレートを作成する
func (s *Handler) createTemplate(yamlStruct *YamlStruct, outputFile *os.File) error {
	tmp, err := template.New("handlerTemplate").Parse(handlerTemplate)
	if err != nil {
		return err
	}

	if err := tmp.ExecuteTemplate(
		outputFile,
		"handlerTemplate",
		TemplateStruct{
			Name:    yamlStruct.Name,
			Package: yamlStruct.Package,
			Comment: yamlStruct.Comment,
			Script:  s.createScript(yamlStruct),
		},
	); err != nil {
		return err
	}

	return nil
}

// createScript スクリプトを作成する
func (s *Handler) createScript(yamlStruct *YamlStruct) string {
	var methodScripts []string

	for _, method := range s.getStructure(yamlStruct.Structures) {
		methodScripts = append(methodScripts, s.createMethod(yamlStruct, method))
	}

	return strings.Join(methodScripts, "\n")
}

// getStructure フィールド構造体を取得する
func (s *Handler) getStructure(structures map[string]Structure) []*Structure {
	var sortStructures []*Structure
	for _, value := range structures {
		sortStructures = append(
			sortStructures,
			&Structure{
				Name:     value.Name,
				Method:   value.Method,
				Auth:     value.Auth,
				Request:  value.Request,
				Response: value.Response,
				Number:   value.Number,
				Comment:  value.Comment,
			},
		)
	}

	sort.Slice(sortStructures, func(i, j int) bool {
		return sortStructures[i].Number < sortStructures[j].Number
	})

	return sortStructures
}

func (s *Handler) createMethod(yamlStruct *YamlStruct, structure *Structure) string {
	return fmt.Sprintf(
		`// %s %s
		func (s *%sHandler) %s(ctx context.Context, req *%s.%s) (*%s.%s, error) {%s
			res, err := s.%sUsecase.%s(ctx, req)
			if err != nil {
				return nil, errors.NewMethodError("s.%sUsecase.%s", err)
			}
		
			return res, nil
		}
		`,
		changes.SnakeToUpperCamel(structure.Name),
		structure.Comment,
		yamlStruct.Package,
		changes.SnakeToUpperCamel(structure.Name),
		yamlStruct.Package,
		structure.Request,
		yamlStruct.Package,
		structure.Response,
		s.createJwtMethod(structure),
		yamlStruct.Package,
		changes.SnakeToUpperCamel(structure.Name),
		yamlStruct.Package,
		changes.SnakeToUpperCamel(structure.Name),
	)
}

func (s *Handler) createJwtMethod(structure *Structure) string {
	if structure.Auth {
		return fmt.Sprintf(
			`if err := tokens.CheckJwtClaims(ctx, req.UserId); err != nil {
				return nil, errors.NewMethodError("internal.CheckJwtClaims", err)
			}`,
		)
	}

	return ""
}
