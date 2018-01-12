package hashes

import (
	"net/http"
	"crypto/sha1"

	"rest/errors"
	"rest/content"
)

const Sha1Path = content.HashesPath + "/sha1"

func Sha1(data []byte) []byte {
	b := sha1.New()

	return hashSum(data, b)
}

func Sha1Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data, err, s := content.InputHttpBytes(r)

	if err == nil {
		content.OutputHttpHash(w, r, Sha1(data))
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
