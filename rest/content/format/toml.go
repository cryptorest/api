package format

import (
	"log"
	"strings"
	"net/http"
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

func TomlInputData(s *InputStructure) {
	_, err := toml.DecodeFile(s.ConfigFile, &s)

	if err != nil {
		log.Fatalf("Unmarshal TOML: %v", err)
	}
}

func InitTomlInputData(s *InputStructure) {
	if s.ConfigFile == "" {
		return
	}

	switch strings.Trim(filepath.Ext(s.ConfigFile), ".") {
	case tomlExtensions[0], tomlExtensions[1]:
		TomlInputData(&*s)
	}
}

func InputToml(s *InputStructure, hr bool) (string, error) {
	return "", nil
}

func OutputToml(w http.ResponseWriter, s *OutputStructure, hr bool) error {
	hr = true

	err := toml.NewEncoder(w).Encode(&s)

	return err
}
