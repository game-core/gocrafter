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

	"github.com/redis/go-redis/v9"

	{{.Import}}
	"github.com/game-core/gocrafter/configs/database"
	"github.com/game-core/gocrafter/internal/keys"
	"github.com/game-core/gocrafter/internal/errors"
)

type {{.CamelName}}Dao struct {
	ReadRedisConn  *redis.Client
	WriteRedisConn *redis.Client
}

func New{{.Name}}Dao(conn *database.RedisHandler) {{.Package}}.{{.Name}}RedisRepository {
	return &{{.CamelName}}Dao{
		ReadRedisConn:  conn.Common.ReadRedisConn,
		WriteRedisConn: conn.Common.WriteRedisConn,
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

	if err := NewMysqlRepository().generate(path, domainPath); err != nil {
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

	// Set
	if len(yamlStruct.Primary) > 0 {
		methods = append(methods, s.createSet(yamlStruct, strings.Split(yamlStruct.Primary[0], ",")))
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
	var keyList []string
	for _, field := range primaryFields {
		keys[field] = yamlStruct.Structures[field]
		keyList = append(keyList, changes.UpperCamelToCamel(field))
	}

	return fmt.Sprintf(
		`func (s *%sDao) Find(ctx context.Context, %s) (*%s.%s, error) {
			t := New%s()
			data, err := s.ReadRedisConn.HGet(ctx, t.TableName(), %s).Result()
			if err != nil {
				return nil, err
			}
		
			if err := t.JsonToTable(data); err != nil {
				return nil, err
			}
		
			return %s, nil
		}`,
		changes.UpperCamelToCamel(yamlStruct.Name),
		s.createParam(keys),
		yamlStruct.Package,
		yamlStruct.Name,
		yamlStruct.Name,
		s.createKey(keyList, keyList),
		s.createModelSetter(yamlStruct),
	)
}

// createFindOrNil FindOrNilを作成する
func (s *Dao) createFindOrNil(yamlStruct *YamlStruct, primaryFields []string) string {
	keys := make(map[string]Structure)
	var keyList []string
	for _, field := range primaryFields {
		keys[field] = yamlStruct.Structures[field]
		keyList = append(keyList, changes.UpperCamelToCamel(field))
	}

	return fmt.Sprintf(
		`func (s *%sDao) FindOrNil(ctx context.Context, %s) (*%s.%s, error) {
			t := New%s()
			data, err := s.ReadRedisConn.HGet(ctx, t.TableName(), %s).Result()
			if err != nil {
				if err == redis.Nil {
					return nil, nil
				}
				return nil, err
			}
		
			if err := t.JsonToTable(data); err != nil {
				return nil, err
			}
		
			return %s, nil
		}`,
		changes.UpperCamelToCamel(yamlStruct.Name),
		s.createParam(keys),
		yamlStruct.Package,
		yamlStruct.Name,
		yamlStruct.Name,
		s.createKey(keyList, keyList),
		s.createModelSetter(yamlStruct),
	)
}

// createSet Setを作成する
func (s *Dao) createSet(yamlStruct *YamlStruct, primaryFields []string) string {
	var keys []string
	var values []string
	for _, field := range primaryFields {
		keys = append(keys, changes.UpperCamelToCamel(field))
		values = append(values, fmt.Sprintf("m.%s", field))
	}

	return fmt.Sprintf(
		`func (s *%sDao) Set(ctx context.Context, tx redis.Pipeliner, m *%s.%s) (*%s.%s, error) {
			var conn redis.Pipeliner
			if tx != nil {
				conn = tx
			} else {
				conn = s.WriteRedisConn.TxPipeline()
			}
		
			t := %s
		
			jt, err := t.TableToJson()
			if err != nil {
				return nil, err
			}
		
			if err := conn.HSet(ctx, t.TableName(), %s, jt).Err(); err != nil {
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
		s.createKey(keys, values),
		s.createModelSetter(yamlStruct),
	)
}

// createDelete Deleteを作成する
func (s *Dao) createDelete(yamlStruct *YamlStruct, primaryFields []string) string {
	var keys []string
	var values []string
	for _, field := range primaryFields {
		keys = append(keys, changes.UpperCamelToCamel(field))
		values = append(values, fmt.Sprintf("m.%s", field))
	}

	return fmt.Sprintf(
		`func (s *%sDao) Delete(ctx context.Context, tx redis.Pipeliner, m *%s.%s) error {
			var conn redis.Pipeliner
			if tx != nil {
				conn = tx
			} else {
				conn = s.WriteRedisConn.TxPipeline()
			}
		
			t := New%s()
			if err := conn.HDel(ctx, t.TableName(), %s).Err(); err != nil {
				return err
			}
		
			return nil
		}`,
		changes.UpperCamelToCamel(yamlStruct.Name),
		yamlStruct.Package,
		yamlStruct.Name,
		yamlStruct.Name,
		s.createKey(keys, values),
	)
}

// createKey Keyを作成する
func (s *Dao) createKey(keys, values []string) string {
	var params []string
	for _, k := range keys {
		params = append(params, k+":%v")
	}

	return "fmt.Sprintf(\"%s:" + strings.Join(params, ",") + "\", t.TableName(), " + strings.Join(values, ",") + ")"
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
