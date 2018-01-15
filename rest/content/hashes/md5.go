package hashes

import (
	"net/http"
	"crypto/md5"

	"rest/errors"
	"rest/content"
)

const Md5Path = content.HashesPath + "/md5"

func Md5(data []byte) []byte {
	md := md5.New()

	md.Write(data)

	return md.Sum(nil)
}

func Md5Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, &*r) {
		return
	}

	data, err, s := content.InputHttpBytes(&*r, false, false)

	if err == nil {
		content.OutputHttpByte(w, &*r, Md5(data))
	} else {
		content.OutputHttpError(w, &*r, err, s)
	}
}
