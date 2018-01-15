package hashes

import (
	"net/http"
	"golang.org/x/crypto/ripemd160"

	"rest/errors"
	"rest/content"
)

const Ripemd160Path = content.HashesPath + "/ripemd160"

func Ripemd160(data []byte) []byte {
	rmd := ripemd160.New()

	rmd.Write(data)

	return rmd.Sum(nil)
}

func Ripemd160Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, &*r) {
		return
	}

	data, err, s := content.InputHttpBytes(&*r, false, false)

	if err == nil {
		content.OutputHttpByte(w, &*r, Ripemd160(data))
	} else {
		content.OutputHttpError(w, &*r, err, s)
	}
}
