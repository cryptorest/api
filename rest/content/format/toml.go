package format

import (
	"io"

	"github.com/BurntSushi/toml"
)

var TomlHttpMimeTypes = [5]string {
	"application/vnd.cryptorest+toml",
	"application/x-toml",
	"application/toml",
	"text/x-toml",
	"text/toml",
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
