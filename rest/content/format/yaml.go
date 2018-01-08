package format

import (
	"log"
	"strings"
	"net/http"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

var YamlMimeTypes = [5]string {
	"application/vnd.cryptorest+yaml",
	"application/x-yaml",
	"application/yaml",
	"text/x-yaml",
	"text/yaml",
}

var yamlExtensions = [2]string {
	"yml",
	"yaml",
}

func YamlInputData(s *InputStructure) {
	cFile, err := ioutil.ReadFile(s.ConfigFile)
	if err != nil {
		log.Fatalf("YAML error: #%v ", err)
	}

	err = yaml.Unmarshal(cFile, s)
	if err != nil {
		log.Fatalf("Unmarshal YAML: %v", err)
	}
}

func InitYamlInputData(s *InputStructure) {
	if s.ConfigFile == "" {
		return
	}

	switch strings.Trim(filepath.Ext(s.ConfigFile), ".") {
	case yamlExtensions[0], yamlExtensions[1]:
		YamlInputData(&*s)
	}
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
