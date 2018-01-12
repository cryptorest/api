package config

import (
	"log"
	"io/ioutil"
	"path/filepath"

	"encoding/xml"
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

	err = yaml.Unmarshal(cFile, &*c)

	if err != nil {
		log.Fatalf("Unmarshal YAML: %v", err)
	}
}

func InitTomlFile(c *Structure) {
	_, err := toml.DecodeFile(c.ConfigFile, &*c)

	if err != nil {
		log.Fatalf("Unmarshal TOML: %v", err)
	}
}

func InitJsonFile(c *Structure) {
	cFile, err := ioutil.ReadFile(c.ConfigFile)

	if err != nil {
		log.Fatalf("JSON error: #%v ", err)
	}

	err = json.Unmarshal(cFile, &*c)

	if err != nil {
		log.Fatalf("Unmarshal JSON: %v", err)
	}
}

func InitXmlFile(c *Structure) {
	cFile, err := ioutil.ReadFile(c.ConfigFile)

	if err != nil {
		log.Fatalf("XML error: #%v ", err)
	}

	err = xml.Unmarshal(cFile, &*c)

	if err != nil {
		log.Fatalf("Unmarshal XML: %v", err)
	}
}

func InitFile(c *Structure) {
	if c.ConfigFile == "" {
		return
	}

	switch filepath.Ext(c.ConfigFile) {
	case format.YamlFileExtensions[0], format.YamlFileExtensions[1]:
		InitYamlFile(&*c)
	case format.TomlFileExtensions[0], format.TomlFileExtensions[1]:
		InitTomlFile(&*c)
	case format.JsonFileExtensions[0]:
		InitJsonFile(&*c)
	case format.XmlFileExtensions[0]:
		InitXmlFile(&*c)
	}
}
