package hashes

import (
	"net/http"
	"golang.org/x/crypto/ripemd160"

	"rest/errors"
	"rest/content"
)

const Ripemd160Path = content.HashesPath + "/ripemd160"

func RIPEMD160(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data, err, s := content.InputHttpBytes(r)
	rmd          := ripemd160.New()

	if err == nil {
		rmd.Write(data)
		content.OutputHttpHash(w, r, rmd.Sum(nil))
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
