package content

import (
	"io"
	"fmt"
	"time"
	"strings"
	"net/http"
	"runtime/debug"

	"rest/content/format"
)

func OutputHttpMimeType(r *http.Request) string {
	return r.Header.Get(HttpMimeTypeOutputKey)
}

func DefaultOutputHttpFormat(o *Output) {
	o.IsHumanReadable = false
	o.Format          = &format.Text
	o.HttpMimeType    = format.TextHttpMimeTypes[0]
}

type Output struct {
	IsHumanReadable bool
	HttpMimeType    string
	Structure       *format.OutputStructure
	Writer          http.ResponseWriter
	Format          *format.Structure
}

func (o *Output) FormatFind() {
	outputHttpMimeType := o.HttpMimeType
	o.HttpMimeType      = EmptyString

	for _, mimeType := range strings.Split(outputHttpMimeType, HttpMimeTypeSeparator) {
		for _, f := range &Formats {
			for _, httpMimeType := range *f.MimeTypes {
				if mimeType == httpMimeType {
					o.IsHumanReadable = HumanReadableFormat(httpMimeType)
					o.HttpMimeType    = httpMimeType
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

	if o.Structure.Status.Message == EmptyString {
		o.Structure.Status.Code = http.StatusOK
	} else {
		o.Structure.Content     = EmptyString
	}

	o.Structure.Time.RFC3339 = t.Format(time.RFC3339)
	o.Structure.Time.Stamp   = t.Unix()
	o.Structure.Status.Text  = http.StatusText(o.Structure.Status.Code)

	o.Writer.Header().Set(HttpMimeTypeInputKey, o.HttpMimeType)

	err := o.Format.OutputFormatFunc(o.Writer, &*o.Structure, o.IsHumanReadable)

	if err != nil {
		o.Structure.Status.Message = err.Error()
		o.Structure.Status.Code    = http.StatusUnsupportedMediaType
		o.Structure.Status.Text    = http.StatusText(o.Structure.Status.Code)
		o.Structure.Content        = EmptyString

		io.WriteString(o.Writer, o.Structure.Status.Message)
	}
}

func (o *Output) Clean() {
	o.Structure = nil
	o.Format    = nil
	o.Writer    = nil

	o.IsHumanReadable = false
	o.HttpMimeType    = EmptyString

	debug.FreeOSMemory()
}

var OutputHttpExecute = func(w http.ResponseWriter, r *http.Request, c string, e error, s int) {
	var output Output

	output.Structure             = &format.OutputStructure{}
	output.Structure.Status.Code = s
	output.Structure.Content     = c
	output.HttpMimeType          = OutputHttpMimeType(&*r)
	output.Writer                = w

	if e == nil {
		output.Structure.Status.Message = EmptyString
	} else {
		output.Structure.Status.Message = e.Error()
	}

	output.FormatFind()
	output.Build()
	output.Clean()
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
