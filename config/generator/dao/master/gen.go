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
	Table   string                 `yaml:"table"`
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
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/patrickmn/go-cache"
	
	"github.com/game-core/gocrafter/config/database"
	"github.com/game-core/gocrafter/domain/entity/master/{{.Package}}"
	{{.Package}}Repository "github.com/game-core/gocrafter/domain/repository/master/{{.Package}}"
)

type {{.Package}}Dao struct {
	Read  *gorm.DB
	Write *gorm.DB
	Cache *cache.Cache
}

func New{{.Name}}Dao(conn *database.SqlHandler) {{.Package}}Repository.{{.Name}}Repository {
	return &{{.Package}}Dao{
		Read:  conn.Master.ReadConn,
		Write: conn.Master.WriteConn,
		Cache: cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

{{range $methodName, $MethodType := .Methods }}
	{{.Script}}
{{end}}

func cacheKey(method string, key string) string {
	return fmt.Sprintf("{{.Table}}:%s:%s", method, key)
}
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

	outputFileName := filepath.Join(outputDir, fmt.Sprintf("%s_dao.gen.go", structInfo.Package))
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("outputFileName file %s create error: %v", outputFileName, err)
	}
	defer outputFile.Close()

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
		Name    string
		Package string
		Table   string
		Methods map[string]MethodType
	}{
		Name:    structInfo.Name,
		Package: structInfo.Package,
		Table:   structInfo.Table,
		Methods: generateMethods(structInfo),
	}

	if err := tmpl.ExecuteTemplate(outputFile, "daoTemplate", data); err != nil {
		return fmt.Errorf("template error: %v", err)
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

	// FindByIndex
	for _, index := range structInfo.Index {
		indexFields := strings.Split(index, ",")
		methods[fmt.Sprintf("FindBy%s", strings.Join(indexFields, "And"))] = MethodType{
			Script: generateFindByIndex(structInfo, indexFields),
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
		`func (d *%sDao) FindByID(ID int64) (*%s.%s, error) {
			cachedResult, found := d.Cache.Get(cacheKey("FindByID", %s))
			if found {
				if cachedEntity, ok := cachedResult.(*%s.%s); ok {
					return cachedEntity, nil
				}
			}

			entity := &%s.%s{}
			res := d.Read.Where("id = ?", ID).Find(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			d.Cache.Set(cacheKey("FindByID", %s), entity, cache.DefaultExpiration)

			return entity, nil
		}
		`,
		structInfo.Package,
		structInfo.Package,
		structInfo.Name,
		`fmt.Sprintf("%d_", ID)`,
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
		`fmt.Sprintf("%d_", ID)`,
	)
}

func generateFindByIndex(structInfo *StructInfo, indexFields []string) string {
	params := make([]struct{ Name, Type string }, len(indexFields))
	paramStrings := make([]string, len(indexFields))
	scriptStrings := make([]string, len(indexFields))
	sprints := make([]string, len(indexFields))
	sprintParams := make([]string, len(indexFields))

	for i, field := range indexFields {
		params[i] = struct{ Name, Type string }{field, structInfo.Fields[field].Type}
		paramStrings[i] = fmt.Sprintf("%s %s", field, structInfo.Fields[field].Type)
		scriptStrings[i] = fmt.Sprintf("Where(\"%s = ?\", %s)", structInfo.Fields[field].Name, field)
		sprintParams[i] = field

		switch structInfo.Fields[field].Type {
		case "string":
			sprints[i] = "%s_"
		default:
			sprints[i] = "%d_"
		}
	}

	return fmt.Sprintf(
		`func (d *%sDao) FindBy%s(%s) (*%s.%s, error) {
			cachedResult, found := d.Cache.Get(cacheKey("FindBy%s", %s))
			if found {
				if cachedEntity, ok := cachedResult.(*%s.%s); ok {
					return cachedEntity, nil
				}
			}

			entity := &%s.%s{}
			res := d.Read.%s.Find(entity)
			if err := res.Error; err != nil {
				return nil, err
			}

			d.Cache.Set(cacheKey("FindBy%s", %s), entity, cache.DefaultExpiration)
		
			return entity, nil
		}
		`,
		structInfo.Package,
		strings.Join(indexFields, "And"),
		strings.Join(paramStrings, ","),
		structInfo.Package,
		structInfo.Name,
		strings.Join(indexFields, "And"),
		fmt.Sprintf(`fmt.Sprintf("%s", %s)`, strings.Join(sprints, ""), strings.Join(sprintParams, ",")),
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
		strings.Join(scriptStrings, "."),
		strings.Join(indexFields, "And"),
		fmt.Sprintf(`fmt.Sprintf("%s", %s)`, strings.Join(sprints, ""), strings.Join(sprintParams, ",")),
	)
}

func generateList(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`func (d *%sDao) List(limit int64) (*%s.%ss, error) {
			cachedResult, found := d.Cache.Get(cacheKey("List", ""))
			if found {
				if cachedEntity, ok := cachedResult.(*%s.%ss); ok {
					return cachedEntity, nil
				}
			}

			entity := &%s.%ss{}
			res := d.Read.Limit(limit).Find(entity)
			if err := res.Error; err != nil {
				return nil, err
			}

			d.Cache.Set(cacheKey("List", ""), entity, cache.DefaultExpiration)
		
			return entity, nil
		}
		`,
		structInfo.Package,
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
	)
}

func generateListByIndex(structInfo *StructInfo, indexFields []string) string {
	params := make([]struct{ Name, Type string }, len(indexFields))
	paramStrings := make([]string, len(indexFields))
	scriptStrings := make([]string, len(indexFields))
	sprints := make([]string, len(indexFields))
	sprintParams := make([]string, len(indexFields))

	for i, field := range indexFields {
		params[i] = struct{ Name, Type string }{field, structInfo.Fields[field].Type}
		paramStrings[i] = fmt.Sprintf("%s %s", field, structInfo.Fields[field].Type)
		scriptStrings[i] = fmt.Sprintf("Where(\"%s = ?\", %s)", structInfo.Fields[field].Name, field)
		sprintParams[i] = field

		switch structInfo.Fields[field].Type {
		case "string":
			sprints[i] = "%s_"
		default:
			sprints[i] = "%d_"
		}
	}

	return fmt.Sprintf(
		`func (d *%sDao) ListBy%s(%s) (*%s.%ss, error) {
			cachedResult, found := d.Cache.Get(cacheKey("ListBy%s", %s))
			if found {
				if cachedEntity, ok := cachedResult.(*%s.%ss); ok {
					return cachedEntity, nil
				}
			}

			entity := &%s.%ss{}
			res := d.Read.%s.Find(entity)
			if err := res.Error; err != nil {
				return nil, err
			}

			d.Cache.Set(cacheKey("ListBy%s", %s), entity, cache.DefaultExpiration)
		
			return entity, nil
		}
		`,
		structInfo.Package,
		strings.Join(indexFields, "And"),
		strings.Join(paramStrings, ","),
		structInfo.Package,
		structInfo.Name,
		strings.Join(indexFields, "And"),
		fmt.Sprintf(`fmt.Sprintf("%s", %s)`, strings.Join(sprints, ""), strings.Join(sprintParams, ",")),
		structInfo.Package,
		structInfo.Name,
		structInfo.Package,
		structInfo.Name,
		strings.Join(scriptStrings, "."),
		strings.Join(indexFields, "And"),
		fmt.Sprintf(`fmt.Sprintf("%s", %s)`, strings.Join(sprints, ""), strings.Join(sprintParams, ",")),
	)
}

func generateCreate(structInfo *StructInfo) string {
	return fmt.Sprintf(
		`func (d *%sDao) Create(entity *%s.%s, tx *gorm.DB) (*%s.%s, error) {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = d.Write
			}
		
			res := conn.Model(&%s.%s{}).Create(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`,
		structInfo.Package,
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
		`func (d *%sDao) Update(entity *%s.%s, tx *gorm.DB) (*%s.%s, error) {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = d.Write
			}
		
			res := conn.Model(&%s.%s{}).Where("id = ?", entity.ID).Update(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`,
		structInfo.Package,
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
		`func (d *%sDao) Delete(entity *%s.%s, tx *gorm.DB) error {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = d.Write
			}
		
			res := conn.Model(&%s.%s{}).Where("id = ?", entity.ID).Delete(entity)
			if err := res.Error; err != nil {
				return err
			}
		
			return nil
		}
		`,
		structInfo.Package,
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
	outputBaseDir := "../../../../infra/dao/master"
	yamlFiles, err := filepath.Glob("../../../../docs/entity/master/*.yaml")
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
