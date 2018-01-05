package hashes

import (
	"net/http"
	"golang.org/x/crypto/ripemd160"

	"rest/utils"
	"rest/errors"
	"rest/handlers/online"
)

const Ripemd160Path = online.HashesPath + "/ripemd160"

func RIPEMD160(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data := []byte("data")
	md := ripemd160.New()

	md.Write(data)
	utils.WriteHash(w, md.Sum(nil))
}
