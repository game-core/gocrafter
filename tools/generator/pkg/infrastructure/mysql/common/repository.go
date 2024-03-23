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

type MysqlRepository struct{}

func NewMysqlRepository() *MysqlRepository {
	return &MysqlRepository{}
}

const repositoryTemplate = `
{{.Mock}}
// Package {{.Package}} {{.Comment}}
package {{.Package}}

import (
	"context"
	"time"

	"gorm.io/gorm"

	"github.com/game-core/gocrafter/pkg/domain/enum"
)

{{.Script}}
`

// generate 生成する
func (s *MysqlRepository) generate(path string, base string) error {
	yamlStruct, err := s.getYamlStruct(path)
	if err != nil {
		return err
	}

	if err := os.MkdirAll(base, os.ModePerm); err != nil {
		return err
	}

	if err := s.createOutputFile(yamlStruct, s.getOutputFileName(base, changes.UpperCamelToSnake(yamlStruct.Name))); err != nil {
		return err
	}

	return nil
}

// getYamlStruct yaml構造体を取得する
func (s *MysqlRepository) getYamlStruct(file string) (*YamlStruct, error) {
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
func (s *MysqlRepository) createOutputFile(yamlStruct *YamlStruct, outputFileName string) error {
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
func (s *MysqlRepository) getOutputFileName(dir, name string) string {
	return filepath.Join(dir, fmt.Sprintf("%s_repository.gen.go", changes.UpperCamelToSnake(name)))
}

// createTemplate テンプレートを作成する
func (s *MysqlRepository) createTemplate(yamlStruct *YamlStruct, outputFile *os.File) error {
	tmp, err := template.New("repositoryTemplate").Parse(repositoryTemplate)
	if err != nil {
		return err
	}

	if err := tmp.ExecuteTemplate(
		outputFile,
		"repositoryTemplate",
		TemplateStruct{
			Name:       yamlStruct.Name,
			Package:    yamlStruct.Package,
			PluralName: changes.SingularToPlural(yamlStruct.Name),
			CamelName:  changes.UpperCamelToCamel(yamlStruct.Name),
			Comment:    yamlStruct.Comment,
			Script:     s.createScript(yamlStruct),
			Import:     importCode,
			Mock:       createMock(yamlStruct),
		},
	); err != nil {
		return err
	}

	return nil
}

// createMock mockを作成する
func createMock(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		"//go:generate mockgen -source=./%s_repository.gen.go -destination=./%s_repository_mock.gen.go -package=%s",
		changes.UpperCamelToSnake(yamlStruct.Name),
		changes.UpperCamelToSnake(yamlStruct.Name),
		changes.UpperCamelToCamel(yamlStruct.Package),
	)
}

// createScript スクリプトを作成する
func (s *MysqlRepository) createScript(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`type %sMysqlRepository interface {
			%s
		}`,
		yamlStruct.Name,
		strings.Join(s.createMethods(yamlStruct), "\n"),
	)
}

// createMethods メソッドを作成する
func (s *MysqlRepository) createMethods(yamlStruct *YamlStruct) []string {
	var methods []string

	// Find
	if len(yamlStruct.Primary) > 0 {
		methods = append(methods, s.createFind(yamlStruct, strings.Split(yamlStruct.Primary[0], ",")))
	}

	// FindOrNil
	if len(yamlStruct.Primary) > 0 {
		methods = append(methods, s.createFindOrNil(yamlStruct, strings.Split(yamlStruct.Primary[0], ",")))
	}

	// FindByIndex
	for _, index := range yamlStruct.Index {
		methods = append(methods, s.createFindByIndex(yamlStruct, strings.Split(index, ",")))
	}

	// FindOrNilByIndex
	for _, index := range yamlStruct.Index {
		methods = append(methods, s.createFindOrNilByIndex(yamlStruct, strings.Split(index, ",")))
	}

	// FindList
	methods = append(methods, s.createFindList(yamlStruct))

	// ListByIndex
	for _, index := range yamlStruct.Index {
		methods = append(methods, s.createFindListByIndex(yamlStruct, strings.Split(index, ",")))
	}

	// Create
	methods = append(methods, s.createCreate(yamlStruct))

	// CreateList
	methods = append(methods, s.createCreateList(yamlStruct))

	// Update
	if len(yamlStruct.Primary) > 0 {
		methods = append(methods, s.createUpdate(yamlStruct, strings.Split(yamlStruct.Primary[0], ",")))
	}

	// Delete
	if len(yamlStruct.Primary) > 0 {
		methods = append(methods, s.createDelete(yamlStruct, strings.Split(yamlStruct.Primary[0], ",")))
	}

	return methods
}

// createFind Findを作成する
func (s *MysqlRepository) createFind(yamlStruct *YamlStruct, primaryFields []string) string {
	keys := make(map[string]Structure)
	for _, field := range primaryFields {
		keys[field] = yamlStruct.Structures[field]
	}

	return fmt.Sprintf(
		`Find(ctx context.Context, %s) (*%s, error)`,
		s.createParam(keys),
		yamlStruct.Name,
	)
}

// createFindOrNil FindOrNilを作成する
func (s *MysqlRepository) createFindOrNil(yamlStruct *YamlStruct, primaryFields []string) string {
	keys := make(map[string]Structure)
	for _, field := range primaryFields {
		keys[field] = yamlStruct.Structures[field]
	}

	return fmt.Sprintf(
		`FindOrNil(ctx context.Context, %s) (*%s, error)`,
		s.createParam(keys),
		yamlStruct.Name,
	)
}

// createFindByIndex FindByIndexを作成する
func (s *MysqlRepository) createFindByIndex(yamlStruct *YamlStruct, indexFields []string) string {
	keys := make(map[string]Structure)
	for _, field := range indexFields {
		keys[field] = yamlStruct.Structures[field]
	}

	return fmt.Sprintf(
		`FindBy%s(ctx context.Context, %s) (*%s, error)`,
		strings.Join(indexFields, "And"),
		s.createParam(keys),
		yamlStruct.Name,
	)
}

// createFindOrNilByIndex FindOrNilByIndexを作成する
func (s *MysqlRepository) createFindOrNilByIndex(yamlStruct *YamlStruct, indexFields []string) string {
	keys := make(map[string]Structure)
	for _, field := range indexFields {
		keys[field] = yamlStruct.Structures[field]
	}

	return fmt.Sprintf(
		`FindOrNilBy%s(ctx context.Context, %s) (*%s, error)`,
		strings.Join(indexFields, "And"),
		s.createParam(keys),
		yamlStruct.Name,
	)
}

// createFindList FindListを作成する
func (s *MysqlRepository) createFindList(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`FindList(ctx context.Context) (%s, error)`,
		changes.SnakeToUpperCamel(changes.SingularToPlural(changes.UpperCamelToSnake(yamlStruct.Name))),
	)
}

// createFindListByIndex FindListByIndexを作成する
func (s *MysqlRepository) createFindListByIndex(yamlStruct *YamlStruct, indexFields []string) string {
	keys := make(map[string]Structure)
	for _, field := range indexFields {
		keys[field] = yamlStruct.Structures[field]
	}

	return fmt.Sprintf(
		`FindListBy%s(ctx context.Context, %s) (%s, error)`,
		strings.Join(indexFields, "And"),
		s.createParam(keys),
		changes.SnakeToUpperCamel(changes.SingularToPlural(changes.UpperCamelToSnake(yamlStruct.Name))),
	)
}

// createCreate Createを作成する
func (s *MysqlRepository) createCreate(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`Create(ctx context.Context, tx *gorm.DB, m *%s) (*%s, error)`,
		yamlStruct.Name,
		yamlStruct.Name,
	)
}

// createCreate CreateListを作成する
func (s *MysqlRepository) createCreateList(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`CreateList(ctx context.Context, tx *gorm.DB, ms %s) (%s, error)`,
		changes.SnakeToUpperCamel(changes.SingularToPlural(changes.UpperCamelToSnake(yamlStruct.Name))),
		changes.SnakeToUpperCamel(changes.SingularToPlural(changes.UpperCamelToSnake(yamlStruct.Name))),
	)
}

// createUpdate Updateを作成する
func (s *MysqlRepository) createUpdate(yamlStruct *YamlStruct, primaryFields []string) string {
	keys := make(map[string]Structure)
	for _, field := range primaryFields {
		keys[field] = yamlStruct.Structures[field]
	}

	return fmt.Sprintf(
		`Update(ctx context.Context, tx *gorm.DB, m *%s) (*%s, error)`,
		yamlStruct.Name,
		yamlStruct.Name,
	)
}

// createDelete Deleteを作成する
func (s *MysqlRepository) createDelete(yamlStruct *YamlStruct, primaryFields []string) string {
	keys := make(map[string]Structure)
	for _, field := range primaryFields {
		keys[field] = yamlStruct.Structures[field]
	}

	return fmt.Sprintf(
		`Delete(ctx context.Context, tx *gorm.DB, m *%s) error`,
		yamlStruct.Name,
	)
}

// createParam Paramを作成する
func (s *MysqlRepository) createParam(keys map[string]Structure) string {
	var paramStrings []string
	for _, field := range s.getStructures(keys) {
		paramStrings = append(paramStrings, fmt.Sprintf("%s %s", changes.SnakeToCamel(field.Name), s.getType(field)))
	}

	return strings.Join(paramStrings, ",")
}

// getStructures フィールド構造体を取得する
func (s *MysqlRepository) getStructures(structures map[string]Structure) []*Structure {
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
func (s *MysqlRepository) getType(field *Structure) string {
	var result string

	switch field.Type {
	case "time":
		result = "time.Time"
	case "structure":
		if field.Package != "" {
			result = fmt.Sprintf("%s.%s", changes.SnakeToCamel(field.Name), changes.SnakeToUpperCamel(field.Name))
		} else {
			result = changes.SnakeToUpperCamel(field.Name)
		}
	case "structures":
		if field.Package != "" {
			result = fmt.Sprintf("%s.%s", changes.SnakeToCamel(field.Name), changes.SnakeToUpperCamel(changes.SingularToPlural(field.Name)))
		} else {
			result = changes.SnakeToUpperCamel(changes.SingularToPlural(field.Name))
		}
	case "enum":
		result = fmt.Sprintf("enum.%s", changes.SnakeToUpperCamel(field.Name))
	default:
		result = field.Type
	}

	if field.Nullable {
		result = fmt.Sprintf("*%s", result)
	}

	return result
}
