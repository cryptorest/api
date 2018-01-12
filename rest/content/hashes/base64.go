package hashes

import (
	"net/http"

	"rest/errors"
	"rest/content"
	"rest/content/data"
)

const Base64Path = content.HashesPath + "/base64"

func Base64(d []byte) string {
	str, _ := data.Base64Encode(d)

	return str
}

func Base64Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, &*r) {
		return
	}

	d, err, s := content.InputHttpBytes(&*r)

	if err == nil {
		// Encode
		content.OutputHttpString(w, &*r, Base64(d))
	} else {
		content.OutputHttpError(w, &*r, err, s)
	}
}
