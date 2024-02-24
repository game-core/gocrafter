package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"

	"gopkg.in/yaml.v3"

	"github.com/game-core/gocrafter/internal/changes"
)

const enumTemplate = `
// Package {{.Package}} {{.Comment}}
package {{.Package}}

type {{.Name}} int32

{{.Script}}
`

type Enum struct{}

func NewEnum() *Enum {
	return &Enum{}
}

// generate 生成する
func (s *Enum) generate(path string, base string) error {
	yamlStruct, err := s.getYamlStruct(path)
	if err != nil {
		return err
	}

	outputDir := filepath.Join(base, strings.Replace(filepath.Dir(path), "/../../docs/yaml", "", -1))
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	if err := s.createOutputFile(yamlStruct, s.getOutputFileName(outputDir, filepath.Base(path[:len(path)-len(filepath.Ext(path))]))); err != nil {
		return err
	}

	return nil
}

// getYamlStruct yaml構造体を取得する
func (s *Enum) getYamlStruct(file string) (*YamlStruct, error) {
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

// createOutputFile ファイルを作成する
func (s *Enum) createOutputFile(yamlStruct *YamlStruct, outputFileName string) error {
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return err
	}

	if err := s.createTemplate(yamlStruct, outputFile); err != nil {
		return err
	}

	return nil
}

// getOutputFileName ファイル名を取得する
func (s *Enum) getOutputFileName(dir, name string) string {
	return filepath.Join(dir, fmt.Sprintf("%s_enum.gen.go", changes.UpperCamelToSnake(name)))
}

// createTemplate テンプレートを作成する
func (s *Enum) createTemplate(yamlStruct *YamlStruct, outputFile *os.File) error {
	tmp, err := template.New("enumTemplate").Parse(enumTemplate)
	if err != nil {
		return err
	}

	if err := tmp.ExecuteTemplate(
		outputFile,
		"enumTemplate",
		TemplateStruct{
			Name:    yamlStruct.Name,
			Package: yamlStruct.Package,
			Comment: yamlStruct.Comment,
			Script:  s.createScript(yamlStruct),
		},
	); err != nil {
		return err
	}

	return nil
}

// createScript スクリプトを作成する
func (s *Enum) createScript(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`const (
			%s
		)`,
		strings.Join(s.createConstants(yamlStruct), "\n"),
	)
}

// createConstants
func (s *Enum) createConstants(yamlStruct *YamlStruct) []string {
	var constants []string
	for _, constant := range s.getStructure(yamlStruct.Structures) {
		constants = append(
			constants,
			fmt.Sprintf(
				"%s_%s %s = %d",
				yamlStruct.Name,
				constant.Name,
				yamlStruct.Name,
				constant.Number,
			),
		)
	}

	return constants
}

// getStructure フィールド構造体を取得する
func (s *Enum) getStructure(structures map[string]Structure) []*Structure {
	var sortStructures []*Structure
	for _, value := range structures {
		sortStructures = append(
			sortStructures,
			&Structure{
				Name:    value.Name,
				Number:  value.Number,
				Comment: value.Comment,
			},
		)
	}

	sort.Slice(sortStructures, func(i, j int) bool {
		return sortStructures[i].Number < sortStructures[j].Number
	})

	return sortStructures
}
