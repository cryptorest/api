package content

import (
	"io"
	"fmt"
	"time"
	"net/http"

	"rest/config"
	"rest/content/format"
)

var OutputFormatFuncs = [4]func(w io.Writer, s *format.OutputStructure, hr bool) error {
	format.OutputText,
	format.OutputJson,
	format.OutputYaml,
	format.OutputToml,
}

func OutputHttpMimeType(r *http.Request) string {
	return r.Header.Get(MimeKeyResponse)
}

type Output struct {
	Format          func(w io.Writer, s *format.OutputStructure, hr bool) error
	HttpMimeType    string
	Writer          http.ResponseWriter
	Structure       format.OutputStructure
	IsHumanReadable bool
}

func DefaultOutputHttpFormat(output *Output) {
	output.HttpMimeType    = format.TextHttpMimeTypes[4]
	output.IsHumanReadable = false
	output.Format          = format.OutputText
}

func OutputFormat(output *Output) {
	outputHttpMimeType := output.HttpMimeType
	output.HttpMimeType = ""

	for i, formatHttpMimeType := range HttpMimeTypes {
		for _, httpMimeType := range formatHttpMimeType {
			if outputHttpMimeType == httpMimeType {
				output.HttpMimeType    = httpMimeType
				output.IsHumanReadable = humanReadableFormat(httpMimeType)
				output.Format          = OutputFormatFuncs[i]

				break
			}
		}
	}

	if output.HttpMimeType == "" {
		DefaultOutputHttpFormat(&*output)
	}
}

func OutputBuild(output *Output) {
	tm  := time.Now().UTC()

	if output.Structure.Error == "" {
		output.Structure.Status = "OK"
	} else {
		output.Structure.Status  = ""
		output.Structure.Content = ""
	}

	output.Structure.Date      = tm.String()
	output.Structure.Timestamp = tm.Unix()

	if output.HttpMimeType != "" {
		output.Structure.Host = config.Server.Host
		output.Structure.Port = config.Server.Port

		if output.Writer != nil {
			output.Writer.Header().Set(MimeKeyRequest, output.HttpMimeType)
		}
	}

	err := output.Format(output.Writer, &output.Structure, output.IsHumanReadable)

	if err != nil {
		output.Structure.Error   = errorMessage(err)
		output.Structure.Status  = ""
		output.Structure.Content = ""

		if output.Writer != nil {
			io.WriteString(output.Writer, output.Structure.Error)
		}
	}
}

func OutputHttpExecute(w http.ResponseWriter, r *http.Request, c string, e error) {
	var output    Output
	var structure format.OutputStructure

	output.Writer            = w
	output.HttpMimeType      = OutputHttpMimeType(r)
	output.Structure         = structure
	output.Structure.Error   = errorMessage(e)
	output.Structure.Content = c

	OutputFormat(&output)
	OutputBuild(&output)
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
	OutputHttpExecute(w, r, "", e)
}
