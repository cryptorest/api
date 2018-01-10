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

	data := content.InputBytes(r)
	b    := sha1.New()

	b.Write(data)
	content.OutputHash(w, r, b.Sum(nil))
}
