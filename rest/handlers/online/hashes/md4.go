package hashes

import (
	"net/http"
	"golang.org/x/crypto/md4"

	"rest/handlers"
	"rest/handlers/online"
)

const Md4Path string = online.HashesPath + "/md4"

func MD4(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	data := []byte("data")
	md := md4.New()

	md.Write(data)
	handlers.WriteHash(w, md.Sum(nil))
}
