package hashes

import (
	"net/http"
	"crypto/md5"

	"rest/handlers"
	"rest/handlers/online"
)

const Md5Path string = online.HashesPath + "/md5"

func MD5(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	data := []byte("data")
	md := md5.New()

	md.Write(data)
	handlers.WriteHash(w, md.Sum(nil))
}
