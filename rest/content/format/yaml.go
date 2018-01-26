package format

import (
	"io"

	"gopkg.in/yaml.v2"
)

var Yaml = Structure {
	&YamlHttpMimeTypes,
	&YamlFileExtensions,
	InputYamlFile,
	InputYaml,
	nil,
	OutputYaml,
}

var YamlHttpMimeTypes = []string {
	// RFC 2046
	"text/yaml",
	// RFC 2046
	"application/yaml",
	"application/x-yaml",
}

var YamlFileExtensions = []string {
	".yml",
	".yaml",
}

func InputYamlFile(s *InputStructure) error {
	return nil
}

func InputYaml(b []byte, f *Format) error {
	return nil
}

func OutputYaml(w io.Writer, s *OutputStructure, hr bool) error {
	hr = true
	var err error
	var b   []byte

	b, err = yaml.Marshal(&*s)

	if err == nil {
		_, err = w.Write(b)
	}

	return err
}
