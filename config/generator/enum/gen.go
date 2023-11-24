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
	Name    string   `yaml:"name"`
	Package string   `yaml:"package"`
	Fields  []string `yaml:"structure"`
}

const templateCode = `
package {{.Package}}

type {{.Name}} int

{{.Script}}

`

func generateEnum(yamlFile string, outputBaseDir string) error {
	structInfo, err := getStructInfo(yamlFile)
	if err != nil {
		return err
	}

	outputDir := filepath.Join(outputBaseDir, structInfo.Package)
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("error creating output directory %s: %v", outputDir, err)
	}

	outputFileName := filepath.Join(outputDir, fmt.Sprintf("%s_enum.gen.go", transform.KebabToCamel(structInfo.Name)))
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		return fmt.Errorf("outputFileName file %s create error: %v", outputFileName, err)
	}

	if err := generateTemplate(structInfo, outputFile); err != nil {
		return fmt.Errorf("faild to generateTemplate: %v", err)
	}

	fmt.Printf("Created %s Enum in %s\n", structInfo.Name, outputFileName)

	return nil
}

func generateTemplate(structInfo *StructInfo, outputFile *os.File) error {
	tmpl, err := template.New("structTemplate").Parse(templateCode)
	if err != nil {
		return err
	}

	constList := strings.Split(structInfo.Fields[0], ",")
	constList[0] = fmt.Sprintf("%s %s = iota", constList[0], structInfo.Name)
	constBlock := fmt.Sprintf("\n\t\t%s\n\t", strings.Join(constList, "\n\t\t"))

	namesList := strings.Split(structInfo.Fields[0], ",")

	script := fmt.Sprintf(
		`
		const (
			%s
		)
		
		func (e %s) ToString() string {
			names := [...]string{%s}
			if e < %s || e > %s {
				return "Unknown"
			}

			return names[e]
		}
		`,
		constBlock,
		structInfo.Name,
		`"`+strings.Join(namesList, `","`)+`"`,
		namesList[0],
		namesList[len(namesList)-1],
	)

	if err := tmpl.ExecuteTemplate(outputFile, "structTemplate", struct {
		Name    string
		Package string
		Script  string
	}{
		Name:    structInfo.Name,
		Package: structInfo.Package,
		Script:  script,
	}); err != nil {
		return err
	}

	return nil
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
	outputBaseDir := "../../../domain/enum"
	yamlFiles, err := filepath.Glob("../../../docs/enum/*.yaml")
	if err != nil {
		log.Fatalf("Error finding YAML files: %v", err)
	}

	for _, yamlFile := range yamlFiles {
		err := generateEnum(yamlFile, outputBaseDir)
		if err != nil {
			log.Printf("Error generating enum from YAML file %s: %v", yamlFile, err)
		}
	}
}
