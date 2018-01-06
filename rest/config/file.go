package config

import (
	"log"
	"strings"
	"io/ioutil"
	"path/filepath"

	"encoding/json"
	"gopkg.in/yaml.v2"
	"github.com/BurntSushi/toml"
)

var extensions = [5]string {
	"yml",
	"yaml",
	"tml",
	"toml",
	"json",
}

func InitFileYAML(c *Configuration) {
	cFile, err := ioutil.ReadFile(c.ConfigFile)
	if err != nil {
		log.Fatalf("YAML error: #%v ", err)
	}

	err = yaml.Unmarshal(cFile, c)
	if err != nil {
		log.Fatalf("Unmarshal YAML: %v", err)
	}
}

func InitFileTOML(c *Configuration) {
	_, err := toml.DecodeFile(c.ConfigFile, &c)

	if err != nil {
		log.Fatalf("Unmarshal TOML: %v", err)
	}
}

func InitFileJSON(c *Configuration) {
	cFile, err := ioutil.ReadFile(c.ConfigFile)
	if err != nil {
		log.Fatalf("JSON error: #%v ", err)
	}

	err =json.Unmarshal(cFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal JSON: %v", err)
	}
}

func InitFile(c *Configuration) {
	if c.ConfigFile == "" {
		return
	}

	switch strings.Trim(filepath.Ext(c.ConfigFile), ".") {
	case extensions[0], extensions[1]:
		InitFileYAML(&*c)
	case extensions[2], extensions[3]:
		InitFileTOML(&*c)
	case extensions[4]:
		InitFileJSON(&*c)
	}
}
