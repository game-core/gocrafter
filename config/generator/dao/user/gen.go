//go:generate go run .

package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"

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

type MethodType struct {
	Script string
}

const daoTemplateCode = `
package {{.Package}}

import (
	"github.com/jinzhu/gorm"

	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/user/{{.Package}}"
	{{.Package}}Repository "github.com/game-core/gocrafter/domain/repository/user/{{.Package}}"
)

type {{.CamelName}}Dao struct {
	ShardConn *database.ShardConn
}

func New{{.Name}}Dao(conn *database.SqlHandler) {{.Package}}Repository.{{.Name}}Repository {
	return &{{.CamelName}}Dao{
		ShardConn: conn.User,
	}
}

{{range $methodName, $MethodType := .Methods }}
	{{.Script}}
{{end}}
`

func generateDao(yamlFilePath string, outputBaseDir string) error {
	structInfo, err := getStructInfo(yamlFilePath)
	if err != nil {
		return err
	}

	outputDir := filepath.Join(outputBaseDir, structInfo.Package)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating output directory %s: %v", outputDir, err)
	}

	outputFileName := filepath.Join(outputDir, fmt.Sprintf("%s_dao.gen.go", transform.KebabToCamel(structInfo.Name)))
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("outputFileName file %s create error: %v", outputFileName, err)
	}

	if err := generateTemplate(structInfo, outputFile); err != nil {
		return fmt.Errorf("template error: %v", err)
	}

	fmt.Printf("Created %s Dao in %s\n", structInfo.Name, outputFileName)

	return nil
}

func generateTemplate(structInfo *StructInfo, outputFile *os.File) error {
	tmpl, err := template.New("daoTemplate").Parse(daoTemplateCode)
	if err != nil {
		return fmt.Errorf("error parsing DAO template: %v", err)
	}

	data := struct {
		Name      string
		CamelName string
		Package   string
		Methods   map[string]MethodType
	}{
		Name:      structInfo.Name,
		CamelName: transform.KebabToCamel(structInfo.Name),
		Package:   structInfo.Package,
		Methods:   generateMethods(structInfo),
	}

	if err := tmpl.ExecuteTemplate(outputFile, "daoTemplate", data); err != nil {
		return fmt.Errorf("faild to generateTemplate: %v", err)
	}

	return nil
}

func generateMethods(structInfo *StructInfo) map[string]MethodType {
	methods := make(map[string]MethodType)

	// FindByID
	if len(structInfo.Primary) > 0 {
		methods["FindByID"] = MethodType{
			Script: generateFindByID(structInfo),
		}
	}

	// FindOrNilByID
	if len(structInfo.Primary) > 0 {
		methods["FindOrNilByID"] = MethodType{
			Script: generateFindOrNilByID(structInfo),
		}
	}

	// FindByIndex
	for _, index := range structInfo.Index {
		indexFields := strings.Split(index, ",")
		methods[fmt.Sprintf("FindBy%s", strings.Join(indexFields, "And"))] = MethodType{
			Script: generateFindByIndex(structInfo, indexFields),
		}
	}

	// FindOrNilByIndex
	for _, index := range structInfo.Index {
		indexFields := strings.Split(index, ",")
		methods[fmt.Sprintf("FindOrNilBy%s", strings.Join(indexFields, "And"))] = MethodType{
			Script: generateFindOrNilByIndex(structInfo, indexFields),
		}
	}

	// List
	methods["List"] = MethodType{
		Script: generateList(structInfo),
	}

	// ListByIndex
	for _, index := range structInfo.Index {
		indexFields := strings.Split(index, ",")
		methods[fmt.Sprintf("ListBy%s", strings.Join(indexFields, "And"))] = MethodType{
			Script: generateListByIndex(structInfo, indexFields),
		}
	}

	// Create
	methods["Create"] = MethodType{
		Script: generateCreate(structInfo),
	}

	// Update
	methods["Update"] = MethodType{
		Script: generateUpdate(structInfo),
	}

	// Delete
	methods["Delete"] = MethodType{
		Script: generateDelete(structInfo),
	}

	return methods
}

func generateFindByID(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`func (d *%sDao) FindByID(ID int64, shardKey int) (*%s.%s, error) {
			entity := &%s.%s{}
			res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Find(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`,
		transform.KebabToCamel(structInfo.Name),
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
	)
}

func generateFindOrNilByID(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`func (d *%sDao) FindOrNilByID(ID int64, shardKey int) (*%s.%s, error) {
			entity := &%s.%s{}
			res := d.ShardConn.Shards[shardKey].ReadConn.Where("id = ?", ID).Find(entity)
			if res.RecordNotFound() {
				return nil, nil
			}
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`,
		transform.KebabToCamel(structInfo.Name),
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
	)
}

func generateFindByIndex(structInfo *StructInfo, indexFields []string) string {
	params := make([]struct{ Name, Type string }, len(indexFields))
	paramStrings := make([]string, len(indexFields))
	scriptStrings := make([]string, len(indexFields))

	for i, field := range indexFields {
		params[i] = struct{ Name, Type string }{field, structInfo.Fields[field].Type}
		paramStrings[i] = fmt.Sprintf("%s %s", field, structInfo.Fields[field].Type)
		scriptStrings[i] = fmt.Sprintf("Where(\"%s = ?\", %s)", structInfo.Fields[field].Name, field)
	}

	return fmt.Sprintf(
		`func (d *%sDao) FindBy%s(%s, shardKey int) (*%s.%s, error) {
			entity := &%s.%s{}
			res := d.ShardConn.Shards[shardKey].ReadConn.%s.Find(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`,
		transform.KebabToCamel(structInfo.Name),
		strings.Join(indexFields, "And"),
		strings.Join(paramStrings, ","),
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
		strings.Join(scriptStrings, "."),
	)
}

func generateFindOrNilByIndex(structInfo *StructInfo, indexFields []string) string {
	params := make([]struct{ Name, Type string }, len(indexFields))
	paramStrings := make([]string, len(indexFields))
	scriptStrings := make([]string, len(indexFields))

	for i, field := range indexFields {
		params[i] = struct{ Name, Type string }{field, structInfo.Fields[field].Type}
		paramStrings[i] = fmt.Sprintf("%s %s", field, structInfo.Fields[field].Type)
		scriptStrings[i] = fmt.Sprintf("Where(\"%s = ?\", %s)", structInfo.Fields[field].Name, field)
	}

	return fmt.Sprintf(
		`func (d *%sDao) FindOrNilBy%s(%s, shardKey int) (*%s.%s, error) {
			entity := &%s.%s{}
			res := d.ShardConn.Shards[shardKey].ReadConn.%s.Find(entity)
			if res.RecordNotFound() {
				return nil, nil
			}
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`,
		transform.KebabToCamel(structInfo.Name),
		strings.Join(indexFields, "And"),
		strings.Join(paramStrings, ","),
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
		strings.Join(scriptStrings, "."),
	)
}

func generateList(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`func (d *%sDao) List(limit int64, shardKey int) (*%s.%s, error) {
			entity := &%s.%s{}
			res := d.ShardConn.Shards[shardKey].ReadConn.Limit(limit).Find(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`,
		transform.KebabToCamel(structInfo.Name),
		structInfo.Package,
		transform.SingularToPlural(structInfo.Name),
		structInfo.Package,
		transform.SingularToPlural(structInfo.Name),
	)
}

func generateListByIndex(structInfo *StructInfo, indexFields []string) string {
	params := make([]struct{ Name, Type string }, len(indexFields))
	paramStrings := make([]string, len(indexFields))
	scriptStrings := make([]string, len(indexFields))

	for i, field := range indexFields {
		params[i] = struct{ Name, Type string }{field, structInfo.Fields[field].Type}
		paramStrings[i] = fmt.Sprintf("%s %s", field, structInfo.Fields[field].Type)
		scriptStrings[i] = fmt.Sprintf("Where(\"%s = ?\", %s)", structInfo.Fields[field].Name, field)
	}

	return fmt.Sprintf(
		`func (d *%sDao) ListBy%s(%s, shardKey int) (*%s.%s, error) {
			entity := &%s.%s{}
			res := d.ShardConn.Shards[shardKey].ReadConn.%s.Find(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`,
		transform.KebabToCamel(structInfo.Name),
		strings.Join(indexFields, "And"),
		strings.Join(paramStrings, ","),
		structInfo.Package,
		transform.SingularToPlural(structInfo.Name),
		structInfo.Package,
		transform.SingularToPlural(structInfo.Name),
		strings.Join(scriptStrings, "."),
	)
}

func generateCreate(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`func (d *%sDao) Create(entity *%s.%s, shardKey int, tx *gorm.DB) (*%s.%s, error) {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = d.ShardConn.Shards[shardKey].WriteConn
			}
		
			res := conn.Model(&%s.%s{}).Create(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`,
		transform.KebabToCamel(structInfo.Name),
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
	)
}

func generateUpdate(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`func (d *%sDao) Update(entity *%s.%s, shardKey int, tx *gorm.DB) (*%s.%s, error) {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = d.ShardConn.Shards[shardKey].WriteConn
			}
		
			res := conn.Model(&%s.%s{}).Where("id = ?", entity.ID).Update(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`,
		transform.KebabToCamel(structInfo.Name),
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
	)
}

func generateDelete(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`func (d *%sDao) Delete(entity *%s.%s, shardKey int, tx *gorm.DB) error {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = d.ShardConn.Shards[shardKey].WriteConn
			}
		
			res := conn.Model(&%s.%s{}).Where("id = ?", entity.ID).Delete(entity)
			if err := res.Error; err != nil {
				return err
			}
		
			return nil
		}
		`,
		transform.KebabToCamel(structInfo.Name),
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
	)
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

func main() {
	outputBaseDir := "../../../../infra/dao/user"
	yamlFiles, err := filepath.Glob("../../../../docs/entity/user/*.yaml")
	if err != nil {
		log.Fatalf("Error finding YAML files: %v", err)
	}

	for _, yamlFile := range yamlFiles {
		err := generateDao(yamlFile, outputBaseDir)
		if err != nil {
			log.Printf("Error generating dao from YAML file %s: %v", yamlFile, err)
		}
	}
}
