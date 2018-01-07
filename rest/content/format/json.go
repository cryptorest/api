package format

import (
	"log"
	"strings"
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

func JsonInputData(c *InputStructure) {
	cFile, err := ioutil.ReadFile(c.ConfigFile)
	if err != nil {
		log.Fatalf("JSON error: #%v ", err)
	}

	err =json.Unmarshal(cFile, &c)
	if err != nil {
		log.Fatalf("Unmarshal JSON: %v", err)
	}
}

func InitJsonInputData(c *InputStructure) {
	if c.ConfigFile == "" {
		return
	}

	switch strings.Trim(filepath.Ext(c.ConfigFile), ".") {
	case jsonExtensions[0]:
		JsonInputData(&*c)
	}
}
