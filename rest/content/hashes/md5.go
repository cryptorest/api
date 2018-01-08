package hashes

import (
	"net/http"
	"crypto/md5"

	"rest/content"
	"rest/errors"
)

const Md5Path = content.HashesPath + "/md5"

func MD5(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bData := []byte("data")
	md := md5.New()

	md.Write(bData)
	content.OutputHash(w, r, md.Sum(nil))
}
