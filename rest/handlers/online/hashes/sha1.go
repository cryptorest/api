package hashes

import (
	"net/http"
	"crypto/sha1"

	"rest/utils"
	"rest/errors"
	"rest/handlers/online"
)

const Sha1Path = online.HashesPath + "/sha1"

func SHA1(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data := []byte("data")
	b := sha1.New()

	b.Write(data)
	utils.WriteHash(w, b.Sum(nil))
}
