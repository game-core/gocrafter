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
	"strings"
	"text/template"
	"time"

	"github.com/game-core/gocrafter/config/transform"
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
	Primary []string               `yaml:"primary"`
	Index   []string               `yaml:"index"`
}

const templateCode = `
CREATE TABLE {{.Table}} (
{{range $field := sortByNumber .Fields}}
	{{$field.Column}}{{$field.Type}}{{$field.TypeWithPointer}}{{$field.Config}},
{{end}}
	{{.PrimaryKey}},
	{{.IndexKey}}
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
`

func generateEntity(yamlFilePath string, outputDir string) error {
	structInfo, err := getStructInfo(yamlFilePath)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating output directory %s: %v", outputDir, err)
	}

	fileName := fmt.Sprintf("%s_%s.sql", time.Now().Format("20060102"), structInfo.Package)
	if fileExistsWithDifferentDateTime(outputDir, fileName) {
		return nil
	}

	outputFileName := filepath.Join(outputDir, fileName)
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("outputFileName file %s create error: %v", outputFileName, err)
	}

	var primaryStrings []string
	for _, primary := range structInfo.Primary {
		for field := range structInfo.Fields {
			if field == primary {
				primaryStrings = append(primaryStrings, fmt.Sprintf("PRIMARY KEY(%s)", structInfo.Fields[field].Name))
			}
		}
	}

	var indexStrings []string
	for _, index := range structInfo.Index {
		for field := range structInfo.Fields {
			if field == index {
				indexStrings = append(indexStrings, fmt.Sprintf("INDEX(%s)", structInfo.Fields[field].Name))
			}
		}
	}

	if err := generateTemplate(structInfo, outputFile, primaryStrings, indexStrings); err != nil {
		return fmt.Errorf("faild to generateTemplate: %v", err)
	}

	fmt.Printf("Created %s SQL in %s\n", structInfo.Name, outputFileName)

	return nil
}

func generateTemplate(structInfo *StructInfo, outputFile *os.File, primaryStrings, indexStrings []string) error {
	tmpl, err := template.New("structTemplate").Funcs(template.FuncMap{
		"sortByNumber": sortByNumber,
	}).Parse(templateCode)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	if err := tmpl.ExecuteTemplate(outputFile, "structTemplate", struct {
		Table      string
		Fields     map[string]StructField
		PrimaryKey string
		IndexKey   string
	}{
		Table:      transform.CamelToSnake(structInfo.Package),
		Fields:     structInfo.Fields,
		PrimaryKey: strings.Join(primaryStrings, ","),
		IndexKey:   strings.Join(indexStrings, ","),
	}); err != nil {
		return fmt.Errorf("template error: %v", err)
	}

	return nil
}

func sortByNumber(fields map[string]StructField) []struct {
	Name            string
	FieldInfo       StructField
	Column          string
	Type            string
	TypeWithPointer string
	Config          string
} {
	var sortedFields []struct {
		Name            string
		FieldInfo       StructField
		Column          string
		Type            string
		TypeWithPointer string
		Config          string
	}

	for name, fieldInfo := range fields {
		sortedFields = append(sortedFields, struct {
			Name            string
			FieldInfo       StructField
			Column          string
			Type            string
			TypeWithPointer string
			Config          string
		}{
			Name:            name,
			FieldInfo:       fieldInfo,
			Column:          fieldInfo.Name,
			Type:            getType(fieldInfo),
			TypeWithPointer: getTypeWithPointer(fieldInfo),
			Config:          getConfig(fieldInfo),
		})
	}

	sort.Slice(sortedFields, func(i, j int) bool {
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

func getType(fieldInfo StructField) string {
	switch fieldInfo.Type {
	case "string":
		return " VARCHAR(255)"
	case "int64":
		return " BIGINT"
	case "int":
		return " INT"
	case "bool":
		return " TINYINT"
	case "time.Time":
		return " TIMESTAMP"
	default:
		return ""
	}
}

func getTypeWithPointer(fieldInfo StructField) string {
	if fieldInfo.Nullable {
		return " DEFAULT NULL"
	}

	return " NOT NULL"
}

func getConfig(fieldInfo StructField) string {
	if fieldInfo.Name == "id" {
		return " AUTO_INCREMENT"
	}

	return ""
}

func getPrimary(fieldInfo StructField) string {
	if fieldInfo.Name == "id" {
		return " PRIMARY KEY (id),"
	}

	return ""
}

func fileExistsWithDifferentDateTime(directory, fileName string) bool {
	fileList, err := filepath.Glob(filepath.Join(directory, "*"))
	if err != nil {
		return false
	}

	targetFileName := extractFileName(fileName)

	for _, existingFile := range fileList {
		existingFileName := extractFileName(existingFile)
		if existingFileName == targetFileName {
			return true
		}
	}

	return false
}

func extractFileName(filePath string) string {
	elements := strings.Split(filePath, "/")
	fileName := elements[len(elements)-1]
	parts := strings.Split(fileName, "_")
	if len(parts) > 1 {
		fileName = parts[1]
	}

	return fileName
}

func main() {
	outputDir := "../../../../docs/sql/user"
	yamlFiles, err := filepath.Glob("../../../../docs/entity/user/*.yaml")
	if err != nil {
		log.Fatalf("Error finding YAML files: %v", err)
	}

	for _, yamlFile := range yamlFiles {
		err := generateEntity(yamlFile, outputDir)
		if err != nil {
			log.Printf("Error generating entity from YAML file %s: %v", yamlFile, err)
		}
	}
}
