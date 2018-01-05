package hashes

import (
	"net/http"
	"encoding/base64"

	"rest/data"
	"rest/errors"
)

const Base64Path = data.HashesPath + "/base64"

func Base64(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bData := []byte("data")
	str := base64.StdEncoding.EncodeToString(bData)

	data.WriteString(w, str)
}
