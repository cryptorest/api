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

func DefaultInputHttpFormat(input *Input) {
	input.Format       = format.InputText
	input.HttpMimeType = format.TextHttpMimeTypes[0]
}

func (input *Input) FormatFind() {
	inputHttpMimeType := input.HttpMimeType
	input.HttpMimeType = EmptyString

	for i, formatHttpMimeType := range HttpMimeTypes {
		for _, httpMimeType := range formatHttpMimeType {
			if inputHttpMimeType == httpMimeType {
				input.HttpMimeType = httpMimeType
				input.Format       = InputFormatFuncs[i]

				break
			}
		}
	}

	if input.HttpMimeType == EmptyString {
		DefaultInputHttpFormat(&*input)
	}
}

func (input Input) Build() []byte {
	return input.Structure.Content
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
