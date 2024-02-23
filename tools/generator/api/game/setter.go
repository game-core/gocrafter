//go:generate go run .

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"

	"github.com/game-core/gocrafter/internal"
)

const templateCode = `
// Package {{.Package}} {{.Comment}}
package {{.Package}}

import (
	{{.Import}}
)

{{.Script}}
`

type Setter struct{}

func NewSetter() *Setter {
	return &Setter{}
}

// generate 生成する
func (s *Setter) generate(file string, base string) error {
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
		if err := s.createOutputFile(yamlStruct, fileName); err != nil {
			return err
		}
		return nil
	case strings.Contains(fileName, "_response"):
		if err := s.createOutputFile(yamlStruct, fileName); err != nil {
			return err
		}
		return nil
	case strings.Contains(fileName, "_structure"):
		if err := s.createOutputFile(yamlStruct, fileName); err != nil {
			return err
		}
		return nil
	case strings.Contains(fileName, "_enum"):
		return nil
	default:
		return nil
	}
}

// getYamlStruct yaml構造体を取得する
func (s *Setter) getYamlStruct(file string) (*YamlStruct, error) {
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
func (s *Setter) getOutputFileName(dir, name string) string {
	return filepath.Join(dir, fmt.Sprintf("%s_model.gen.go", internal.UpperCamelToSnake(name)))
}

// createOutputFile ファイルを作成する
func (s *Setter) createOutputFile(yamlStruct *YamlStruct, outputFileName string) error {
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
func (s *Setter) createTemplate(yamlStruct *YamlStruct, outputFile *os.File) error {
	tmp, err := template.New("templateCode").Parse(templateCode)
	if err != nil {
		return err
	}

	if err := tmp.ExecuteTemplate(
		outputFile,
		"templateCode",
		TemplateStruct{
			Name:    yamlStruct.Name,
			Package: yamlStruct.Package,
			Comment: yamlStruct.Comment,
			Script:  s.createScript(yamlStruct),
			Import:  importCode,
		},
	); err != nil {
		return err
	}

	return nil
}

// createScript スクリプトを作成する
func (s *Setter) createScript(yamlStruct *YamlStruct) string {
	var fieldScript []string
	var paramScript []string
	var returnScript []string

	for _, field := range s.getStructure(yamlStruct.Structures) {
		fieldScript = append(fieldScript, fmt.Sprintf("%s %s", internal.SnakeToUpperCamel(field.Name), s.getType(field)))
		paramScript = append(paramScript, fmt.Sprintf("%s %s", internal.SnakeToCamel(field.Name), s.getType(field)))
		returnScript = append(returnScript, fmt.Sprintf("%s: %s,", internal.SnakeToUpperCamel(field.Name), internal.SnakeToCamel(field.Name)))
	}

	return fmt.Sprintf(`%s`, s.createSetter(yamlStruct.Name, strings.Join(paramScript, ","), strings.Join(returnScript, "\n")))
}

// createSetter Setterを作成する
func (s *Setter) createSetter(name, paramScript, returnScript string) string {
	return fmt.Sprintf(
		`func Set%s(%s) *%s {
			return &%s{
				%s
			}
		}`,
		name,
		paramScript,
		name,
		name,
		returnScript,
	)
}

// getStructure フィールド構造体を取得する
func (s *Setter) getStructure(structures map[string]Structure) []*Structure {
	var sortStructures []*Structure
	for _, value := range structures {
		sortStructures = append(
			sortStructures,
			&Structure{
				Name:     value.Name,
				Type:     value.Type,
				Package:  value.Package,
				Nullable: value.Nullable,
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

// getType 型を取得する
func (s *Setter) getType(field *Structure) string {
	var result string

	switch field.Type {
	case "time":
		importCode = fmt.Sprintf("%s\n%s", importCode, "\"google.golang.org/protobuf/types/known/timestamppb\"")
		// timeの場合はポインタ固定にする
		return "*timestamppb.Timestamp"
	case "structure":
		if field.Package != "" {
			result = fmt.Sprintf("%s", internal.SnakeToUpperCamel(field.Name))
		} else {
			result = internal.SnakeToUpperCamel(field.Name)
		}
	case "structures":
		if field.Package != "" {
			result = fmt.Sprintf("%s", internal.SnakeToUpperCamel(internal.SingularToPlural(field.Name)))
		} else {
			result = internal.SnakeToUpperCamel(internal.SingularToPlural(field.Name))
		}
	case "enum":
		if field.Package != "" {
			result = fmt.Sprintf("%s", internal.SnakeToUpperCamel(field.Name))
		} else {
			result = internal.SnakeToUpperCamel(field.Name)
		}
	default:
		result = field.Type
	}

	if field.Nullable {
		result = fmt.Sprintf("*%s", result)
	}

	return result
}
