package hashes

import (
	"net/http"
	"crypto/sha1"

	"rest/data"
	"rest/errors"
)

const Sha1Path = data.HashesPath + "/sha1"

func SHA1(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bData := []byte("data")
	b := sha1.New()

	b.Write(bData)
	data.WriteHash(w, b.Sum(nil))
}
