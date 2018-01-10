package format

import (
	"io"
)

var TextHttpMimeTypes = [2]string {
	// RFC2046, RFC3676, RFC5147
	"text/plane",
}

var TextFileExtensions = [1]string {
	".txt",
}

func InputText(w io.Reader, s *InputStructure, hr bool) error {
	return nil
}

func OutputText(w io.Writer, s *OutputStructure, hr bool) error {
	hr = false
	var data []byte

	if s.Error == "" {
		data = []byte(s.Content)
	} else {
		data = []byte(s.Error)
	}
	_, err := w.Write(data)

	return err
}
