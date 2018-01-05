package hashes

import (
	"net/http"
	"crypto/md5"

	"rest/utils"
	"rest/errors"
	"rest/handlers/online"
)

const Md5Path = online.HashesPath + "/md5"

func MD5(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data := []byte("data")
	md := md5.New()

	md.Write(data)
	utils.WriteHash(w, md.Sum(nil))
}
