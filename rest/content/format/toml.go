package format

import (
	"io"

	"github.com/BurntSushi/toml"
)

var TomlHttpMimeTypes = [5]string {
	"text/toml",
	"text/x-toml",
	"application/toml",
	"application/x-toml",
	"application/vnd.cryptorest+toml",
}

var TomlFileExtensions = [2]string {
	".tml",
	".toml",
}

func InputToml(w io.Reader, s *InputStructure, hr bool) error {
	return nil
}

func OutputToml(w io.Writer, s *OutputStructure, hr bool) error {
	hr = true

	err := toml.NewEncoder(w).Encode(&s)

	return err
}
