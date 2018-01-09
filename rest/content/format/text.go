package format

import (
	"io"
)

var TextHttpMimeTypes = [5]string {
	"text/plane",
	"text/x-plane",
	"application/text",
	"application/x-text",
	"application/vnd.cryptorest+text",
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
