package content

import (
	"io"
	"fmt"
	"time"
	"net/http"

	"rest/config"
	"rest/content/format"
)

const StatusOKString = "OK"

var OutputFormatFuncs = [5]func(w io.Writer, s *format.OutputStructure, hr bool) error {
	format.OutputText,
	format.OutputJson,
	format.OutputYaml,
	format.OutputToml,
	format.OutputXml,
}

func OutputHttpMimeType(r *http.Request) string {
	return r.Header.Get(MimeKeyResponse)
}

func DefaultOutputHttpFormat(output *Output) {
	output.IsHumanReadable = false
	output.Format          = format.OutputText
	output.HttpMimeType    = format.TextHttpMimeTypes[0]
}

type Output struct {
	IsHumanReadable bool
	HttpMimeType    string
	Writer          http.ResponseWriter
	Structure       format.OutputStructure
	Format          func(w io.Writer, s *format.OutputStructure, hr bool) error
}

func (output *Output) FormatFind() {
	outputHttpMimeType := output.HttpMimeType
	output.HttpMimeType = EmptyString

	for i, formatHttpMimeType := range HttpMimeTypes {
		for _, httpMimeType := range formatHttpMimeType {
			if outputHttpMimeType == httpMimeType {
				output.HttpMimeType    = httpMimeType
				output.IsHumanReadable = HumanReadableFormat(httpMimeType)
				output.Format          = OutputFormatFuncs[i]

				break
			}
		}
	}

	if output.HttpMimeType == EmptyString {
		DefaultOutputHttpFormat(&*output)
	}
}

func (output *Output) Build() {
	tm := time.Now().UTC()

	if output.Structure.Error == EmptyString {
		output.Structure.Status = StatusOKString
	} else {
		output.Structure.Status  = EmptyString
		output.Structure.Content = EmptyString
	}

	output.Structure.Date      = tm.String()
	output.Structure.Timestamp = tm.Unix()

	if output.HttpMimeType != EmptyString {
		output.Structure.Host = config.Server.Host
		output.Structure.Port = config.Server.Port

		output.Writer.Header().Set(MimeKeyRequest, output.HttpMimeType)
	}

	err := output.Format(output.Writer, &output.Structure, output.IsHumanReadable)

	if err != nil {
		output.Structure.Error   = errorMessage(err)
		output.Structure.Status  = EmptyString
		output.Structure.Content = EmptyString

		io.WriteString(output.Writer, output.Structure.Error)
	}
}

func OutputHttpExecute(w http.ResponseWriter, r *http.Request, c string, e error) {
	var output Output

	output.Writer            = w
	output.HttpMimeType      = OutputHttpMimeType(r)
	output.Structure         = format.OutputStructure{}
	output.Structure.Error   = errorMessage(e)
	output.Structure.Content = c

	output.FormatFind()
	output.Build()
}

func OutputHash(w http.ResponseWriter, r *http.Request, b []byte) {
	OutputHttpExecute(w, r, fmt.Sprintf("%x", b), nil)
}

func OutputBytes(w http.ResponseWriter, r *http.Request, b []byte) {
	OutputHttpExecute(w, r, fmt.Sprintf("%s", b), nil)
}

func OutputString(w http.ResponseWriter, r *http.Request, s string) {
	OutputHttpExecute(w, r, s, nil)
}

func OutputUInt8(w http.ResponseWriter, r *http.Request, i uint8) {
	OutputHttpExecute(w, r, fmt.Sprintf("%x", i), nil)
}

func OutputUInt32(w http.ResponseWriter, r *http.Request, i uint32) {
	OutputHttpExecute(w, r, fmt.Sprintf("%x", i), nil)
}

func OutputUInt64(w http.ResponseWriter, r *http.Request, i uint64) {
	OutputHttpExecute(w, r, fmt.Sprintf("%x", i), nil)
}

func Output32Byte(w http.ResponseWriter, r *http.Request, b [32]byte) {
	OutputHttpExecute(w, r, fmt.Sprintf("%x", b), nil)
}

func Output48Byte(w http.ResponseWriter, r *http.Request, b [48]byte) {
	OutputHttpExecute(w, r, fmt.Sprintf("%x", b), nil)
}

func Output64Byte(w http.ResponseWriter, r *http.Request, b [64]byte) {
	OutputHttpExecute(w, r, fmt.Sprintf("%x", b), nil)
}

func OutputError(w http.ResponseWriter, r *http.Request, e error) {
	OutputHttpExecute(w, r, EmptyString, e)
}
