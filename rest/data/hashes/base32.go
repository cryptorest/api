package hashes

import (
	"net/http"
	"encoding/base32"

	"rest/data"
	"rest/errors"
)

const Base32Path = data.HashesPath + "/base32"

func Base32(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bData := []byte("data")
	str := base32.StdEncoding.EncodeToString(bData)

	data.WriteString(w, str)
}
