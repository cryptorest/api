package format

import (
	"net/http"

	"encoding/json"
)

var JsonMimeTypes = [5]string {
	"application/vnd.cryptorest+json",
	"application/x-json",
	"application/json",
	"text/x-json",
	"text/json",
}

var JsonExtensions = [1]string {
	"json",
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
