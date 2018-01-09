package format

import (
	"io"

	"gopkg.in/yaml.v2"
)

var YamlHttpMimeTypes = [5]string {
	"application/vnd.cryptorest+yaml",
	"application/x-yaml",
	"application/yaml",
	"text/x-yaml",
	"text/yaml",
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
