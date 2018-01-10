package format

import (
	"io"

	"github.com/BurntSushi/toml"
)

var TomlHttpMimeTypes = [2]string {
	// RFC 2046
	"text/toml",
	// RFC 2046
	"application/toml",
}

var TomlFileExtensions = []string {
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
