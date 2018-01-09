package content

import (
	"io"
	"fmt"
	"time"
	"net/http"

	"rest/config"
	"rest/content/format"
)

var OutputFuncs = [4]func(w http.ResponseWriter, c *format.OutputStructure, hr bool) error {
	format.OutputText,
	format.OutputJson,
	format.OutputYaml,
	format.OutputToml,
}

func DefaultOutputFormat() (string, bool, func(w http.ResponseWriter, c *format.OutputStructure, hr bool) error) {
	return format.TextMimeTypes[4], false, format.OutputText
}

func OutputFormat(r *http.Request) (string, bool, func(w http.ResponseWriter, c *format.OutputStructure, hr bool) error) {
	var mimeType string
	var hr bool
	var f func(w http.ResponseWriter, c *format.OutputStructure, hr bool) error
	rType := r.Header.Get(MimeKeyResponse)

	if rType == "" {
		return DefaultOutputFormat()
	}

	for i, m := range MimeTypes {
		for _, t := range m {
			if rType == t {
				mimeType = rType
				hr = humanReadableFormat(mimeType)
				f = OutputFuncs[i]

				break
			}
		}
	}

	if mimeType == "" {
		mimeType, hr, f = DefaultOutputFormat()
	}

	return mimeType, hr, f
}

func Output(w http.ResponseWriter, r *http.Request, d string, e error) {
	var eStr string
	tm := time.Now().UTC()
	t, hr, f := OutputFormat(r)

	if e == nil {
		eStr = ""
	} else {
		eStr = errorContent(e)
		d = ""
	}

	var outputStruct = &format.OutputStructure{
		Host:      config.Server.Host,
		Port:      config.Server.Port,
		Content:   d,
		Error:     eStr,
		Date:      tm.String(),
		Timestamp: tm.Unix(),
	}

	err := f(w, outputStruct, hr)

	if err != nil {
		d = ""
		outputStruct.Content = d
		outputStruct.Error = errorContent(e)

		io.WriteString(w, outputStruct.Error)
	}

	w.Header().Set(MimeKeyRequest, t)
}

func OutputHash(w http.ResponseWriter, r *http.Request, b []byte) {
	Output(w, r, fmt.Sprintf("%x", b), nil)
}

func OutputBytes(w http.ResponseWriter, r *http.Request, b []byte) {
	Output(w, r, fmt.Sprintf("%s", b), nil)
}

func OutputString(w http.ResponseWriter, r *http.Request, s string) {
	Output(w, r, s, nil)
}

func OutputUInt8(w http.ResponseWriter, r *http.Request, i uint8) {
	Output(w, r, fmt.Sprintf("%x", i), nil)
}

func OutputUInt32(w http.ResponseWriter, r *http.Request, i uint32) {
	Output(w, r, fmt.Sprintf("%x", i), nil)
}

func OutputUInt64(w http.ResponseWriter, r *http.Request, i uint64) {
	Output(w, r, fmt.Sprintf("%x", i), nil)
}

func Output32Byte(w http.ResponseWriter, r *http.Request, b [32]byte) {
	Output(w, r, fmt.Sprintf("%x", b), nil)
}

func Output48Byte(w http.ResponseWriter, r *http.Request, b [48]byte) {
	Output(w, r, fmt.Sprintf("%x", b), nil)
}

func Output64Byte(w http.ResponseWriter, r *http.Request, b [64]byte) {
	Output(w, r, fmt.Sprintf("%x", b), nil)
}

func OutputError(w http.ResponseWriter, r *http.Request, e error) {
	Output(w, r, "", e)
}
