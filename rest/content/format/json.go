package format

import (
	"io"

	"encoding/json"
)

var JsonHttpMimeTypes = [5]string {
	"application/vnd.cryptorest+json",
	"application/x-json",
	"application/json",
	"text/x-json",
	"text/json",
}

var JsonFileExtensions = [1]string {
	"json",
}

func InputJson(w io.Reader, s *InputStructure, hr bool) error {
	return nil
}

func OutputJson(w io.Writer, s *OutputStructure, hr bool) error {
	var err error
	var b   []byte

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
