package format

import (
	"io"
	"log"
	"io/ioutil"

	"encoding/json"
)

var Json = Structure {
	&JsonHttpMimeTypes,
	&JsonFileExtensions,
	InputJsonFile,
	InputJson,
	OutputJson,
}

var JsonHttpMimeTypes = []string {
	// RFC 2046, for human readable mode
	"text/json",
	// RFC 8259
	"application/json",
}

var JsonFileExtensions = []string {
	".json",
}

func InputJsonFile(s *InputStructure) error {
	cFile, err := ioutil.ReadFile(s.File)

	if err != nil {
		log.Fatalf("JSON error: %v ", err)
	}

	err = json.Unmarshal(cFile, &*s)

	if err != nil {
		log.Fatalf("Unmarshal JSON: %v", err)
	}

	return err
}

func InputJson(w io.Reader, s *InputStructure, hr bool) error {
	return nil
}

func OutputJson(w io.Writer, s *OutputStructure, hr bool) error {
	var err error
	var b   []byte

	if hr {
		b, err = json.MarshalIndent(&*s, HumanReadablePrefix, HumanReadableIndent)
	} else {
		b, err = json.Marshal(&*s)
	}

	if err == nil {
		_, err = w.Write(b)
	}

	return err
}
