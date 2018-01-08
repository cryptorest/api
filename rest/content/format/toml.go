package format

import (
	"net/http"

	"github.com/BurntSushi/toml"
)

var TomlMimeTypes = [5]string {
	"application/vnd.cryptorest+toml",
	"application/x-toml",
	"application/toml",
	"text/x-toml",
	"text/toml",
}

var TomlExtensions = [2]string {
	"tml",
	"toml",
}

func InputToml(s *InputStructure, hr bool) (string, error) {
	return "", nil
}

func OutputToml(w http.ResponseWriter, s *OutputStructure, hr bool) error {
	hr = true

	err := toml.NewEncoder(w).Encode(&s)

	return err
}
