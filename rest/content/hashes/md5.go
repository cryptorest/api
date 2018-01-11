package hashes

import (
	"net/http"
	"crypto/md5"

	"rest/errors"
	"rest/content"
)

const Md5Path = content.HashesPath + "/md5"

func MD5(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data, err, s := content.InputHttpBytes(r)
	md           := md5.New()

	if err == nil {
		md.Write(data)
		content.OutputHttpHash(w, r, md.Sum(nil))
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
