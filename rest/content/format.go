package content

import (
	"strings"

	"rest/content/format"
)

const MimeKeyRequest  = "Content-Type"
const MimeKeyResponse = "Accept"

const EmptyString = ""

var HttpMimePrefixes = [2]string {
	// For human readable mode
	"text",
	// For applications
	"application",
}

var Formats = [5]format.Structure {
	format.Text,
	format.Json,
	format.Yaml,
	format.Toml,
	format.Xml,
}

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
