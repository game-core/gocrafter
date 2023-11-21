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
	Name    string                 `yaml:"name"`
	Package string                 `yaml:"package"`
	Fields  map[string]StructField `yaml:"structure"`
	Primary []string               `yaml:"primary"`
	Index   []string               `yaml:"index"`
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
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating output directory %s: %v", outputDir, err)
	}

	outputFileName := filepath.Join(outputDir, fmt.Sprintf("%s_repository.gen.go", structInfo.Package))
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("outputFileName file %s create error: %v", outputFileName, err)
	}
	defer outputFile.Close()

	repositoryTmpl, err := template.New("repositoryTemplate").Parse(repositoryTemplateCode)
	if err != nil {
		return fmt.Errorf("error parsing repository template: %v", err)
	}

	if err = repositoryTmpl.ExecuteTemplate(outputFile, "repositoryTemplate", struct {
		Name     string
		Package  string
		Database string
		Methods  map[string]methodType
	}{
		Name:    structInfo.Name,
		Package: structInfo.Package,
		Methods: generateMethods(structInfo),
	}); err != nil {
		return fmt.Errorf("template error: %v", err)
	}

	fmt.Printf("Created %s Repository in %s\n", structInfo.Name, outputFileName)

	return nil
}

func generateMethods(structInfo *StructInfo) map[string]methodType {
	methods := make(map[string]methodType)

	// FindByID
	methods["FindByID"] = methodType{
		Script: generateFindByID(structInfo),
	}

	// FindByIndex
	for _, index := range structInfo.Index {
		indexFields := strings.Split(index, ",")
		methods[fmt.Sprintf("FindBy%s", strings.Join(indexFields, "And"))] = methodType{
			Script: generateFindByIndex(structInfo, indexFields),
		}
	}

	// List
	methods["List"] = methodType{
		Script: generateList(structInfo),
	}

	// ListByIndex
	for _, index := range structInfo.Index {
		indexFields := strings.Split(index, ",")
		methods[fmt.Sprintf("ListBy%s", strings.Join(indexFields, "And"))] = methodType{
			Script: generateListByIndex(structInfo, indexFields),
		}
	}

	// Create
	methods["Create"] = methodType{
		Script: generateCreate(structInfo),
	}

	// Update
	methods["Update"] = methodType{
		Script: generateUpdate(structInfo),
	}

	// Delete
	methods["Delete"] = methodType{
		Script: generateDelete(structInfo),
	}

	return methods
}

func generateFindByID(structInfo *StructInfo) string {
	return fmt.Sprintf(`FindByID(ID int64) (*%s.%s, error)`, structInfo.Package, structInfo.Name)
}

func generateList(structInfo *StructInfo) string {
	return fmt.Sprintf(`List(limit int64) (*%s.%ss, error)`, structInfo.Package, structInfo.Name)
}

func generateCreate(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`Create(%s *%s.%s, tx *gorm.DB) (*%s.%s, error)`,
		structInfo.Package,
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
	)
}

func generateUpdate(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`Update(%s *%s.%s, tx *gorm.DB) (*%s.%s, error)`,
		structInfo.Package,
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
	)
}

func generateDelete(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`Delete(%s *%s.%s, tx *gorm.DB) error`,
		structInfo.Package,
		structInfo.Package,
		structInfo.Name,
	)
}

func generateFindByIndex(structInfo *StructInfo, indexFields []string) string {
	params := make([]struct{ Name, Type string }, len(indexFields))
	var paramStrings []string

	for i, field := range indexFields {
		paramStrings = append(paramStrings, fmt.Sprintf("%s %s", field, structInfo.Fields[field].Type))
		params[i] = struct{ Name, Type string }{field, structInfo.Fields[field].Type}
	}

	return fmt.Sprintf(
		`FindBy%s(%s) (*%s.%s, error)`,
		strings.Join(indexFields, "And"),
		strings.Join(paramStrings, ","),
		structInfo.Package,
		structInfo.Name,
	)
}

func generateListByIndex(structInfo *StructInfo, indexFields []string) string {
	params := make([]struct{ Name, Type string }, len(indexFields))
	var paramStrings []string

	for i, field := range indexFields {
		paramStrings = append(paramStrings, fmt.Sprintf("%s %s", field, structInfo.Fields[field].Type))
		params[i] = struct{ Name, Type string }{field, structInfo.Fields[field].Type}
	}

	return fmt.Sprintf(
		`ListBy%s(%s) (*%s.%ss, error)`,
		strings.Join(indexFields, "And"),
		strings.Join(paramStrings, ","),
		structInfo.Package,
		structInfo.Name,
	)
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
	outputBaseDir := "../../../../domain/repository/user"
	yamlFiles, err := filepath.Glob("../../../../docs/entity/user/*.yaml")
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
