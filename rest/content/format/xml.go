package format

import (
	"io"

	"encoding/xml"
)

var XmlHttpMimeTypes = [2]string {
	// RFC 7303, for human readable mode
	"text/xml",
	// RFC RFC7303
	"application/xml",
}

var XmlFileExtensions = []string {
	".xml",
}

func InputXml(w io.Reader, s *InputStructure, hr bool) error {
	return nil
}

func OutputXml(w io.Writer, s *OutputStructure, hr bool) error {
	var err error
	var b   []byte

	if hr {
		b, err = xml.MarshalIndent(&s, HumanReadablePrefix, HumanReadableIndent)
	} else {
		b, err = xml.Marshal(&s)
	}

	if err == nil {
		_, err = w.Write(b)
	}

	return err
}
