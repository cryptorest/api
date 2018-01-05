package hashes

import (
	"net/http"
	"golang.org/x/crypto/md4"

	"rest/data"
	"rest/errors"
)

const Md4Path = data.HashesPath + "/md4"

func MD4(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bData := []byte("data")
	md := md4.New()

	md.Write(bData)
	data.WriteHash(w, md.Sum(nil))
}
