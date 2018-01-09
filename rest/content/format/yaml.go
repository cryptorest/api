package format

import (
	"io"

	"gopkg.in/yaml.v2"
)

var YamlHttpMimeTypes = [2]string {
	// RFC 2046
	"text/yaml",
	// RFC 2046
	"application/yaml",
}

var YamlFileExtensions = [2]string {
	".yml",
	".yaml",
}

func InputYaml(w io.Reader, s *InputStructure, hr bool) error {
	return nil
}

func OutputYaml(w io.Writer, s *OutputStructure, hr bool) error {
	hr = true
	var err error
	var b   []byte

	b, err = yaml.Marshal(&s)

	if err == nil {
		_, err = w.Write(b)
	}

	return err
}
