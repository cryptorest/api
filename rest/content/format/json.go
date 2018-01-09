package format

import (
	"io"

	"encoding/json"
)

var JsonHttpMimeTypes = [2]string {
	// RFC 2046, for human readable mode
	"text/json",
	// RFC8259
	"application/json",
}

var JsonFileExtensions = [1]string {
	".json",
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
