package hashes

import (
	"net/http"
	"golang.org/x/crypto/ripemd160"

	"rest/handlers"
	"rest/handlers/online"
)

const RIPEMD160Path string = online.HashesPath + "/ripemd160"

func RIPEMD160(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	data := []byte("data")
	md := ripemd160.New()

	md.Write(data)
	handlers.WriteHash(w, md.Sum(nil))
}
