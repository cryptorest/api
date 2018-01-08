package hashes

import (
	"net/http"
	"golang.org/x/crypto/ripemd160"

	"rest/content"
	"rest/errors"
)

const Ripemd160Path = content.HashesPath + "/ripemd160"

func RIPEMD160(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bData := []byte("data")
	md := ripemd160.New()

	md.Write(bData)
	content.OutputHash(w, r, md.Sum(nil))
}
