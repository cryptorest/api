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

	bData := []byte("data")
	str := base32.StdEncoding.EncodeToString(bData)

	content.WriteString(w, str)
}
