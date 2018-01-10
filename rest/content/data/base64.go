package data

import (
	"net/http"
	"encoding/base64"

	"rest/content"
	"rest/errors"
)

const Base64Path = content.DataPath + "/base64"

var Base64Actions = []string{
	"encode",
	"decode",
}

func Base64(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	action := content.Path2Action(r)

	switch action {
	case Base64Actions[0]:
		data := content.InputBytes(r)
		str  := base64.StdEncoding.EncodeToString(data)

		content.OutputString(w, r, str)
	case Base64Actions[1]:
		str        := content.InputString(r)
		bData, err := base64.StdEncoding.DecodeString(str)

		if err != nil {
			content.OutputError(w, r, err)

			return
		}

		content.OutputBytes(w, r, bData)
	}
}
