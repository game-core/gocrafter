package main

import (
	"fmt"
	"github.com/game-core/gocrafter/internal/changes"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type Di struct{}

func NewDi() *Di {
	return &Di{}
}

var diCode string

const diTemplate = `//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/game-core/gocrafter/configs/database"

	{{.Import}}
)

{{.Script}}
`

// generate 生成する
func (s *Di) generate() error {
	importCode = ""

	// interceptor
	if err := s.create("interceptor", "../../../../api/game/presentation/interceptor"); err != nil {
		return err
	}

	// handler
	if err := s.create("handler", "../../../../api/game/presentation/handler"); err != nil {
		return err
	}

	// usecase
	if err := s.create("usecase", "../../../../api/game/usecase"); err != nil {
		return err
	}

	// service
	if err := s.create("service", "../../../../pkg/domain/model"); err != nil {
		return err
	}

	if err := s.createTemplate(); err != nil {
		return err
	}

	return nil
}

// create 作成する
func (s *Di) create(layer, path string) error {
	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if strings.Contains(info.Name(), fmt.Sprintf("_%s", layer)) {
			if err := s.parseFile(layer, path); err != nil {
				return err
			}
		}

		return nil
	}); err != nil {
		return err
	}

	return nil
}

// createTemplate テンプレートを作成する
func (s *Di) createTemplate() error {
	outputFile, err := os.Create("../../../../api/game/di/wire.go")
	if err != nil {
		return err
	}

	tmp, err := template.New("diTemplate").Parse(diTemplate)
	if err != nil {
		return err
	}

	if err := tmp.ExecuteTemplate(
		outputFile,
		"diTemplate",
		TemplateStruct{
			Import: importCode,
			Script: diCode,
		},
	); err != nil {
		return err
	}

	return nil
}

// parseFile ファイルを解析する
func (s *Di) parseFile(layer, filePath string) error {
	file, err := parser.ParseFile(token.NewFileSet(), filePath, nil, parser.AllErrors)
	if err != nil {
		return err
	}

	diCode = fmt.Sprintf("%s\n%s", diCode, s.createScript(layer, filePath, file))

	return nil
}

// getStructName ファイル名から構造体名文字列を取得
func (s *Di) getStructName(layer, filePath string) string {
	parts := strings.Split(strings.TrimSuffix(filepath.Base(filePath), fmt.Sprintf("_%s.go", layer)), "_")
	if len(parts) > 0 {
		return fmt.Sprintf("%s%s", changes.SnakeToCamel(strings.Join(parts[:len(parts)-1], "_")), layer)
	}

	return ""
}

// createScript
func (s *Di) createScript(layer, filePath string, file *ast.File) string {
	var scripts []string
	structName := s.getStructName(changes.SnakeToUpperCamel(layer), filePath)

	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				if ts, ok := spec.(*ast.TypeSpec); ok {
					if ts.Name.Name == structName {
						if st, ok := ts.Type.(*ast.StructType); ok {
							switch layer {
							case "interceptor":
								scripts = append(scripts, s.interceptorScript(structName, st.Fields.List))
							case "handler":
								scripts = append(scripts, s.handlerScript(structName, st.Fields.List))
							case "usecase":
								scripts = append(scripts, s.usecaseScript(structName, st.Fields.List))
							case "service":
								scripts = append(scripts, s.serviceScript(structName, st.Fields.List))
							default:
							}
						}
					}
				}
			}
		}
	}

	return strings.Join(scripts, "\n")
}

// interceptorScript interceptorを生成する
func (s *Di) interceptorScript(structName string, fields []*ast.Field) string {
	importCode = fmt.Sprintf(
		"%s\n%s",
		importCode,
		fmt.Sprintf("%s \"github.com/game-core/gocrafter/api/game/presentation/interceptor/%s\"", structName, strings.Replace(structName, "Interceptor", "", -1)),
	)
	var scripts []string

	for _, field := range fields {
		for _, field := range field.Names {
			fieldName := field.Name

			if strings.Contains(fieldName, "Service") {
				name := strings.Replace(fieldName, "Service", "", -1)
				importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("%sService \"github.com/game-core/gocrafter/pkg/domain/model/%s\"", name, name))
				scripts = append(scripts, fmt.Sprintf("Initialize%sService,", changes.CamelToUpperCamel(name)))
			}
		}
	}

	return fmt.Sprintf(
		`func Initialize%s() %s.%s {
			wire.Build(
				%s.New%s,
				%s
			)
			return nil
		}
		`,
		changes.CamelToUpperCamel(structName),
		structName,
		changes.CamelToUpperCamel(structName),
		structName,
		changes.CamelToUpperCamel(structName),
		strings.Join(scripts, "\n"),
	)
}

// handlerScript handlerを生成する
func (s *Di) handlerScript(structName string, fields []*ast.Field) string {
	importCode = fmt.Sprintf(
		"%s\n%s",
		importCode,
		fmt.Sprintf("%s \"github.com/game-core/gocrafter/api/game/presentation/handler/%s\"", structName, strings.Replace(structName, "Handler", "", -1)),
	)
	var scripts []string

	for _, field := range fields {
		for _, field := range field.Names {
			fieldName := field.Name

			if strings.Contains(fieldName, "Usecase") {
				name := strings.Replace(fieldName, "Usecase", "", -1)
				importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("%sUsecase \"github.com/game-core/gocrafter/api/game/usecase/%s\"", name, name))
				scripts = append(scripts, fmt.Sprintf("Initialize%sUsecase,", changes.CamelToUpperCamel(name)))
			}
		}
	}

	return fmt.Sprintf(
		`func Initialize%s() %s.%s {
			wire.Build(
				%s.New%s,
				%s
			)
			return nil
		}
		`,
		changes.CamelToUpperCamel(structName),
		structName,
		changes.CamelToUpperCamel(structName),
		structName,
		changes.CamelToUpperCamel(structName),
		strings.Join(scripts, "\n"),
	)
}

// usecaseScript usecaseを生成する
func (s *Di) usecaseScript(structName string, fields []*ast.Field) string {
	var scripts []string

	for _, field := range fields {
		for _, field := range field.Names {
			fieldName := field.Name

			if strings.Contains(fieldName, "Service") {
				name := strings.Replace(fieldName, "Service", "", -1)
				importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("%sService \"github.com/game-core/gocrafter/pkg/domain/model/%s\"", name, name))
				scripts = append(scripts, fmt.Sprintf("Initialize%sService,", changes.CamelToUpperCamel(name)))
			}
		}
	}

	return fmt.Sprintf(
		`func Initialize%s() %s.%s {
			wire.Build(
				%s.New%s,
				%s
			)
			return nil
		}
		`,
		changes.CamelToUpperCamel(structName),
		structName,
		changes.CamelToUpperCamel(structName),
		structName,
		changes.CamelToUpperCamel(structName),
		strings.Join(scripts, "\n"),
	)
}

// serviceScript serviceを生成する
func (s *Di) serviceScript(structName string, fields []*ast.Field) string {
	var scripts []string

	for _, field := range fields {
		for _, field := range field.Names {
			fieldName := field.Name

			if strings.Contains(fieldName, "Service") {
				name := strings.Replace(fieldName, "Service", "", -1)
				importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("%sService \"github.com/game-core/gocrafter/pkg/domain/model/%s\"", name, name))
				scripts = append(scripts, fmt.Sprintf("Initialize%sService,", changes.CamelToUpperCamel(name)))
			}

			if strings.Contains(fieldName, "Repository") {
				name := strings.Replace(fieldName, "Repository", "", -1)
				importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("%sDao \"github.com/game-core/gocrafter/pkg/infrastructure/mysql/%s/%s\"", name, s.getDaoDir(name), name))
				scripts = append(scripts, fmt.Sprintf("%sDao.New%sDao,", name, changes.CamelToUpperCamel(name)))
			}
		}
	}

	return fmt.Sprintf(
		`func Initialize%s() %s.%s {
			wire.Build(
				database.NewDB,
				%s.New%s,
				%s
			)
			return nil
		}
		`,
		changes.CamelToUpperCamel(structName),
		structName,
		changes.CamelToUpperCamel(structName),
		structName,
		changes.CamelToUpperCamel(structName),
		strings.Join(scripts, "\n"),
	)
}

// getDaoDir daoのディレクトリを取得する
func (s *Di) getDaoDir(name string) string {
	if fileExists("../../../../pkg/infrastructure/mysql/common", fmt.Sprintf("%s_dao.gen.go", changes.CamelToSnake(name))) {
		return "common"
	}

	if fileExists("../../../../pkg/infrastructure/mysql/master", fmt.Sprintf("%s_dao.gen.go", changes.CamelToSnake(name))) {
		return "master"
	}

	if fileExists("../../../../pkg/infrastructure/mysql/user", fmt.Sprintf("%s_dao.gen.go", changes.CamelToSnake(name))) {
		return "user"
	}

	return ""
}

// fileExists ファイルの存在を確認する
func fileExists(root, target string) bool {
	var exists bool

	if err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if info.Name() == target {
			exists = true
			return filepath.SkipDir
		}

		return nil
	}); err != nil {
		fmt.Println(err)
	}

	return exists
}
