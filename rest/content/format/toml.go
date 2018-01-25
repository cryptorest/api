package format

import (
	"io"

	"github.com/BurntSushi/toml"
)

var Toml = Structure {
	&TomlHttpMimeTypes,
	&TomlFileExtensions,
	InputTomlFile,
	InputToml,
	nil,
	OutputToml,
}

var TomlHttpMimeTypes = []string {
	// RFC 2046
	"text/toml",
	// RFC 2046
	"application/toml",
	"application/x-toml",
}

var TomlFileExtensions = []string {
	".tml",
	".toml",
}

func InputTomlFile(s *InputStructure) error {
	return nil
}

func InputToml(b []byte, s *struct{}) error {
	return nil
}

func OutputToml(w io.Writer, s *OutputStructure, hr bool) error {
	hr = true

	err := toml.NewEncoder(w).Encode(&*s)

	return err
}
