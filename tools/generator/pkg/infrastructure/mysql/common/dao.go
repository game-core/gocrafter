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
	"github.com/game-core/gocrafter/internal/errors"
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
	"github.com/game-core/gocrafter/internal/cashes"
	"github.com/game-core/gocrafter/internal/errors"
)

type {{.CamelName}}Dao struct {
	ReadConn  *gorm.DB
	WriteConn *gorm.DB
}

func New{{.Name}}Dao(conn *database.SqlHandler) {{.Package}}.{{.Name}}Repository {
	return &{{.CamelName}}Dao{
		ReadConn:  conn.Common.ReadConn,
		WriteConn: conn.Common.WriteConn,
	}
}

{{.Script}}
`

type Dao struct{}

func NewDao() *Dao {
	return &Dao{}
}

// generate 生成する
func (s *Dao) generate(path string, base string) error {
	importCode = ""

	yamlStruct, err := s.getYamlStruct(path)
	if err != nil {
		return err
	}

	domainPath, err := s.getDomainPath(fmt.Sprintf("%s_model.gen.go", changes.UpperCamelToSnake(yamlStruct.Name)))
	if err != nil {
		return err
	}

	if err := NewRepository().generate(path, domainPath); err != nil {
		return err
	}

	outputDir := filepath.Join(base, yamlStruct.Package)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return err
	}

	if err := s.createOutputFile(yamlStruct, s.getOutputFileName(outputDir, changes.UpperCamelToSnake(yamlStruct.Name))); err != nil {
		return err
	}

	return nil
}

// getDomainPath ドメインのpathを取得する関数
func (s *Dao) getDomainPath(name string) (string, error) {
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
		return "", err
	}

	if target == "" {
		return "", errors.NewError("file does not exist")
	}

	importPath := fmt.Sprintf("\"github.com/game-core/gocrafter/%s\"", strings.Replace(target, "../../../../../../", "", -1))
	importCode = fmt.Sprintf("%s\n%s", importCode, importPath)

	return target, nil
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
	return filepath.Join(dir, fmt.Sprintf("%s_dao.gen.go", changes.UpperCamelToSnake(name)))
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
			PluralName: changes.SingularToPlural(yamlStruct.Name),
			CamelName:  changes.UpperCamelToCamel(yamlStruct.Name),
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
			method,
		)
	}

	return methods
}

// createMethods メソッドを作成する
func (s *Dao) createMethods(yamlStruct *YamlStruct) []string {
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
func (s *Dao) createFind(yamlStruct *YamlStruct, primaryFields []string) string {
	keys := make(map[string]Structure)
	for _, field := range primaryFields {
		keys[field] = yamlStruct.Structures[field]
	}

	return fmt.Sprintf(
		`func (s *%sDao) Find(ctx context.Context, %s) (*%s.%s, error) {
			t := New%s()
			res := s.ReadConn.WithContext(ctx).%s.Find(t)
			if err := res.Error; err != nil {
				return nil, err
			}
			if res.RowsAffected == 0 {
				return nil, errors.NewError("record does not exist")
			}

			return %s, nil
		}
		`,
		changes.UpperCamelToCamel(yamlStruct.Name),
		s.createParam(keys),
		yamlStruct.Package,
		yamlStruct.Name,
		yamlStruct.Name,
		s.createQuery(keys),
		s.createModelSetter(yamlStruct),
	)
}

// createFindOrNil FindOrNilを作成する
func (s *Dao) createFindOrNil(yamlStruct *YamlStruct, primaryFields []string) string {
	keys := make(map[string]Structure)
	for _, field := range primaryFields {
		keys[field] = yamlStruct.Structures[field]
	}

	return fmt.Sprintf(
		`func (s *%sDao) FindOrNil(ctx context.Context, %s) (*%s.%s, error) {
			t := New%s()
			res := s.ReadConn.WithContext(ctx).%s.Find(t)
			if err := res.Error; err != nil {
				return nil, err
			}
			if res.RowsAffected == 0 {
				return nil, nil
			}

			return %s, nil
		}
		`,
		changes.UpperCamelToCamel(yamlStruct.Name),
		s.createParam(keys),
		yamlStruct.Package,
		yamlStruct.Name,
		yamlStruct.Name,
		s.createQuery(keys),
		s.createModelSetter(yamlStruct),
	)
}

// createFindByIndex FindByIndexを作成する
func (s *Dao) createFindByIndex(yamlStruct *YamlStruct, indexFields []string) string {
	keys := make(map[string]Structure)
	for _, field := range indexFields {
		keys[field] = yamlStruct.Structures[field]
	}

	return fmt.Sprintf(
		`func (s *%sDao) FindBy%s(ctx context.Context, %s) (*%s.%s, error) {
			t := New%s()
			res := s.ReadConn.WithContext(ctx).%s.Find(t)
			if err := res.Error; err != nil {
				return nil, err
			}
			if res.RowsAffected == 0 {
				return nil, errors.NewError("record does not exist")
			}

			return %s, nil
		}
		`,
		changes.UpperCamelToCamel(yamlStruct.Name),
		strings.Join(indexFields, "And"),
		s.createParam(keys),
		yamlStruct.Package,
		yamlStruct.Name,
		yamlStruct.Name,
		s.createQuery(keys),
		s.createModelSetter(yamlStruct),
	)
}

// createFindOrNilByIndex FindOrNilByIndexを作成する
func (s *Dao) createFindOrNilByIndex(yamlStruct *YamlStruct, indexFields []string) string {
	keys := make(map[string]Structure)
	for _, field := range indexFields {
		keys[field] = yamlStruct.Structures[field]
	}

	return fmt.Sprintf(
		`func (s *%sDao) FinOrNilBy%s(ctx context.Context, %s) (*%s.%s, error) {
			t := New%s()
			res := s.ReadConn.WithContext(ctx).%s.Find(t)
			if err := res.Error; err != nil {
				return nil, err
			}
			if res.RowsAffected == 0 {
				return nil, nil
			}

			return %s, nil
		}
		`,
		changes.UpperCamelToCamel(yamlStruct.Name),
		strings.Join(indexFields, "And"),
		s.createParam(keys),
		yamlStruct.Package,
		yamlStruct.Name,
		yamlStruct.Name,
		s.createQuery(keys),
		s.createModelSetter(yamlStruct),
	)
}

// createFindList FindListを作成する
func (s *Dao) createFindList(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`func (s *%sDao) FindList(ctx context.Context) (%s.%s, error) {
			ts := New%s()
			res := s.ReadConn.WithContext(ctx).Find(&ts)
			if err := res.Error; err != nil {
				return nil, err
			}

			%s

			return ms, nil
		}
		`,
		changes.UpperCamelToCamel(yamlStruct.Name),
		yamlStruct.Package,
		changes.SingularToPlural(yamlStruct.Name),
		changes.SingularToPlural(yamlStruct.Name),
		s.createModelSetters(yamlStruct),
	)
}

// createFindListByIndex FindListByIndexを作成する
func (s *Dao) createFindListByIndex(yamlStruct *YamlStruct, indexFields []string) string {
	keys := make(map[string]Structure)
	for _, field := range indexFields {
		keys[field] = yamlStruct.Structures[field]
	}

	return fmt.Sprintf(
		`func (s *%sDao) FindListBy%s(ctx context.Context, %s) (%s.%s, error) {
			ts := New%s()
			res := s.ReadConn.WithContext(ctx).%s.Find(&ts)
			if err := res.Error; err != nil {
				return nil, err
			}

			%s

			return ms, nil
		}
		`,
		changes.UpperCamelToCamel(yamlStruct.Name),
		strings.Join(indexFields, "And"),
		s.createParam(keys),
		yamlStruct.Package,
		changes.SingularToPlural(yamlStruct.Name),
		changes.SingularToPlural(yamlStruct.Name),
		s.createQuery(keys),
		s.createModelSetters(yamlStruct),
	)
}

// createCreate Createを作成する
func (s *Dao) createCreate(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`func (s *%sDao) Create(ctx context.Context, tx *gorm.DB, m *%s.%s) (*%s.%s, error) {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = s.WriteConn
			}

			t := %s
			res := conn.Model(New%s()).WithContext(ctx).Create(t)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return %s, nil
		}`,
		changes.UpperCamelToCamel(yamlStruct.Name),
		yamlStruct.Package,
		yamlStruct.Name,
		yamlStruct.Package,
		yamlStruct.Name,
		s.createTableSetter(yamlStruct),
		yamlStruct.Name,
		s.createModelSetter(yamlStruct),
	)
}

// createCreate CreateListを作成する
func (s *Dao) createCreateList(yamlStruct *YamlStruct) string {
	return fmt.Sprintf(
		`func (s *%sDao) CreateList(ctx context.Context, tx *gorm.DB, ms %s.%s) (%s.%s, error) {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = s.WriteConn
			}

			ts := New%s()
			for _, m := range ms {
				t := %s
				ts = append(ts, t)
			}

			res := conn.Model(New%s()).WithContext(ctx).Create(ts)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return ms, nil
		}`,
		changes.UpperCamelToCamel(yamlStruct.Name),
		yamlStruct.Package,
		changes.SingularToPlural(yamlStruct.Name),
		yamlStruct.Package,
		changes.SingularToPlural(yamlStruct.Name),
		changes.SingularToPlural(yamlStruct.Name),
		s.createTableSetter(yamlStruct),
		yamlStruct.Name,
	)
}

// createUpdate Updateを作成する
func (s *Dao) createUpdate(yamlStruct *YamlStruct, primaryFields []string) string {
	keys := make(map[string]Structure)
	for _, field := range primaryFields {
		keys[field] = yamlStruct.Structures[field]
	}

	return fmt.Sprintf(
		`func (s *%sDao) Update(ctx context.Context, tx *gorm.DB, m *%s.%s) (*%s.%s, error) {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = s.WriteConn
			}

			t := %s
			res := conn.Model(New%s()).WithContext(ctx).%s.Updates(t)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return %s, nil
		}`,
		changes.UpperCamelToCamel(yamlStruct.Name),
		yamlStruct.Package,
		yamlStruct.Name,
		yamlStruct.Package,
		yamlStruct.Name,
		s.createTableSetter(yamlStruct),
		yamlStruct.Name,
		s.createModelQuery(keys),
		s.createModelSetter(yamlStruct),
	)
}

// createDelete Deleteを作成する
func (s *Dao) createDelete(yamlStruct *YamlStruct, primaryFields []string) string {
	keys := make(map[string]Structure)
	for _, field := range primaryFields {
		keys[field] = yamlStruct.Structures[field]
	}

	return fmt.Sprintf(
		`func (s *%sDao) Delete(ctx context.Context, tx *gorm.DB, m *%s.%s) error {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = s.WriteConn
			}
		
			res := conn.Model(New%s()).WithContext(ctx).%s.Delete(New%s())
			if err := res.Error; err != nil {
				return err
			}
		
			return nil
		}`,
		changes.UpperCamelToCamel(yamlStruct.Name),
		yamlStruct.Package,
		yamlStruct.Name,
		yamlStruct.Name,
		s.createModelQuery(keys),
		yamlStruct.Name,
	)
}

// createQuery Queryを作成する
func (s *Dao) createQuery(keys map[string]Structure) string {
	var queryStrings []string
	for _, field := range s.getStructures(keys) {
		queryStrings = append(queryStrings, fmt.Sprintf("Where(\"%s = ?\", %s)", field.Name, changes.SnakeToCamel(field.Name)))
	}

	return strings.Join(queryStrings, ".")
}

// createModelQuery Queryを作成する
func (s *Dao) createModelQuery(keys map[string]Structure) string {
	var queryStrings []string
	for _, field := range s.getStructures(keys) {
		queryStrings = append(queryStrings, fmt.Sprintf("Where(\"%s = ?\", m.%s)", field.Name, changes.SnakeToUpperCamel(field.Name)))
	}

	return strings.Join(queryStrings, ".")
}

// createParam Paramを作成する
func (s *Dao) createParam(keys map[string]Structure) string {
	var paramStrings []string
	for _, field := range s.getStructures(keys) {
		paramStrings = append(paramStrings, fmt.Sprintf("%s %s", changes.SnakeToCamel(field.Name), s.getType(field)))
	}

	return strings.Join(paramStrings, ",")
}

// createModelSetter createModelSetterを作成する
func (s *Dao) createModelSetter(yamlStruct *YamlStruct) string {
	var paramStrings []string
	for _, field := range s.getStructures(yamlStruct.Structures) {
		if field.Name != "created_at" && field.Name != "updated_at" {
			paramStrings = append(paramStrings, fmt.Sprintf("t.%s,", changes.SnakeToUpperCamel(field.Name)))
		}
	}

	return fmt.Sprintf(
		`%s.Set%s(%s)`,
		yamlStruct.Package,
		yamlStruct.Name,
		strings.Join(paramStrings, ""),
	)
}

// createModelSetters createModelSettersを作成する
func (s *Dao) createModelSetters(yamlStruct *YamlStruct) string {
	var paramStrings []string
	for _, field := range s.getStructures(yamlStruct.Structures) {
		if field.Name != "created_at" && field.Name != "updated_at" {
			paramStrings = append(paramStrings, fmt.Sprintf("t.%s,", changes.SnakeToUpperCamel(field.Name)))
		}
	}

	return fmt.Sprintf(
		`ms := %s.New%s()
		for _, t := range ts {
			ms = append(ms, %s)
		}`,
		yamlStruct.Package,
		changes.SingularToPlural(yamlStruct.Name),
		fmt.Sprintf(
			`%s.Set%s(%s)`,
			yamlStruct.Package,
			yamlStruct.Name,
			strings.Join(paramStrings, ""),
		),
	)
}

// createTableSetter createTableSetterを作成する
func (s *Dao) createTableSetter(yamlStruct *YamlStruct) string {
	var paramStrings []string
	for _, field := range s.getStructures(yamlStruct.Structures) {
		if field.Name != "created_at" && field.Name != "updated_at" {
			paramStrings = append(paramStrings, fmt.Sprintf("%s: m.%s,", changes.SnakeToUpperCamel(field.Name), changes.SnakeToUpperCamel(field.Name)))
		}
	}

	return fmt.Sprintf(
		`&%s{
			%s
		}`,
		yamlStruct.Name,
		strings.Join(paramStrings, "\n"),
	)
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
			result = fmt.Sprintf("%s.%s", changes.SnakeToCamel(field.Name), changes.SnakeToUpperCamel(field.Name))
		} else {
			result = changes.SnakeToUpperCamel(field.Name)
		}
	case "structures":
		if field.Package != "" {
			importCode = fmt.Sprintf("%s\n%s", importCode, fmt.Sprintf("\"github.com/game-core/gocrafter/pkg/domain/model/%s\"", field.Package))
			result = fmt.Sprintf("%s.%s", changes.SnakeToCamel(field.Name), changes.SnakeToUpperCamel(changes.SingularToPlural(field.Name)))
		} else {
			result = changes.SnakeToUpperCamel(changes.SingularToPlural(field.Name))
		}
	case "enum":
		importCode = fmt.Sprintf("%s\n%s", importCode, "\"github.com/game-core/gocrafter/pkg/domain/enum\"")
		result = fmt.Sprintf("enum.%s", changes.SnakeToUpperCamel(field.Name))
	default:
		result = field.Type
	}

	if field.Nullable {
		result = fmt.Sprintf("*%s", result)
	}

	return result
}
