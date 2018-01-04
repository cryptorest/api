package hashes

import (
	"net/http"
	"crypto/sha1"

	"rest/handlers"
	"rest/handlers/online"
)

const Sha1Path string = online.HashesPath + "/sha1"

func SHA1(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	data := []byte("data")
	b := sha1.New()

	b.Write(data)
	handlers.WriteHash(w, b.Sum(nil))
}
