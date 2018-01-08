package hashes

import (
	"net/http"
	"crypto/sha1"

	"rest/content"
	"rest/errors"
)

const Sha1Path = content.HashesPath + "/sha1"

func SHA1(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bData := []byte("data")
	b := sha1.New()

	b.Write(bData)
	content.OutputHash(w, r, b.Sum(nil))
}
