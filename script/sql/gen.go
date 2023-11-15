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
	"time"
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
	Table    string                 `yaml:"table"`
	Database string                 `yaml:"database"`
	Package  string                 `yaml:"package"`
	Fields   map[string]StructField `yaml:"structure"`
	Primary  []string               `yaml:"primary"`
	Index    []string               `yaml:"index"`
}

const templateCode = 
`CREATE TABLE {{.Table}} (
{{range $field := sortByNumber .Fields}}
	{{$field.Column}}{{$field.Type}}{{$field.TypeWithPointer}}{{$field.Config}},
{{end}}
	{{.PrimaryKey}},
	{{.IndexKey}}
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;
`

func generateEntity(yamlFilePath string, outputBaseDir string) error {
	yamlData, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		return fmt.Errorf("error reading YAML file %s: %v", yamlFilePath, err)
	}

	var structInfo StructInfo
	if err := yaml.Unmarshal(yamlData, &structInfo); err != nil {
		return fmt.Errorf("error unmarshalling YAML in file %s: %v", yamlFilePath, err)
	}

	tmpl, err := template.New("structTemplate").Funcs(template.FuncMap{
		"sortByNumber": sortByNumber,
	}).Parse(templateCode)
	if err != nil {
		return fmt.Errorf("error parsing template: %v", err)
	}

	outputDir := fmt.Sprintf("%s/%s", outputBaseDir, structInfo.Database)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating output directory %s: %v", outputDir, err)
	}

	fileName := fmt.Sprintf("%s_%s.sql", time.Now().Format("20060102"), structInfo.Package)
	outputFileName := filepath.Join(outputDir, fileName)
	if fileExistsWithDifferentDateTime(outputDir, fileName) {
		fmt.Printf("Skipped %s Entity because the file already exists: %s\n", structInfo.Name, outputFileName)
		return nil
	}

	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("outputFileName file %s create error: %v", outputFileName, err)
	}
	defer outputFile.Close()

	fieldsOrdered := make([]string, 0, len(structInfo.Fields))
	for fieldName := range structInfo.Fields {
		fieldsOrdered = append(fieldsOrdered, fieldName)
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

	if err = tmpl.ExecuteTemplate(outputFile, "structTemplate", struct {
		Table       string
		Fields      map[string]StructField
		PrimaryKey  string
		IndexKey    string
	}{
		Table:       structInfo.Table,
		Fields:      structInfo.Fields,
		PrimaryKey:  strings.Join(primaryStrings, ","),
		IndexKey:    strings.Join(indexStrings, ","),
	}); err != nil {
		return fmt.Errorf("template error: %v", err)
	}

	fmt.Printf("Created %s Entity in %s\n", structInfo.Name, outputFileName)

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

func getType(fieldInfo StructField)  string {
	if fieldInfo.Type == "string" {
		return " VARCHAR(255)"
	} else if fieldInfo.Type == "int64" {
		return " BIGINT"
	} else if fieldInfo.Type == "int" {
		return " INT"
	} else if fieldInfo.Type == "time.Time" {
		return " TIMESTAMP"
	} 

	return ""
}

func getTypeWithPointer(fieldInfo StructField) string {
	if fieldInfo.Nullable {
		return " DEFAULT NULL"
	}

	return " NOT NULL"
}

func getConfig(fieldInfo StructField)  string {
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
	// ディレクトリ内のファイル一覧を取得
	fileList, err := filepath.Glob(filepath.Join(directory, "*"))
	if err != nil {
		fmt.Println("Error:", err)
		return false
	}

	// 指定されたファイル名から日時情報を取り除く
	targetFileName := extractFileName(fileName)

	// ディレクトリ内のファイルを走査して同じファイル名が存在するか確認
	for _, existingFile := range fileList {
		existingFileName := extractFileName(existingFile)

		// 同じファイル名が存在する場合、日時情報が異なっていてもtrueを返す
		if existingFileName == targetFileName {
			return true
		}
	}

	// 同じファイル名が見つからない場合はfalseを返す
	return false
}

// ファイル名から日時情報を取り除く関数
func extractFileName(filePath string) string {
	// ファイルパスをスラッシュで分割して最後の要素を取得
	elements := strings.Split(filePath, "/")
	fileName := elements[len(elements)-1]

	// ファイル名から日時情報を取り除く
	parts := strings.Split(fileName, "_")
	if len(parts) > 1 {
		fileName = parts[1]
	}

	return fileName
}

func main() {
	userOutput := "../../docs/sql"
	userYamlFiles, err := filepath.Glob("../../docs/entity/*.yaml")
	if err != nil {
		log.Fatalf("Error finding YAML files: %v", err)
	}

	for _, yamlFile := range userYamlFiles {
		err := generateEntity(yamlFile, userOutput)
		if err != nil {
			log.Printf("Error generating entity from YAML file %s: %v", yamlFile, err)
		}
	}
}
