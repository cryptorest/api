package hashes

import (
	"net/http"
	"encoding/base32"

	"rest/errors"
	"rest/content"
)

const Base32Path = content.HashesPath + "/base32"

func Base32(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data, err, s := content.InputHttpBytes(r)

	if err == nil {
		// Encode
		str := base32.StdEncoding.EncodeToString(data)

		content.OutputHttpString(w, r, str)
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
