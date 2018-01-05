package hashes

import (
	"net/http"
	"golang.org/x/crypto/ripemd160"

	"rest/data"
	"rest/errors"
	"rest/handlers/online"
)

const Ripemd160Path = online.HashesPath + "/ripemd160"

func RIPEMD160(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bData := []byte("data")
	md := ripemd160.New()

	md.Write(bData)
	data.WriteHash(w, md.Sum(nil))
}
