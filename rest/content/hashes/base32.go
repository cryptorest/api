package hashes

import (
	"net/http"
	"encoding/base32"

	"rest/content"
	"rest/errors"
)

const Base32Path = content.HashesPath + "/base32"

func Base32(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data := content.InputBytes(r)
	str  := base32.StdEncoding.EncodeToString(data)

	content.OutputString(w, r, str)
}
