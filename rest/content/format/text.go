package format

import (
	"log"
	"strings"
	"net/http"
	"path/filepath"

	"github.com/BurntSushi/toml"
)

var TextMimeTypes = [5]string {
	"application/vnd.cryptorest+text",
	"application/x-text",
	"application/text",
	"text/x-plane",
	"text/plane",
}

var textExtensions = [1]string {
	"txt",
}

func TextInputData(s *InputStructure) {
	_, err := toml.DecodeFile(s.ConfigFile, &s)

	if err != nil {
		log.Fatalf("Unmarshal TEXT: %v", err)
	}
}

func InitTextInputData(s *InputStructure) {
	if s.ConfigFile == "" {
		return
	}

	switch strings.Trim(filepath.Ext(s.ConfigFile), ".") {
	case textExtensions[0]:
		TextInputData(&*s)
	}
}

func InputText(s *InputStructure, hr bool) (string, error) {
	return "", nil
}

func OutputText(w http.ResponseWriter, s *OutputStructure, hr bool) error {
	hr      = true
	s.Error = ""

	_, err := w.Write([]byte(s.Content))

	return err
}
