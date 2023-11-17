//go:generate go run .

package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"text/template"
)

type StructField struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Nullable bool   `yaml:"nullable"`
	Number   int    `yaml:"number"`
}

type StructInfo struct {
	Name    string                 `yaml:"name"`
	Package string                 `yaml:"package"`
	Fields  map[string]StructField `yaml:"structure"`
}

const templateCode = `
package {{.Package}}

type {{.Name}} struct {
{{range $field := sortByNumber .Fields}}
	{{$field.Name}} {{$field.TypeWithPointer}} ` + "`json:\"{{$field.Json}}\"{{if eq $field.Name \"CreatedAt\"}} gorm:\"autoCreateTime\"{{else if eq $field.Name \"UpdatedAt\"}} gorm:\"autoUpdateTime\"{{end}}`" + `
{{end}}
}
`

func generateRequest(yamlFilePath string, outputBaseDir string) error {
	structInfo, err := getStructInfo(yamlFilePath)
	if err != nil {
		return err
	}

	tmpl, err := template.New("structTemplate").Funcs(template.FuncMap{
		"sortByNumber": sortByNumber,
	}).Parse(templateCode)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	outputDir := filepath.Join(outputBaseDir, structInfo.Package)
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating output directory %s: %v", outputDir, err)
	}

	outputFileName := filepath.Join(outputDir, fmt.Sprintf("%s_request.gen.go", structInfo.Name))
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("outputFileName file %s create error: %v", outputFileName, err)
	}
	defer outputFile.Close()

	fieldsOrdered := make([]string, 0, len(structInfo.Fields))
	for fieldName := range structInfo.Fields {
		fieldsOrdered = append(fieldsOrdered, fieldName)
	}

	err = tmpl.ExecuteTemplate(outputFile, "structTemplate", struct {
		Name    string
		Package string
		Fields  map[string]StructField
	}{
		Name:    structInfo.Name,
		Package: structInfo.Package,
		Fields:  structInfo.Fields,
	})
	if err != nil {
		return fmt.Errorf("template error: %v", err)
	}

	fmt.Printf("Created %s Request in %s\n", structInfo.Name, outputFileName)

	return nil
}

func sortByNumber(fields map[string]StructField) []struct {
	Name            string
	FieldInfo       StructField
	TypeWithPointer string
	Json            string
} {
	var sortedFields []struct {
		Name            string
		FieldInfo       StructField
		TypeWithPointer string
		Json            string
	}

	for name, fieldInfo := range fields {
		sortedFields = append(sortedFields, struct {
			Name            string
			FieldInfo       StructField
			TypeWithPointer string
			Json            string
		}{
			Name:            name,
			FieldInfo:       fieldInfo,
			TypeWithPointer: getTypeWithPointer(fieldInfo),
			Json:            fieldInfo.Name,
		})
	}

	sort.SliceStable(sortedFields, func(i, j int) bool {
		return fields[sortedFields[i].Name].Number < fields[sortedFields[j].Name].Number
	})

	return sortedFields
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

func getTypeWithPointer(fieldInfo StructField) string {
	if fieldInfo.Nullable {
		return "*" + fieldInfo.Type
	}

	return fieldInfo.Type
}

func main() {
	userOutput := "../../../api/presentation/request"
	userYamlFiles, err := filepath.Glob("../../../docs/api/request/**/*.yaml")
	if err != nil {
		log.Fatalf("Error finding YAML files: %v", err)
	}

	for _, yamlFile := range userYamlFiles {
		err := generateRequest(yamlFile, userOutput)
		if err != nil {
			log.Printf("Error generating request from YAML file %s: %v", yamlFile, err)
		}
	}
}
