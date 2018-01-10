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

func DefaultOutputHttpFormat(o *Output) {
	o.IsHumanReadable = false
	o.Format          = format.OutputText
	o.HttpMimeType    = format.TextHttpMimeTypes[0]
}

type Output struct {
	IsHumanReadable bool
	HttpMimeType    string
	Writer          http.ResponseWriter
	Structure       format.OutputStructure
	Format          func(w io.Writer, s *format.OutputStructure, hr bool) error
}

func (o *Output) FormatFind() {
	outputHttpMimeType := o.HttpMimeType
	o.HttpMimeType = EmptyString

	for f, formatHttpMimeType := range HttpMimeTypes {
		for _, httpMimeType := range formatHttpMimeType {
			if outputHttpMimeType == httpMimeType {
				o.HttpMimeType    = httpMimeType
				o.IsHumanReadable = HumanReadableFormat(httpMimeType)
				o.Format          = OutputFormatFuncs[f]

				break
			}
		}
	}

	if o.HttpMimeType == EmptyString {
		DefaultOutputHttpFormat(&*o)
	}
}

func (o *Output) Build() {
	tm := time.Now().UTC()

	if o.Structure.Error == EmptyString {
		o.Structure.Status = StatusOKString
	} else {
		o.Structure.Status  = EmptyString
		o.Structure.Content = EmptyString
	}

	o.Structure.Date      = tm.String()
	o.Structure.Timestamp = tm.Unix()

	if o.HttpMimeType != EmptyString {
		o.Structure.Host = config.Server.Host
		o.Structure.Port = config.Server.Port

		o.Writer.Header().Set(MimeKeyRequest, o.HttpMimeType)
	}

	err := o.Format(o.Writer, &o.Structure, o.IsHumanReadable)

	if err != nil {
		o.Structure.Error   = errorMessage(err)
		o.Structure.Status  = EmptyString
		o.Structure.Content = EmptyString

		io.WriteString(o.Writer, o.Structure.Error)
	}
}

var OutputHttpExecute = func(w http.ResponseWriter, r *http.Request, c string, e error) {
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
