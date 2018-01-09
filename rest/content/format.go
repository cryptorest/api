package content

import (
	"fmt"
	"strings"

	"rest/content/format"
)

const MimeKeyRequest  = "Content-Type"
const MimeKeyResponse = "Accept"

var HttpMimePrefixs = [2]string {
	"text",
	"application",
}

var HttpMimeTypes = [4][5]string {
	format.TextHttpMimeTypes,
	format.JsonHttpMimeTypes,
	format.YamlHttpMimeTypes,
	format.TomlHttpMimeTypes,
}

func humanReadableFormat(m string) bool {
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
