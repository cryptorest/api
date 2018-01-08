package format

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

var YamlMimeTypes = [5]string {
	"application/vnd.cryptorest+yaml",
	"application/x-yaml",
	"application/yaml",
	"text/x-yaml",
	"text/yaml",
}

var YamlExtensions = [2]string {
	"yml",
	"yaml",
}

func InputYaml(s *InputStructure, hr bool) (string, error) {
	return "", nil
}

func OutputYaml(w http.ResponseWriter, s *OutputStructure, hr bool) error {
	hr = true
	var err error
	var b []byte

	b, err = yaml.Marshal(&s)

	if err == nil {
		_, err = w.Write(b)
	}

	return err
}
