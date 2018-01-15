package content

import (
	"strings"

	"rest/content/format"
)

const HttpMimeTypeSeparator = ";"
const HttpMimeTypeInputKey  = "Content-Type"
const HttpMimeTypeInputSize = "Content-Length"
const HttpMimeTypeInputFile = "multipart/form-data"
const HttpMimeTypeOutputKey = "Accept"

const EmptyString = ""

var HttpMimePrefixes = [2]string {
	// For human readable mode
	"text/",
	// For applications
	"application/",
}

var Formats = [5]format.Structure {
	format.Text,
	format.Json,
	format.Yaml,
	format.Toml,
	format.Xml,
}

type conf struct {
	BufferSize      int
	FileSizeLimit   int64
	BodySizeLimit   int64
	TemporaryDeploy bool
	TmpDir          *string
	UploadDir       *string
}

var Config = &conf{}

func HumanReadableFormat(m string) bool {
	var hr bool

	switch {
	case strings.HasPrefix(m, HttpMimePrefixes[0]):
		hr = true
	case strings.HasPrefix(m, HttpMimePrefixes[1]):
		hr = false
	default:
		hr = false
	}

	return hr
}
