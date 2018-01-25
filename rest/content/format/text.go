package format

import (
	"io"
)

var Text = Structure {
	&TextHttpMimeTypes,
	&TextFileExtensions,
	InputTextFile,
	InputText,
	nil,
	OutputText,
}

var TextHttpMimeTypes = []string {
	// RFC2046, RFC3676, RFC5147
	"text/plane",
}

var TextFileExtensions = []string {
	".txt",
	".text",
}

func InputTextFile(s *InputStructure) error {
	return nil
}

func InputText(b []byte, s *struct{}) error {
	return nil
}

func OutputText(w io.Writer, s *OutputStructure, hr bool) error {
	hr = false
	var data []byte

	if s.Status.Message == "" {
		data = []byte(s.Content)
	} else {
		data = []byte(s.Status.Message)
	}

	_, err := w.Write(data)

	return err
}
