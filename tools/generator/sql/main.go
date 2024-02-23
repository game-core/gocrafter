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
	Primary    []string             `yaml:"primary"`
	Unique     []string             `yaml:"unique"`
	Index      []string             `yaml:"index"`
}

type Structure struct {
	Name     string `yaml:"name"`
	Type     string `yaml:"type"`
	Package  string `yaml:"package"`
	Nullable bool   `yaml:"nullable"`
	Number   int    `yaml:"number"`
	Comment  string `yaml:"comment"`
}

type TemplateStruct struct {
	Package   string
	Name      string
	SnakeName string
	Comment   string
	Script    string
}

func generate(yamls, base string) error {
	if err := filepath.Walk(yamls, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Printf("error accessing path %s: %v", path, err)
			return nil
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".yaml") {
			if err := NewSql().generate(path, base); err != nil {
				log.Printf("failed to NewSql().generate: %s", err)
			}
		}
		return nil
	}); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := generate("../../../docs/yaml/pkg/infrastructure/mysql/common", "../../../docs/sql/common/ddl"); err != nil {
		log.Printf("failed to coomon: %s", err)
	}

	if err := generate("../../../docs/yaml/pkg/infrastructure/mysql/master", "../../../docs/sql/master/ddl"); err != nil {
		log.Printf("failed to master: %s", err)
	}

	if err := generate("../../../docs/yaml/pkg/infrastructure/mysql/user", "../../../docs/sql/user/ddl"); err != nil {
		log.Printf("failed to user: %s", err)
	}
}
