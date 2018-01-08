package format

import (
	"log"
	"strings"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"encoding/json"
)

var JsonMimeTypes = [5]string {
	"application/vnd.cryptorest+json",
	"application/x-json",
	"application/json",
	"text/x-json",
	"text/json",
}

var jsonExtensions = [1]string {
	"json",
}

func InputJsonFile(s *InputStructure) {
	cFile, err := ioutil.ReadFile(s.ConfigFile)
	if err != nil {
		log.Fatalf("JSON error: #%v ", err)
	}

	err =json.Unmarshal(cFile, &s)
	if err != nil {
		log.Fatalf("Unmarshal JSON: %v", err)
	}
}

func InitJsonInputData(s *InputStructure) {
	if s.ConfigFile == "" {
		return
	}

	switch strings.Trim(filepath.Ext(s.ConfigFile), ".") {
	case jsonExtensions[0]:
		InputJsonFile(&*s)
	}
}

func InputJson(s *InputStructure, hr bool) (string, error) {
	return "", nil
}

func OutputJson(w http.ResponseWriter, s *OutputStructure, hr bool) error {
	var b []byte
	var err error

	if hr {
		b, err = json.MarshalIndent(&s, HumanReadablePrefix, HumanReadableIndent)
	} else {
		b, err = json.Marshal(&s)
	}

	if err == nil {
		_, err = w.Write(b)
	}

	return err
}
