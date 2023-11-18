//go:generate go run .

package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
)

type StructField struct {
	Type     string `yaml:"type"`
	Nullable bool   `yaml:"nullable"`
	Number   int    `yaml:"number"`
}

type StructInfo struct {
	Name     string                 `yaml:"name"`
	Package  string                 `yaml:"package"`
	Fields   map[string]StructField `yaml:"structure"`
	Primary  []string               `yaml:"primary"`
	Index    []string               `yaml:"index"`
}

type methodType struct {
	Script string
}

const repositoryTemplateCode = `
package {{.Package}}

import (
	"github.com/jinzhu/gorm"
	"github.com/game-core/gocrafter/domain/entity/user/{{.Package}}"
)

type {{.Name}}Repository interface {
{{range $methodName, $methodType := .Methods}}
	{{.Script}}
{{end}}
}
`

func generateRepository(yamlFilePath string, outputBaseDir string) error {
	structInfo, err := getStructInfo(yamlFilePath)
	if err != nil {
		return err
	}

	outputDir := filepath.Join(outputBaseDir, structInfo.Package)
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating output directory %s: %v", outputDir, err)
	}

	outputFileName := filepath.Join(outputDir, fmt.Sprintf("%s_repository.gen.go", structInfo.Package))
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("outputFileName file %s create error: %v", outputFileName, err)
	}
	defer outputFile.Close()

	methods := make(map[string]methodType)

	// FindByID
	if len(structInfo.Primary) > 0 {
		methods["FindByID"] = methodType{
			Script: fmt.Sprintf(`FindByID(ID int64) (*%s.%s, error)`, structInfo.Package, structInfo.Name),
		}
	}

	// FindByIndex
	for _, index := range structInfo.Index {
		indexFields := strings.Split(index, ",")
		params := make([]struct{ Name, Type string }, len(indexFields))

		var paramStrings []string

		for i, field := range indexFields {
			paramStrings = append(paramStrings, fmt.Sprintf("%s %s", field, structInfo.Fields[field].Type))
			params[i] = struct{ Name, Type string }{field, structInfo.Fields[field].Type}
		}

		methods[fmt.Sprintf("FindBy%s", strings.Join(indexFields, "And"))] = methodType{
			Script: fmt.Sprintf(
				`FindBy%s(%s) (*%s.%s, error)`,
				strings.Join(indexFields, "And"),
				strings.Join(paramStrings, ","),
				structInfo.Package,
				structInfo.Name,
			),
		}
	}

	// List
	methods["List"] = methodType{
		Script: fmt.Sprintf(
			`List(limit int64) (*%s.%ss, error)`,
			structInfo.Package,
			structInfo.Name,
		),
	}

	// Create
	methods["Create"] = methodType{
		Script: fmt.Sprintf(
			`Create(%s *%s.%s, tx *gorm.DB) (*%s.%s, error)`,
			structInfo.Package,
			structInfo.Package,
			structInfo.Name,
			structInfo.Package,
			structInfo.Name,
		),
	}

	// Update
	methods["Update"] = methodType{
		Script: fmt.Sprintf(
			`Update(%s *%s.%s, tx *gorm.DB) (*%s.%s, error)`,
			structInfo.Package,
			structInfo.Package,
			structInfo.Name,
			structInfo.Package,
			structInfo.Name,
		),
	}

	// Delete
	methods["Delete"] = methodType{
		Script: fmt.Sprintf(
			`Delete(%s *%s.%s, tx *gorm.DB) error`,
			structInfo.Package,
			structInfo.Package,
			structInfo.Name,
		),
	}

	repositoryTmpl, err := template.New("repositoryTemplate").Parse(repositoryTemplateCode)
	if err != nil {
		return fmt.Errorf("error parsing repository template: %v", err)
	}

	err = repositoryTmpl.ExecuteTemplate(outputFile, "repositoryTemplate", struct {
		Name     string
		Package  string
		Database string
		Methods  map[string]methodType
	}{
		Name:     structInfo.Name,
		Package:  structInfo.Package,
		Methods:  methods,
	})
	if err != nil {
		return fmt.Errorf("template error: %v", err)
	}

	fmt.Printf("Created %s Repository in %s\n", structInfo.Name, outputFileName)

	return nil
}

func getStructInfo(yamlFilePath string) (*StructInfo, error) {
	yamlData, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		return nil, fmt.Errorf("error reading YAML file %s: %v", yamlFilePath, err)
	}

	var structInfo StructInfo
	if err := yaml.Unmarshal(yamlData, &structInfo); err != nil {
		return nil, fmt.Errorf("error unmarshalling YAML in file %s: %v", yamlFilePath, err)
	}

	return &structInfo, nil
}

func main() {
	outputBaseDir := "../../../domain/repository/user"
	yamlFiles, err := filepath.Glob("../../../docs/entity/user/*.yaml")
	if err != nil {
		log.Fatalf("Error finding YAML files: %v", err)
	}

	for _, yamlFile := range yamlFiles {
		err := generateRepository(yamlFile, outputBaseDir)
		if err != nil {
			log.Printf("Error generating repository from YAML file %s: %v", yamlFile, err)
		}
	}
}
