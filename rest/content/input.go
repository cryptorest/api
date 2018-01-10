package content

import (
	"io"
	"net/http"

	"rest/content/format"
)

var InputFormatFuncs = [5]func(w io.Reader, s *format.InputStructure, hr bool) error {
	format.InputText,
	format.InputJson,
	format.InputYaml,
	format.InputToml,
	format.InputXml,
}

func InputHttpMimeType(r *http.Request) string {
	return r.Header.Get(MimeKeyRequest)
}

type Input struct {
	HttpMimeType string
	Reader       *http.Request
	Structure    format.InputStructure
	Format       func(w io.Reader, s *format.InputStructure, hr bool) error
}

func DefaultInputHttpFormat(i *Input) {
	i.Format       = format.InputText
	i.HttpMimeType = format.TextHttpMimeTypes[0]
}

func (i *Input) FormatFind() {
	inputHttpMimeType := i.HttpMimeType
	i.HttpMimeType = EmptyString

	for f, formatHttpMimeType := range HttpMimeTypes {
		for _, httpMimeType := range formatHttpMimeType {
			if inputHttpMimeType == httpMimeType {
				i.HttpMimeType = httpMimeType
				i.Format       = InputFormatFuncs[f]

				break
			}
		}
	}

	if i.HttpMimeType == EmptyString {
		DefaultInputHttpFormat(&*i)
	}
}

func (i *Input) Build() []byte {
	return i.Structure.Content
}

func InputHttpExecute(r *http.Request) []byte {
	var input Input

	input.Reader            = r
	input.HttpMimeType      = InputHttpMimeType(r)
	input.Structure         = format.InputStructure{}
	input.Structure.Content = []byte("data")

	input.FormatFind()

	return input.Build()
}

func InputBytes(r *http.Request) []byte {
	return InputHttpExecute(r)
}

func InputString(r *http.Request) string {
	return string(InputHttpExecute(r))
}
