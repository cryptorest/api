package format

import (
	"log"
	"strings"
	"path/filepath"
	"github.com/BurntSushi/toml"
)

var TomlMimeTypes = [5]string {
	"application/vnd.cryptorest+toml",
	"application/x-toml",
	"application/toml",
	"text/x-toml",
	"text/toml",
}

var tomlExtensions = [2]string {
	"tml",
	"toml",
}

func TomlInputData(c *InputStructure) {
	_, err := toml.DecodeFile(c.ConfigFile, &c)

	if err != nil {
		log.Fatalf("Unmarshal TOML: %v", err)
	}
}

func InitTomlInputData(c *InputStructure) {
	if c.ConfigFile == "" {
		return
	}

	switch strings.Trim(filepath.Ext(c.ConfigFile), ".") {
	case tomlExtensions[0], tomlExtensions[1]:
		TomlInputData(&*c)
	}
}
