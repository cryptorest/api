package hashes

import (
	"net/http"
	"encoding/base64"

	"rest/errors"
	"rest/content"
)

const Base64Path = content.HashesPath + "/base64"

func Base64(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data, err, s := content.InputHttpBytes(r)

	if err == nil {
		// Encode
		str := base64.StdEncoding.EncodeToString(data)

		content.OutputHttpString(w, r, str)
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
