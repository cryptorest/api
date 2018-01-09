package content

import (
	"fmt"
	"strings"

	"rest/content/format"
)

const MimeKeyRequest  = "Content-Type"
const MimeKeyResponse = "Accept"

const EmptyString = ""

var HttpMimePrefixs = [2]string {
	// For human readable mode
	"text",
	// For applications
	"application",
}

var HttpMimeTypes = [4][2]string {
	format.TextHttpMimeTypes,
	format.JsonHttpMimeTypes,
	format.YamlHttpMimeTypes,
	format.TomlHttpMimeTypes,
}

func HumanReadableFormat(m string) bool {
	var hr bool

	switch {
	case strings.HasPrefix(m, HttpMimePrefixs[0]):
		hr = true
	case strings.HasPrefix(m, HttpMimePrefixs[1]):
		hr = false
	default:
		hr = false
	}

	return hr
}

func errorMessage(e error) string {
	var str string

	if e == nil {
		str = ""
	} else {
		str = fmt.Sprintf("%s", e)
	}

	return str
}
