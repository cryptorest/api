package format

import (
	"log"
	"strings"
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

func TextInputData(c *InputStructure) {
	_, err := toml.DecodeFile(c.ConfigFile, &c)

	if err != nil {
		log.Fatalf("Unmarshal TEXT: %v", err)
	}
}

func InitTextInputData(c *InputStructure) {
	if c.ConfigFile == "" {
		return
	}

	switch strings.Trim(filepath.Ext(c.ConfigFile), ".") {
	case textExtensions[0]:
		TextInputData(&*c)
	}
}
