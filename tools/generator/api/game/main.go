//go:generate go run .

package main

import (
	"log"
	"os"
	"path/filepath"
	"strings"
)

type YamlStruct struct {
	Name       string               `yaml:"name"`
	Package    string               `yaml:"package"`
	Comment    string               `yaml:"comment"`
	Structures map[string]Structure `yaml:"structure"`
}

type Structure struct {
	Name     string `yaml:"name"`
	Method   string `yaml:"method"`
	Auth     bool   `yaml:"auth"`
	Request  string `yaml:"request"`
	Response string `yaml:"response"`
	Type     string `yaml:"type"`
	Package  string `yaml:"package"`
	Nullable bool   `yaml:"nullable"`
	Number   int    `yaml:"number"`
	Comment  string `yaml:"comment"`
}

type TemplateStruct struct {
	Package string
	Name    string
	Comment string
	Script  string
	Import  string
}

var importCode = ""

func main() {
	yamls := "../../../../docs/yaml/api/game"
	base := "../../../../docs/proto/api/game"

	if err := filepath.Walk(yamls, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("error accessing path %s: %v", path, err)
			return nil
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".yaml") {
			if err := NewProto().generate(path, base); err != nil {
				log.Printf("failed to NewProto().generate: %s", err)
			}
		}

		return nil
	}); err != nil {
		log.Fatalf("failed to filepath.Walk: %s", err)
	}
}
