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

var HttpMimeTypes = [5][2]string {
	format.TextHttpMimeTypes,
	format.JsonHttpMimeTypes,
	format.YamlHttpMimeTypes,
	format.TomlHttpMimeTypes,
	format.XmlHttpMimeTypes,
}

var FileExtensions = [5][]string {
	format.TextFileExtensions,
	format.JsonFileExtensions,
	format.YamlFileExtensions,
	format.TomlFileExtensions,
	format.XmlFileExtensions,
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
