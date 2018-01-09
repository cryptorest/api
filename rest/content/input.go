package content

import (
	"io"
	"fmt"
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

func InputFormat(input *Input) {
	inputHttpMimeType := input.HttpMimeType
	input.HttpMimeType = ""

	for i, formatHttpMimeType := range HttpMimeTypes {
		for _, httpMimeType := range formatHttpMimeType {
			if inputHttpMimeType == httpMimeType {
				input.HttpMimeType = httpMimeType
				input.Format       = InputFormatFuncs[i]

				break
			}
		}
	}

	if input.HttpMimeType == "" {
		DefaultInputHttpFormat(&*input)
	}
}

func InputBuild(input *Input) {

}

func InputHttpExecute(r *http.Request, c string) {
	var input Input

	input.Reader            = r
	input.HttpMimeType      = InputHttpMimeType(r)
	input.Structure         = format.InputStructure{}
	input.Structure.Content = c

	InputFormat(&input)
	InputBuild(&input)
}

func InputHash(r *http.Request, b []byte) {
	InputHttpExecute(r, fmt.Sprintf("%x", b))
}

func InputBytes(r *http.Request, b []byte) {
	InputHttpExecute(r, fmt.Sprintf("%s", b))
}
