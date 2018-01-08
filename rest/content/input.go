package content

import (
	"net/http"

	"rest/content/format"
)


var InputFuncs = [4]func(c *format.InputStructure, hr bool) (string, error) {
	format.InputText,
	format.InputJson,
	format.InputYaml,
	format.InputToml,
}

func DefaultInputFormat() (string, bool, func(c *format.InputStructure, hr bool) (string, error)) {
	return format.TextMimeTypes[4], false, format.InputText
}

func InputFormat(r *http.Request) (string, bool, func(c *format.InputStructure, hr bool) (string, error)) {
	var mimeType string
	var hr bool
	var f func(c *format.InputStructure, hr bool) (string, error)
	rType := r.Header.Get(MimeKeyRequest)

	if rType == "" {
		return DefaultInputFormat()
	}

	for i, m := range MimeTypes {
		for _, t := range m {
			if rType == t {
				mimeType = rType
				hr = humanReadableFormat(mimeType)
				f = InputFuncs[i]

				break
			}
		}
	}

	if mimeType == "" {
		mimeType, hr, f = DefaultInputFormat()
	}

	return mimeType, hr, f
}
