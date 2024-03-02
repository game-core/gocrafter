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
)

type Di struct{}

func NewDi() *Di {
	return &Di{}
}

var diCode string

// generate 生成する
func (s *Di) generate() error {
	importCode = ""

	// service
	if err := s.create("service", "../../../../pkg/domain/model"); err != nil {
		return err
	}

	fmt.Printf(diCode)
	return nil
}

// create 作成する
func (s *Di) create(layer, path string) error {
	if err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if info.IsDir() {
			return nil
		}

		if strings.HasSuffix(info.Name(), fmt.Sprintf("_%s.go", layer)) {
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
func (s *Di) getStructName(filePath, field string) string {
	parts := strings.Split(strings.TrimSuffix(filepath.Base(filePath), fmt.Sprintf("_%s.go", field)), "_")
	if len(parts) > 0 {
		return fmt.Sprintf("%sService", changes.SnakeToCamel(strings.Join(parts[:len(parts)-1], "_")))
	}

	return ""
}

// createScript
func (s *Di) createScript(layer, filePath string, file *ast.File) string {
	var scripts []string
	structName := s.getStructName(filePath, changes.SnakeToUpperCamel(layer))

	for _, decl := range file.Decls {
		switch d := decl.(type) {
		case *ast.GenDecl:
			for _, spec := range d.Specs {
				if ts, ok := spec.(*ast.TypeSpec); ok {
					if ts.Name.Name == structName {
						if st, ok := ts.Type.(*ast.StructType); ok {
							// フィールドを表示
							switch layer {
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

// serviceを生成する
func (s *Di) serviceScript(structName string, fields []*ast.Field) string {
	var scripts []string

	for _, field := range fields {
		for _, field := range field.Names {
			fieldName := field.Name

			if strings.Contains(fieldName, "Service") {
				name := strings.Replace(fieldName, "Service", "", -1)
				importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("%sService github.com/game-core/gocrafter/pkg/domain/model/%s", name, name))
				scripts = append(scripts, fmt.Sprintf("Initialize%sService,", name))
			}

			if strings.Contains(fieldName, "Repository") {
				name := strings.Replace(fieldName, "Repository", "", -1)
				importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("%sDao github.com/game-core/gocrafter/pkg/infrastructure/mysql/common/%s", name, name))
				importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("%sDao github.com/game-core/gocrafter/pkg/infrastructure/mysql/master/%s", name, name))
				importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("%sDao github.com/game-core/gocrafter/pkg/infrastructure/mysql/user/%s", name, name))
				scripts = append(scripts, fmt.Sprintf("%sDao.New%sDao,", name, changes.CamelToUpperCamel(name)))
			}
		}
	}

	return fmt.Sprintf(
		`func Initialize%s() %s.%s {
			wire.Build(
				%s.New%s,
				database.NewDB,
				%s
			)
			return nil
		}`,
		changes.CamelToUpperCamel(structName),
		structName,
		changes.CamelToUpperCamel(structName),
		structName,
		changes.CamelToUpperCamel(structName),
		strings.Join(scripts, "\n"),
	)
}
