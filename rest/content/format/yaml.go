package format

import (
	"log"
	"strings"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var YamlMimeTypes = [5]string {
	"application/vnd.cryptorest+yaml",
	"application/x-yaml",
	"application/yaml",
	"text/x-yaml",
	"text/yaml",
}

var yamlExtensions = [2]string {
	"yml",
	"yaml",
}

func YamlInputData(c *InputStructure) {
	cFile, err := ioutil.ReadFile(c.ConfigFile)
	if err != nil {
		log.Fatalf("YAML error: #%v ", err)
	}

	err = yaml.Unmarshal(cFile, c)
	if err != nil {
		log.Fatalf("Unmarshal YAML: %v", err)
	}
}

func InitYamlInputData(c *InputStructure) {
	if c.ConfigFile == "" {
		return
	}

	switch strings.Trim(filepath.Ext(c.ConfigFile), ".") {
	case yamlExtensions[0], yamlExtensions[1]:
		YamlInputData(&*c)

	}
}
