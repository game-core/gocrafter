package main

import (
	"fmt"
	"github.com/game-core/gocrafter/internal"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"text/template"
)

const daoTemplate = `
// Package {{.Package}} {{.Comment}}
package {{.Package}}

import (
	"context"
	"time"

	"gorm.io/gorm"

	{{.Import}}
	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal"
)

type {{.CamelName}}Dao struct {
	ShardConn *database.ShardConn
}

func New{{.Name}}Dao(conn *database.SqlHandler) {{.Package}}.{{.Name}}Repository {
	return &{{.CamelName}}Dao{
		ShardConn: conn.User,
	}
}

{{.Script}}
`

type Dao struct{}

func NewDao() *Dao {
	return &Dao{}
}

// generate 生成する
func (s *Dao) generate(file string, base string) error {
	importCode = ""

	yamlStruct, err := s.getYamlStruct(file)
	if err != nil {
		return err
	}

	if err := s.getDomainImportPath(fmt.Sprintf("%s_repository.gen.go", internal.UpperCamelToSnake(yamlStruct.Name))); err != nil {
		return err
	}

	outputDir := filepath.Join(base, yamlStruct.Package)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	if err := s.createOutputFile(yamlStruct, s.getOutputFileName(outputDir, internal.UpperCamelToSnake(yamlStruct.Name))); err != nil {
		return err
	}

	return nil
}

// getDomainImportPath ドメインのpathを取得する関数
func (s *Dao) getDomainImportPath(name string) error {
	base := "../../../../../../pkg/domain/model"
	var target string

	if err := filepath.Walk(base, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() {
			return nil
		}
		if info.Name() == name {
			target = filepath.Dir(path)
		}
		return nil
	}); err != nil {
		return err
	}

	if target == "" {
		return fmt.Errorf("file does not exist")
	}

	importPath := fmt.Sprintf("\"github.com/game-core/gocrafter/%s\"", strings.Replace(target, "../../../../../../", "", -1))
	importCode = fmt.Sprintf("%s\n%s", importCode, importPath)
	fmt.Println(importCode)

	return nil
}

// getYamlStruct yaml構造体を取得する
func (s *Dao) getYamlStruct(file string) (*YamlStruct, error) {
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
func (s *Dao) getOutputFileName(dir, name string) string {
	return filepath.Join(dir, fmt.Sprintf("%s_dao.gen.go", internal.UpperCamelToSnake(name)))
}

// createOutputFile ファイルを作成する
func (s *Dao) createOutputFile(yamlStruct *YamlStruct, outputFileName string) error {
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
func (s *Dao) createTemplate(yamlStruct *YamlStruct, outputFile *os.File) error {
	tmp, err := template.New("daoTemplate").Parse(daoTemplate)
	if err != nil {
		return err
	}

	if err := tmp.ExecuteTemplate(
		outputFile,
		"daoTemplate",
		TemplateStruct{
			Name:       yamlStruct.Name,
			Package:    yamlStruct.Package,
			PluralName: internal.SingularToPlural(yamlStruct.Name),
			CamelName:  internal.UpperCamelToCamel(yamlStruct.Name),
			Comment:    yamlStruct.Comment,
			Script:     s.createScript(yamlStruct),
			Import:     importCode,
		},
	); err != nil {
		return err
	}

	return nil
}

// createScript スクリプトを作成する
func (s *Dao) createScript(yamlStruct *YamlStruct) string {
	var methods string

	for _, method := range s.createMethods(yamlStruct) {
		methods = fmt.Sprintf(
			`%s

			%s`,
			methods,
			method.Script,
		)
	}

	return methods
}

// createMethods メソッドを作成する
func (s *Dao) createMethods(yamlStruct *YamlStruct) map[string]MethodType {
	methods := make(map[string]MethodType)

	// Find
	if len(yamlStruct.Primary) > 0 {
		methods["Find"] = MethodType{
			Script: s.createFind(yamlStruct, strings.Split(yamlStruct.Primary[0], ",")),
		}
	}

	/*
		// FindOrNil
		if len(yamlStruct.Primary) > 0 {
			primaryFields := strings.Split(yamlStruct.Primary[0], ",")
			methods["FindOrNil"] = MethodType{
				Script: createDaoFindOrNil(yamlStruct, primaryFields),
			}
		}

		// FindByIndex
		for _, index := range yamlStruct.Index {
			indexFields := strings.Split(index, ",")
			methods[fmt.Sprintf("FindBy%s", strings.Join(indexFields, "And"))] = MethodType{
				Script: createDaoFindByIndex(yamlStruct, indexFields),
			}
		}

		// FindOrNilByIndex
		for _, index := range yamlStruct.Index {
			indexFields := strings.Split(index, ",")
			methods[fmt.Sprintf("FindOrNilBy%s", strings.Join(indexFields, "And"))] = MethodType{
				Script: createDaoFindOrNilByIndex(yamlStruct, indexFields),
			}
		}

		// List
		methods["FindList"] = MethodType{
			Script: createDaoFindList(yamlStruct),
		}

		// ListByIndex
		for _, index := range yamlStruct.Index {
			indexFields := strings.Split(index, ",")
			methods[fmt.Sprintf("FindListBy%s", strings.Join(indexFields, "And"))] = MethodType{
				Script: createDaoFindListByIndex(yamlStruct, indexFields),
			}
		}

		// Create
		methods["Create"] = MethodType{
			Script: createDaoCreate(yamlStruct),
		}

		// Update
		if len(yamlStruct.Primary) > 0 {
			primaryFields := strings.Split(yamlStruct.Primary[0], ",")
			methods["Update"] = MethodType{
				Script: createDaoUpdate(yamlStruct, primaryFields),
			}
		}

		// Delete
		if len(yamlStruct.Primary) > 0 {
			primaryFields := strings.Split(yamlStruct.Primary[0], ",")
			methods["Delete"] = MethodType{
				Script: createDaoDelete(yamlStruct, primaryFields),
			}
		}

	*/

	return methods
}

// createFind Findを作成する
func (s *Dao) createFind(yamlStruct *YamlStruct, primaryFields []string) string {
	keys := make(map[string]Structure)
	for _, field := range primaryFields {
		keys[field] = yamlStruct.Structures[field]
	}

	return fmt.Sprintf(
		`func (s *%sDao) Find(ctx context.Context, %s) (*%s.%s, error) {
			t := New%s()
			res := s.ShardConn.Shards[internal.GetShardKeyByUserId(userId)].ReadConn.WithContext(ctx).%s.Find(t)
			if err := res.Error; err != nil {
				return nil, err
			}
			if res.RowsAffected == 0 {
				return nil, fmt.Errorf("record does not exist")
			}

			return &%s.%s{
				%s
			}, nil
		}
		`,
		internal.UpperCamelToCamel(yamlStruct.Name),
		s.createParam(keys),
		yamlStruct.Package,
		yamlStruct.Name,
		yamlStruct.Name,
		s.createQuery(keys),
		yamlStruct.Package,
		yamlStruct.Name,
		s.createReturn(yamlStruct.Structures),
	)
}

// createQuery Queryを作成する
func (s *Dao) createQuery(keys map[string]Structure) string {
	var queryStrings []string
	for _, field := range s.getStructures(keys) {
		queryStrings = append(queryStrings, fmt.Sprintf("Where(\"%s = ?\", %s)", field.Name, internal.SnakeToCamel(field.Name)))
	}

	return strings.Join(queryStrings, ",")
}

// createParam 引数を作成する
func (s *Dao) createParam(keys map[string]Structure) string {
	var paramStrings []string
	for _, field := range s.getStructures(keys) {
		paramStrings = append(paramStrings, fmt.Sprintf("%s %s", internal.SnakeToCamel(field.Name), s.getType(field)))
	}

	return strings.Join(paramStrings, ",")
}

// createReturn 返り値のフィールドを作成する
func (s *Dao) createReturn(structures map[string]Structure) string {
	var returnStrings []string
	for _, field := range s.getStructures(structures) {
		if field.Name != "created_at" && field.Name != "updated_at" {
			returnStrings = append(returnStrings, fmt.Sprintf("%s: t.%s,", internal.SnakeToUpperCamel(field.Name), internal.SnakeToUpperCamel(field.Name)))
		}
	}

	return strings.Join(returnStrings, "\n")
}

// getStructures フィールド構造体を取得する
func (s *Dao) getStructures(structures map[string]Structure) []*Structure {
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
func (s *Dao) getType(field *Structure) string {
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
			importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("\"github.com/game-core/gocrafter/pkg/domain/%s\"", field.Package))
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
