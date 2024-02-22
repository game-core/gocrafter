//go:generate go run .

package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/game-core/gocrafter/internal"
)

const tableTemplate = `
// Package {{.Package}} {{.Comment}}
package {{.Package}}

import (
	{{.Import}}
)

type {{.PluralName}} []*{{.Name}}

{{.Script}}
`

// generateTable 生成する
func generateTable(file string, base string) error {
	importCode = ""

	yamlStruct, err := getTableYamlStruct(file)
	if err != nil {
		return err
	}

	outputDir := filepath.Join(base, strings.Replace(filepath.Dir(file), "/../../docs/yaml", "", -1))
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	if err := createTableOutputFile(yamlStruct, getTableOutputFileName(outputDir, filepath.Base(file[:len(file)-len(filepath.Ext(file))]))); err != nil {
		return err
	}

	return nil
}

// getTableYamlStruct yaml構造体を取得する
func getTableYamlStruct(file string) (*YamlStruct, error) {
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

// getTableOutputFileName ファイル名を取得する
func getTableOutputFileName(dir, name string) string {
	return filepath.Join(dir, fmt.Sprintf("%s.gen.go", internal.UpperCamelToSnake(name)))
}

// createTableOutputFile ファイルを作成する
func createTableOutputFile(yamlStruct *YamlStruct, outputFileName string) error {
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}

	if err := createTableTemplate(yamlStruct, outputFile); err != nil {
		return err
	}

	return nil
}

// createTableTemplate テンプレートを作成する
func createTableTemplate(yamlStruct *YamlStruct, outputFile *os.File) error {
	tmp, err := template.New("tableTemplate").Parse(tableTemplate)
	if err != nil {
		return err
	}

	if err := tmp.ExecuteTemplate(
		outputFile,
		"tableCode",
		TemplateStruct{
			Name:       yamlStruct.Name,
			Package:    yamlStruct.Package,
			PluralName: internal.SingularToPlural(yamlStruct.Name),
			Comment:    yamlStruct.Comment,
			Script:     createTableScript(yamlStruct),
			Import:     importCode,
		},
	); err != nil {
		return err
	}

	return nil
}

// createTableScript スクリプトを作成する
func createTableScript(yamlStruct *YamlStruct) string {
	var fieldScript []string
	var paramScript []string
	var returnScript []string

	for _, field := range getTableStructure(yamlStruct.Structures) {
		fieldScript = append(fieldScript, fmt.Sprintf("%s %s", internal.SnakeToUpperCamel(field.Name), getTableType(field)))
		paramScript = append(paramScript, fmt.Sprintf("%s %s", internal.SnakeToCamel(field.Name), getTableType(field)))
		returnScript = append(returnScript, fmt.Sprintf("%s: %s,", internal.SnakeToUpperCamel(field.Name), internal.SnakeToCamel(field.Name)))
	}

	return fmt.Sprintf(
		`%s

		%s

		%s

		%s`,
		createTableStruct(yamlStruct.Name, strings.Join(fieldScript, "\n")),
		createTableNew(yamlStruct.Name, internal.SingularToPlural(yamlStruct.Name)),
		createTableSetter(yamlStruct.Name, strings.Join(paramScript, ","), strings.Join(returnScript, "\n")),
		createTableNameScript(yamlStruct.Name),
	)
}

// createStruct Structを作成する
func createTableStruct(name string, fieldScript string) string {
	return fmt.Sprintf(
		`type %s struct {
			%s
		}`,
		name,
		fieldScript,
	)
}

// createNew Newを作成する
func createTableNew(name, pluralName string) string {
	return fmt.Sprintf(
		`func New%s() *%s {
			return &%s{}
		}

		func New%s() %s {
			return %s{}
		}`,
		name,
		name,
		name,
		pluralName,
		pluralName,
		pluralName,
	)
}

// createTableSetter Setterを作成する
func createTableSetter(name, paramScript, returnScript string) string {
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

// createTableNameScript TableNameを作成する
func createTableNameScript(name string) string {
	return fmt.Sprintf(
		`func (t *%s) TableName() string {
			return "%s"
		}`,
		name,
		internal.UpperCamelToSnake(name),
	)
}

// getTableStructure フィールド構造体を取得する
func getTableStructure(structures map[string]Structure) []*Structure {
	var sortStructures []*Structure
	for key, value := range structures {
		sortStructures = append(
			sortStructures,
			&Structure{
				Name:     key,
				Type:     value.Type,
				Package:  value.Package,
				Nullable: value.Nullable,
				Number:   value.Number,
				Comment:  value.Comment,
			},
		)
	}

	sort.SliceStable(sortStructures, func(i, j int) bool {
		return structures[sortStructures[i].Name].Number < structures[sortStructures[j].Name].Number
	})

	return sortStructures
}

// getTableType 型を取得する
func getTableType(field *Structure) string {
	var result string

	switch field.Type {
	case "time":
		importCode = fmt.Sprintf("%s\n%s", importCode, "\"time\"")
		result = "time.Time"
	case "structure":
		if field.Package != "" {
			importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("\"github.com/game-core/gocrafter/pkg/domain/model/%s\"", field.Package))
			result = fmt.Sprintf("%s.%s", internal.SnakeToCamel(field.Name), internal.SnakeToUpperCamel(field.Name))
		} else {
			result = internal.SnakeToUpperCamel(field.Name)
		}
	case "structures":
		if field.Package != "" {
			importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("\"github.com/game-core/gocrafter/pkg/domain/model/%s\"", field.Package))
			result = fmt.Sprintf("%s.%s", internal.SnakeToCamel(field.Name), internal.SnakeToUpperCamel(internal.SingularToPlural(field.Name)))
		} else {
			result = internal.SnakeToUpperCamel(internal.SingularToPlural(field.Name))
		}
	case "enum":
		if field.Package != "" {
			importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("\"github.com/game-core/gocrafter/pkg/domain/model/%s\"", field.Package))
			result = fmt.Sprintf("%s.%s", internal.SnakeToCamel(field.Name), internal.SnakeToUpperCamel(field.Name))
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
