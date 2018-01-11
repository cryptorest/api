package hashes

import (
	"net/http"
	"github.com/cryptorest/go-md2"

	"rest/errors"
	"rest/content"
)

const Md2Path = content.HashesPath + "/md2"

func MD2(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data, err, s := content.InputHttpBytes(r)
	md           := md2.New()

	if err == nil {
		md.Write(data)
		content.OutputHttpHash(w, r, md.Sum(nil))
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
