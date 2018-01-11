package hashes

import (
	"net/http"
	"golang.org/x/crypto/md4"

	"rest/errors"
	"rest/content"
)

const Md4Path = content.HashesPath + "/md4"

func Md4(data []byte) []byte {
	md := md4.New()

	md.Write(data)

	return md.Sum(nil)
}

func Md4Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data, err, s := content.InputHttpBytes(r)

	if err == nil {
		content.OutputHttpByte(w, r, Md4(data))
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
