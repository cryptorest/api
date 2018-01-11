package hashes

import (
	"net/http"

	"rest/errors"
	"rest/content"
	"rest/content/data"
)

const Base32Path = content.HashesPath + "/base32"

func Base32(d []byte) string {
	str, _ := data.Base32Encode(d)

	return str
}

func Base32Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	d, err, s := content.InputHttpBytes(r)

	if err == nil {
		// Encode
		content.OutputHttpString(w, r, Base32(d))
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
