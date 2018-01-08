package content

import (
	"fmt"
	"strings"

	"rest/content/format"
)

const MimeKeyRequest  = "Content-Type"
const MimeKeyResponse = "Accept"

var MimePrefixs = [2]string {
	"text",
	"application",
}

var MimeTypes = [4][5]string {
	format.TextMimeTypes,
	format.JsonMimeTypes,
	format.YamlMimeTypes,
	format.TomlMimeTypes,
}

func humanReadableFormat(m string) bool {
	var hr bool

	switch {
	case strings.HasPrefix(m, MimePrefixs[0]):
		hr = true
	case strings.HasPrefix(m, MimePrefixs[1]):
		hr = false
	default:
		hr = false
	}

	return hr
}

func errorContent(e error) string {
	return fmt.Sprintf("%s", e)
}
