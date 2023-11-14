//go:generate go run .

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"text/template"
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
	Database string                 `yaml:"database"`
	Package  string                 `yaml:"package"`
	Fields   map[string]StructField `yaml:"structure"`
	Primary  []string               `yaml:"primary"`
	Index    []string               `yaml:"index"`
}

type methodType struct {
	Script  string
}

const daoTemplateCode = `
package {{.Package}}

import (
	"github.com/jinzhu/gorm"
	"github.com/architecture-template/echo-ddd/config/database"
	"github.com/architecture-template/echo-ddd/domain/model/{{.Database}}/{{.Package}}"
	{{.RepositoryImportPath}}
)

type {{.Package}}Dao struct {
	Read  *gorm.DB
	Write *gorm.DB
}

func New{{.Name}}Dao(conn *database.SqlHandler) {{.Package}}Repository.{{.RepositoryInterface}} {
	return &{{.Package}}Dao{
		Read:  conn.User.ReadConn,
		Write: conn.User.WriteConn,
	}
}

{{range $methodName, $methodType := .Methods }}
	{{.Script}}
{{end}}
`

func generateDao(yamlFilePath string, outputBaseDir string) error {
	yamlData, err := ioutil.ReadFile(yamlFilePath)
	if err != nil {
		return fmt.Errorf("error reading YAML file %s: %v", yamlFilePath, err)
	}

	var structInfo StructInfo
	err = yaml.Unmarshal(yamlData, &structInfo)
	if err != nil {
		return fmt.Errorf("error unmarshalling YAML in file %s: %v", yamlFilePath, err)
	}

	outputDir := filepath.Join(fmt.Sprintf("%s/%s", outputBaseDir, structInfo.Database), structInfo.Package)
	err = os.MkdirAll(outputDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("error creating output directory %s: %v", outputDir, err)
	}

	outputFileName := filepath.Join(outputDir, fmt.Sprintf("%s_dao.gen.go", structInfo.Package))
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("outputFileName file %s create error: %v", outputFileName, err)
	}
	defer outputFile.Close()

	methods := make(map[string]methodType)

	// FindByID
	if len(structInfo.Primary) > 0 {
		methods["FindByID"] = methodType{
			Script: fmt.Sprintf(`
			func (e *%sDao) FindByID(ID int64) (*%s.%s, error) {
				entity := &%s.%s{}
				res := e.Read.Where("id = ?", ID).Find(entity)
				if err := res.Error; err != nil {
					return nil, err
				}
			
				return entity, nil
			}
			`, structInfo.Package, structInfo.Package, structInfo.Name, structInfo.Package, structInfo.Name),
		}
	}

	// FindByIndex
	for _, index := range structInfo.Index {
		indexFields := strings.Split(index, ",")
		params := make([]struct{ Name, Type string }, len(indexFields))

		var paramStrings []string
		var scriptStrings []string

		for i, field := range indexFields {
			params[i] = struct{ Name, Type string }{field, structInfo.Fields[field].Type}
			paramStrings = append(paramStrings, fmt.Sprintf("%s %s", field, structInfo.Fields[field].Type))
			scriptStrings = append(scriptStrings, fmt.Sprintf("Where(\"%s = ?\", %s)", structInfo.Fields[field].Name, field))
		}

		methods[fmt.Sprintf("FindBy%s", strings.Join(indexFields, "And"))] = methodType{
			Script: fmt.Sprintf(`
			func (e *%sDao) FindBy%s(%s) (*%s.%s, error) {
				entity := &%s.%s{}
				res := e.Read.%s.Find(entity)
				if err := res.Error; err != nil {
					return nil, err
				}
			
				return entity, nil
			}
			`, structInfo.Package, strings.Join(indexFields, "And"), strings.Join(paramStrings, ","), structInfo.Package, structInfo.Name, structInfo.Package, structInfo.Name, strings.Join(scriptStrings, ".")),
		}
	}

	// List
	methods["List"] = methodType{
		Script: fmt.Sprintf(`
		func (e *%sDao) List(limit int64) (*%s.%ss, error) {
			entity := &%s.%ss{}
			res := e.Read.Limit(limit).Find(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`, structInfo.Package, structInfo.Package, structInfo.Name, structInfo.Package, structInfo.Name),
	}

	// Create
	methods["Create"] = methodType{
		Script: fmt.Sprintf(`
		func (e *%sDao) Create(entity *%s.%s, tx *gorm.DB) (*%s.%s, error) {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = e.Write
			}
		
			res := conn.Model(&%s.%s{}).Create(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`, structInfo.Package, structInfo.Package, structInfo.Name, structInfo.Package, structInfo.Name, structInfo.Package, structInfo.Name),
	}

	// Update
	methods["Update"] = methodType{
		Script: fmt.Sprintf(`
		func (e *%sDao) Update(entity *%s.%s, tx *gorm.DB) (*%s.%s, error) {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = e.Write
			}
		
			res := conn.Model(&%s.%s{}).Where("id = ?", entity.ID).Update(entity)
			if err := res.Error; err != nil {
				return nil, err
			}
		
			return entity, nil
		}
		`, structInfo.Package, structInfo.Package, structInfo.Name, structInfo.Package, structInfo.Name, structInfo.Package, structInfo.Name),
	}

	// Delete
	methods["Delete"] = methodType{
		Script: fmt.Sprintf(`
		func (e *%sDao) Delete(entity *%s.%s, tx *gorm.DB) error {
			var conn *gorm.DB
			if tx != nil {
				conn = tx
			} else {
				conn = e.Write
			}
		
			res := conn.Model(&%s.%s{}).Where("id = ?", entity.ID).Delete(entity)
			if err := res.Error; err != nil {
				return err
			}
		
			return nil
		}
		`, structInfo.Package, structInfo.Package, structInfo.Name, structInfo.Package, structInfo.Name),
	}

	daoTmpl, err := template.New("daoTemplate").Parse(daoTemplateCode)
	if err != nil {
		return fmt.Errorf("error parsing DAO template: %v", err)
	}

	err = daoTmpl.ExecuteTemplate(outputFile, "daoTemplate", struct {
		Name                 string
		Package              string
		Database             string
		Methods              map[string]methodType
		RepositoryImportPath string
		RepositoryInterface  string
	}{
		Name:                 structInfo.Name,
		Package:              structInfo.Package,
		Database:             structInfo.Database,
		Methods:              methods,
		RepositoryImportPath: fmt.Sprintf("%s \"github.com/architecture-template/echo-ddd/domain/repository/%s/%s\"", structInfo.Package+"Repository", structInfo.Database, structInfo.Package),
		RepositoryInterface:  fmt.Sprintf("%sRepository", structInfo.Name),
	})
	if err != nil {
		return fmt.Errorf("template error: %v", err)
	}

	fmt.Printf("Created %s Dao in %s\n", structInfo.Name, outputFileName)

	return nil
}

func main() {
	outputBaseDir := "../../infra/dao"
	yamlFiles, err := filepath.Glob("../../docs/entity/*.yaml")
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
