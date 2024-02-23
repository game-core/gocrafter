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
	Name    string `yaml:"name"`
	Number  int    `yaml:"number"`
	Comment string `yaml:"comment"`
}

type TemplateStruct struct {
	Package string
	Name    string
	Comment string
	Script  string
}

func main() {
	yamls := "../../../../../docs/yaml/pkg/domain/enum"
	base := "../../../../../pkg/domain/enum"

	if err := filepath.Walk(yamls, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("error accessing path %s: %v", path, err)
			return nil
		}

		if !info.IsDir() && strings.HasSuffix(info.Name(), ".yaml") {
			if err := NewEnum().generate(path, base); err != nil {
				log.Printf("failed to NewEnum().generate: %s", err)
			}
		}

		return nil
	}); err != nil {
		log.Fatalf("failed to filepath.Walk: %s", err)
	}
}
