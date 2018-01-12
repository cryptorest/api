package content

import (
	"io"
	"fmt"
	"time"
	"strings"
	"net/http"

	"rest/config"
	"rest/content/format"
)

func OutputHttpMimeType(r *http.Request) string {
	return r.Header.Get(MimeKeyResponse)
}

func DefaultOutputHttpFormat(o *Output) {
	o.IsHumanReadable = false
	o.Format          = &format.Text
	o.HttpMimeType    = format.TextHttpMimeTypes[0]
}

type Output struct {
	IsHumanReadable bool
	HttpMimeType    string
	Writer          http.ResponseWriter
	Structure       *format.OutputStructure
	Format          *format.Structure
}

func (o *Output) FormatFind() {
	outputHttpMimeType := o.HttpMimeType
	o.HttpMimeType      = EmptyString

	for _, mimeType := range strings.Split(outputHttpMimeType, ";") {
		for _, f := range &Formats {
			for _, httpMimeType := range *f.MimeTypes {
				if mimeType == httpMimeType {
					o.HttpMimeType    = httpMimeType
					o.IsHumanReadable = HumanReadableFormat(httpMimeType)
					o.Format          = &f

					return
				}
			}
		}
	}

	if o.HttpMimeType == EmptyString {
		DefaultOutputHttpFormat(&*o)
	}
}

func (o *Output) Build() {
	t := time.Now().UTC()

	if o.Structure.Error == EmptyString {
		o.Structure.Status = http.StatusOK
	} else {
		o.Structure.Content = EmptyString
		if o.Structure.Status < 100 {
			o.Structure.Status = http.StatusInternalServerError
		}
	}

	o.Structure.Date      = t.String()
	o.Structure.Timestamp = t.Unix()

	if o.HttpMimeType != EmptyString {
		o.Structure.Host = config.Server.Host
		o.Structure.Port = config.Server.Port

		o.Writer.Header().Set(MimeKeyRequest, o.HttpMimeType)
	}

	err := o.Format.OutputFormatFunc(o.Writer, &*o.Structure, o.IsHumanReadable)

	if err != nil {
		o.Structure.Error   = err.Error()
		o.Structure.Status  = http.StatusUnsupportedMediaType
		o.Structure.Content = EmptyString

		io.WriteString(o.Writer, o.Structure.Error)
	}
}

var OutputHttpExecute = func(w http.ResponseWriter, r *http.Request, c string, e error, s int) {
	var output Output

	output.Writer            = w
	output.HttpMimeType      = OutputHttpMimeType(&*r)
	output.Structure         = &format.OutputStructure{}
	output.Structure.Status  = s
	output.Structure.Content = c

	if e == nil {
		output.Structure.Error = EmptyString
	} else {
		output.Structure.Error = e.Error()
	}

	output.FormatFind()
	output.Build()
}

func OutputHttpHash(w http.ResponseWriter, r *http.Request, b []byte) {
	OutputHttpExecute(w, &*r, fmt.Sprintf("%x", b), nil, 0)
}

func OutputHttpByte(w http.ResponseWriter, r *http.Request, b []byte) {
	OutputHttpExecute(w, &*r, fmt.Sprintf("%s", b), nil, 0)
}

func OutputHttpString(w http.ResponseWriter, r *http.Request, s string) {
	OutputHttpExecute(w, &*r, s, nil, 0)
}

func OutputHttpUInt8(w http.ResponseWriter, r *http.Request, i uint8) {
	OutputHttpExecute(w, &*r, fmt.Sprintf("%x", i), nil, 0)
}

func OutputHttpUInt32(w http.ResponseWriter, r *http.Request, i uint32) {
	OutputHttpExecute(w, &*r, fmt.Sprintf("%x", i), nil, 0)
}

func OutputHttpUInt64(w http.ResponseWriter, r *http.Request, i uint64) {
	OutputHttpExecute(w, &*r, fmt.Sprintf("%x", i), nil, 0)
}

func OutputHttp32Byte(w http.ResponseWriter, r *http.Request, b [32]byte) {
	OutputHttpExecute(w, &*r, fmt.Sprintf("%x", b), nil, 0)
}

func OutputHttp48Byte(w http.ResponseWriter, r *http.Request, b [48]byte) {
	OutputHttpExecute(w, &*r, fmt.Sprintf("%x", b), nil, 0)
}

func OutputHttp64Byte(w http.ResponseWriter, r *http.Request, b [64]byte) {
	OutputHttpExecute(w, &*r, fmt.Sprintf("%x", b), nil, 0)
}

func OutputHttpError(w http.ResponseWriter, r *http.Request, e error, s int) {
	OutputHttpExecute(w, &*r, EmptyString, e, s)
}
