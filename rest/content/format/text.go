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

	_, err := w.Write([]byte(s.Content))

	return err
}
