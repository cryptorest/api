package hashes

import (
	"net/http"
	"crypto/sha1"

	"rest/data"
	"rest/errors"
	"rest/handlers/online"
)

const Sha1Path = online.HashesPath + "/sha1"

func SHA1(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bData := []byte("data")
	b := sha1.New()

	b.Write(bData)
	data.WriteHash(w, b.Sum(nil))
}
