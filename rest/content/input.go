package content

import (
	"io"
	"net/http"

	"rest/content/format"
)

var InputFormatFuncs = [4]func(w io.Reader, s *format.InputStructure, hr bool) error {
	format.InputText,
	format.InputJson,
	format.InputYaml,
	format.InputToml,
}

func InputHttpMimeType(r *http.Request) string {
	return r.Header.Get(MimeKeyRequest)
}

func DefaultInputFormat() (string, bool, func(w io.Reader, c *format.InputStructure, hr bool) error) {
	return format.TextHttpMimeTypes[4], false, format.InputText
}

type Intput struct {
	Format          func(w io.Reader, s *format.InputStructure, hr bool) error
	HttpMimeType    string
	Reader          *http.Request
	Structure       format.InputStructure
	IsHumanReadable bool
}

func DefaultInputHttpFormat(input *Intput) {
	input.HttpMimeType = format.TextHttpMimeTypes[4]
	input.Format       = format.InputText
}
