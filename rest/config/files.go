package config

import (
	"log"
	"strings"
	"io/ioutil"
	"path/filepath"

	"encoding/json"
	"gopkg.in/yaml.v2"
	"github.com/BurntSushi/toml"

	"rest/content/format"
)

func InitYamlFile(c *Structure) {
	cFile, err := ioutil.ReadFile(c.ConfigFile)

	if err != nil {
		log.Fatalf("YAML error: #%v ", err)
	}

	err = yaml.Unmarshal(cFile, c)

	if err != nil {
		log.Fatalf("Unmarshal YAML: %v", err)
	}
}

func InitTomlFile(c *Structure) {
	_, err := toml.DecodeFile(c.ConfigFile, &c)

	if err != nil {
		log.Fatalf("Unmarshal TOML: %v", err)
	}
}

func InitJsonFile(c *Structure) {
	cFile, err := ioutil.ReadFile(c.ConfigFile)

	if err != nil {
		log.Fatalf("JSON error: #%v ", err)
	}

	err =json.Unmarshal(cFile, &c)

	if err != nil {
		log.Fatalf("Unmarshal JSON: %v", err)
	}
}

func InitFile(c *Structure) {
	if c.ConfigFile == "" {
		return
	}

	switch strings.Trim(filepath.Ext(c.ConfigFile), ".") {
	case format.YamlExtensions[0], format.YamlExtensions[1]:
		InitYamlFile(&*c)
	case format.TomlExtensions[0], format.TomlExtensions[1]:
		InitTomlFile(&*c)
	case format.JsonExtensions[0]:
		InitJsonFile(&*c)
	}
}
