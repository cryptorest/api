package hashes

import (
	"net/http"
	"golang.org/x/crypto/md4"

	"rest/errors"
	"rest/content"
)

const Md4Path = content.HashesPath + "/md4"

func MD4(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data, err, s := content.InputHttpBytes(r)
	md           := md4.New()

	if err == nil {
		md.Write(data)
		content.OutputHttpHash(w, r, md.Sum(nil))
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
