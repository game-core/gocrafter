//go:generate go run .

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
	"gopkg.in/yaml.v2"
)

type StructField struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Nullable bool   `yaml:"nullable"`
	Number   int    `yaml:"number"`
}

type StructInfo struct {
	Name     string                 `yaml:"name"`
	Package  string                 `yaml:"package"`
	Fields   map[string]StructField `yaml:"structure"`
	Script   string
}

const templateCode = `
package {{.Package}}

type {{.Name}} struct {
{{range $field := sortByNumber .Fields}}
	{{$field.Name}} {{$field.TypeWithPointer}} ` + "`json:\"{{$field.Json}}\"{{if eq $field.Name \"CreatedAt\"}} gorm:\"autoCreateTime\"{{else if eq $field.Name \"UpdatedAt\"}} gorm:\"autoUpdateTime\"{{end}}`" + `
{{end}}
}

{{.Script}}
`

func generateResponse(yamlFilePath string, outputBaseDir string) error {
	yamlData, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		return fmt.Errorf("error reading YAML file %s: %v", yamlFilePath, err)
	}

	var structInfo StructInfo
	if err = yaml.Unmarshal(yamlData, &structInfo); err != nil {
		return fmt.Errorf("error unmarshalling YAML in file %s: %v", yamlFilePath, err)
	}

	tmpl, err := template.New("structTemplate").Funcs(template.FuncMap{
		"sortByNumber": sortByNumber,
	}).Parse(templateCode)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	outputDir := filepath.Join(outputBaseDir, structInfo.Package)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating output directory %s: %v", outputDir, err)
	}

	outputFileName := filepath.Join(outputDir, fmt.Sprintf("%s_request.gen.go", structInfo.Name))
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("outputFileName file %s create error: %v", outputFileName, err)
	}
	defer outputFile.Close()

	var paramStrings []string
	var scriptStrings []string
	fieldsOrdered := make([]string, 0, len(structInfo.Fields))

	for fieldName := range structInfo.Fields {
		fieldsOrdered = append(fieldsOrdered, fieldName)

		param := strings.ToLower(fieldName[:1]) + fieldName[1:]
		paramStrings = append(paramStrings, fmt.Sprintf("%s %s", param, getTypeWithPointer(structInfo.Fields[fieldName])))
		scriptStrings = append(scriptStrings, fmt.Sprintf("%s: %s", fieldName, param))
	}

	script := fmt.Sprintf(`
	func %sResponse(%s) *%s {
		return &%s{%s}
	}
	`, structInfo.Name, strings.Join(paramStrings, ","), structInfo.Name, structInfo.Name, strings.Join(scriptStrings, ","))

	if err = tmpl.ExecuteTemplate(outputFile, "structTemplate", struct {
		Name        string
		Package     string
		Fields      map[string]StructField
		FieldsOrder []string
		Script      string
	}{
		Name:        structInfo.Name,
		Package:     structInfo.Package,
		Fields:      structInfo.Fields,
		FieldsOrder: fieldsOrdered,
		Script:      script,
	}); err != nil {
		return fmt.Errorf("template error: %v", err)
	}

	fmt.Printf("Created %s Response in %s\n", structInfo.Name, outputFileName)

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

	sort.Slice(sortedFields, func(i, j int) bool {
		return fields[sortedFields[i].Name].Number < fields[sortedFields[j].Name].Number
	})

	return sortedFields
}

func getTypeWithPointer(fieldInfo StructField) string {
	if fieldInfo.Nullable {
		return "*" + fieldInfo.Type
	}
	return fieldInfo.Type
}

func main() {
	userOutput := "../../api/presentation/response"
	userYamlFiles, err := filepath.Glob("../../docs/api/response/**/*.yaml")
	if err != nil {
		log.Fatalf("Error finding YAML files: %v", err)
	}

	for _, yamlFile := range userYamlFiles {
		if err := generateResponse(yamlFile, userOutput); err != nil {
			log.Printf("Error generating response from YAML file %s: %v", yamlFile, err)
		}
	}
}
