package hashes

import (
	"net/http"
	"crypto/md5"

	"rest/data"
	"rest/errors"
)

const Md5Path = data.HashesPath + "/md5"

func MD5(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bData := []byte("data")
	md := md5.New()

	md.Write(bData)
	data.WriteHash(w, md.Sum(nil))
}
