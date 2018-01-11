package hashes

import (
	"net/http"
	"crypto/sha1"

	"rest/errors"
	"rest/content"
)

const Sha1Path = content.HashesPath + "/sha1"

func SHA1(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data, err, s := content.InputHttpBytes(r)
	b            := sha1.New()

	if err == nil {
		b.Write(data)
		content.OutputHttpHash(w, r, b.Sum(nil))
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
