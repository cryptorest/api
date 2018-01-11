package data

import (
	"net/http"
	"encoding/base32"

	"rest/errors"
	"rest/content"
)

const Base32Path = content.DataPath + "/base32"

var Base32Actions = []string{
	"encode",
	"decode",
}

func Base32(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	action := content.Path2Action(r)

	switch action {
	// Encode
	case Base32Actions[0]:
		data, err, s := content.InputHttpBytes(r)

		if err == nil {
			str := base32.StdEncoding.EncodeToString(data)

			content.OutputHttpString(w, r, str)
		} else {
			content.OutputHttpError(w, r, err, s)
		}
	// Decode
	case Base32Actions[1]:
		str, err, s := content.InputHttpString(r)

		if err == nil {
			data, err := base32.StdEncoding.DecodeString(str)

			if err == nil {
				content.OutputHttpBytes(w, r, data)
			} else {
				content.OutputHttpError(w, r, err, http.StatusUnprocessableEntity)
			}
		} else {
			content.OutputHttpError(w, r, err, s)
		}
	}
}
