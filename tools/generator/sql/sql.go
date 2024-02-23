package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"text/template"
	"time"

	"github.com/game-core/gocrafter/internal"
)

const sqlTemplate = `{{.Script}}
`

type Sql struct{}

func NewSql() *Sql {
	return &Sql{}
}

// generate 生成する
func (s *Sql) generate(path string, base string) error {
	yamlStruct, err := s.getYamlStruct(path)
	if err != nil {
		return err
	}

	if s.checkFileExists(base, yamlStruct.Name) {
		return nil
	}

	if err := os.MkdirAll(base, os.ModePerm); err != nil {
		return err
	}

	if err := s.createOutputFile(yamlStruct, s.getOutputFileName(base, yamlStruct.Name)); err != nil {
		return err
	}

	return nil
}

// getYamlStruct yaml構造体を取得する
func (s *Sql) getYamlStruct(file string) (*YamlStruct, error) {
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
func (s *Sql) getOutputFileName(dir, name string) string {
	return filepath.Join(dir, fmt.Sprintf("%s.sql", fmt.Sprintf("%s_%s", time.Now().Format("20060102"), internal.UpperCamelToSnake(name))))
}

// checkFileExists ファイルがあるか確認する
func (s *Sql) checkFileExists(directory, name string) bool {
	targetFileName := fmt.Sprintf("%s.sql", internal.UpperCamelToSnake(s.extractFileName(name)))
	fileList, err := filepath.Glob(filepath.Join(directory, "*"))
	if err != nil {
		return false
	}

	for _, existingFile := range fileList {
		existingFileName := s.extractFileName(existingFile)
		if existingFileName == targetFileName {
			return true
		}
	}

	return false
}

// extractFileName ファイル名を抽出する
func (s *Sql) extractFileName(filePath string) string {
	elements := strings.Split(filePath, "/")
	fileName := elements[len(elements)-1]
	parts := strings.Split(fileName, "_")

	if len(parts) > 1 {
		dateRegex := regexp.MustCompile(`^\d{8}`)
		match := dateRegex.FindString(parts[0])
		if match != "" {
			fileName = fileName[len(match)+1:]
		}
	}

	return fileName
}

// createOutputFile ファイルを作成する
func (s *Sql) createOutputFile(yamlStruct *YamlStruct, outputFileName string) error {
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}

	if err := s.createTemplate(yamlStruct, outputFile); err != nil {
		return err
	}

	return nil
}

// createTemplate テンプレートを作成する
func (s *Sql) createTemplate(yamlStruct *YamlStruct, outputFile *os.File) error {
	tmp, err := template.New("sqlTemplate").Parse(sqlTemplate)
	if err != nil {
		return err
	}

	if err := tmp.ExecuteTemplate(
		outputFile,
		"sqlTemplate",
		TemplateStruct{
			Script: s.createScript(yamlStruct),
		},
	); err != nil {
		return err
	}

	return nil
}

// createScript スクリプトを作成する
func (s *Sql) createScript(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`CREATE TABLE %s
(
    %s
) DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci;`,
		internal.UpperCamelToSnake(yamlStruct.Name),
		strings.Join(s.createFields(yamlStruct), ",\n"),
	)
}

// createFields fieldを作成する
func (s *Sql) createFields(yamlStruct *YamlStruct) []string {
	var fields []string

	for i, field := range s.getStructure(yamlStruct.Structures) {
		fields = append(
			fields,
			fmt.Sprintf(
				`%s%s %s %s %sCOMMENT "%s"`,
				s.getSpace(i),
				field.Name,
				s.getType(field.Type),
				s.getDefault(field.Nullable),
				s.getOption(field.Name),
				field.Comment,
			),
		)
	}

	for _, primary := range yamlStruct.Primary {
		fields = append(fields, fmt.Sprintf("	PRIMARY KEY(%s)", s.getKey(primary)))
	}

	for _, unique := range yamlStruct.Unique {
		fields = append(fields, fmt.Sprintf("	UNIQUE KEY(%s)", s.getKey(unique)))
	}

	for _, index := range yamlStruct.Index {
		fields = append(fields, fmt.Sprintf("	INDEX(%s)", s.getKey(index)))
	}

	return fields
}

// getSpace
func (s *Sql) getSpace(i int) string {
	if i == 0 {
		return ""
	}

	return "	"
}

// getStructure フィールド構造体を取得する
func (s *Sql) getStructure(structures map[string]Structure) []*Structure {
	var sortStructures []*Structure
	for _, value := range structures {
		sortStructures = append(
			sortStructures,
			&Structure{
				Name:     value.Name,
				Type:     value.Type,
				Package:  value.Package,
				Nullable: value.Nullable,
				Number:   value.Number,
				Comment:  value.Comment,
			},
		)
	}

	sort.Slice(sortStructures, func(i, j int) bool {
		return sortStructures[i].Number < sortStructures[j].Number
	})

	return sortStructures
}

// getType 型を取得する
func (s *Sql) getType(t string) string {
	switch t {
	case "string":
		return "VARCHAR(255)"
	case "int64":
		return "BIGINT"
	case "int32":
		return "INT"
	case "int":
		return "INT"
	case "enum":
		return "INT"
	case "bool":
		return "INT"
	case "time":
		return "TIMESTAMP"
	default:
		return "VARCHAR(255)"
	}
}

// getOption オプションを取得する
func (s *Sql) getOption(name string) string {
	if name == "id" {
		return "AUTO_INCREMENT "
	}

	return ""
}

// getDefault デフォルト値を取得する
func (s *Sql) getDefault(p bool) string {
	if p {
		return "DEFAULT NULL"
	}

	return "NOT NULL"
}

func (s *Sql) getKey(keys string) string {
	keyList := strings.Split(keys, ",")
	var k []string
	for _, key := range keyList {
		k = append(k, internal.UpperCamelToSnake(key))
	}

	return strings.Join(k, ",")
}
