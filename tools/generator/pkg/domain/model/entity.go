//go:generate go run .

package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"github.com/game-core/gocrafter/internal"
)

const templateCode = `
// Package {{.Package}} {{.Comment}}
package {{.Package}}

import (
	{{.Import}}
)

type {{.PluralName}} []*{{.Name}}

{{.Script}}
`

var importCode = ""

type YamlStruct struct {
	Name       string               `yaml:"name"`
	Package    string               `yaml:"package"`
	Comment    string               `yaml:"comment"`
	Structures map[string]Structure `yaml:"structure"`
}

type Structure struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Package  string `yaml:"package"`
	Nullable bool   `yaml:"nullable"`
	Number   int    `yaml:"number"`
	Comment  string `yaml:"comment"`
}

type TemplateStruct struct {
	Package    string
	Name       string
	PluralName string
	Comment    string
	Script     string
	Import     string
}

// generate 生成する
func generate(file string, base string) error {
	yamlStruct, err := getYamlStruct(file)
	if err != nil {
		return err
	}

	outputDir := filepath.Join(base, strings.Replace(filepath.Dir(file), "/../../docs/yaml", "", -1))
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	if err := createOutputFile(yamlStruct, getOutputFileName(outputDir, filepath.Base(file[:len(file)-len(filepath.Ext(file))]))); err != nil {
		return err
	}

	return nil
}

// getYamlStruct yaml構造体を取得する
func getYamlStruct(file string) (*YamlStruct, error) {
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
func getOutputFileName(dir, name string) string {
	return filepath.Join(dir, fmt.Sprintf("%s.gen.go", internal.UpperCamelToSnake(name)))
}

// createOutputFile ファイルを作成する
func createOutputFile(yamlStruct *YamlStruct, outputFileName string) error {
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}

	if err := createTemplate(yamlStruct, outputFile); err != nil {
		return err
	}

	return nil
}

// createTemplate テンプレートを作成する
func createTemplate(yamlStruct *YamlStruct, outputFile *os.File) error {
	tmp, err := template.New("templateCode").Parse(templateCode)
	if err != nil {
		return err
	}

	if err := tmp.ExecuteTemplate(
		outputFile,
		"templateCode",
		TemplateStruct{
			Name:       yamlStruct.Name,
			Package:    yamlStruct.Package,
			PluralName: internal.SingularToPlural(yamlStruct.Name),
			Comment:    yamlStruct.Comment,
			Script:     createScript(yamlStruct),
			Import:     importCode,
		},
	); err != nil {
		return err
	}

	return nil
}

// createScript スクリプトを作成する
func createScript(yamlStruct *YamlStruct) string {
	var fieldScript []string
	var paramScript []string
	var returnScript []string

	for _, field := range getStructure(yamlStruct.Structures) {
		fieldScript = append(fieldScript, fmt.Sprintf("%s %s", internal.SnakeToUpperCamel(field.Name), getType(field)))
		paramScript = append(paramScript, fmt.Sprintf("%s %s", internal.SnakeToCamel(field.Name), getType(field)))
		returnScript = append(returnScript, fmt.Sprintf("%s: %s,", internal.SnakeToUpperCamel(field.Name), internal.SnakeToCamel(field.Name)))
	}

	return fmt.Sprintf(
		`%s

		%s

		%s`,
		createStruct(yamlStruct.Name, strings.Join(fieldScript, "\n")),
		createNew(yamlStruct.Name, internal.SingularToPlural(yamlStruct.Name)),
		createSetter(yamlStruct.Name, strings.Join(paramScript, ","), strings.Join(returnScript, "\n")),
	)
}

// createStruct Structを作成する
func createStruct(name string, fieldScript string) string {
	return fmt.Sprintf(
		`type %s struct {
			%s
		}`,
		name,
		fieldScript,
	)
}

// createNew Newを作成する
func createNew(name, pluralName string) string {
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

// createSetter Setterを作成する
func createSetter(name, paramScript, returnScript string) string {
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
func getStructure(structures map[string]Structure) []*Structure {
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

// getType 型を取得する
func getType(field *Structure) string {
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

func main() {
	yamls := "../../../../../docs/yaml/pkg/domain/model"
	base := "../../../../../pkg/domain/model"

	if err := filepath.Walk(yamls, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("error accessing path %s: %v", path, err)
			return nil
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".yaml") {
			if err := generate(path, base); err != nil {
				log.Printf("failed to generate: %s", err)
			}
		}

		return nil
	}); err != nil {
		log.Fatalf("failed to filepath.Walk: %s", err)
	}
}
