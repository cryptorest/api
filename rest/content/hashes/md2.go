package hashes

import (
	"net/http"
	"github.com/cryptorest/go-md2"

	"rest/errors"
	"rest/content"
)

const Md2Path = content.HashesPath + "/md2"

func Md2(data []byte) []byte {
	md := md2.New()

	md.Write(data)

	return md.Sum(nil)
}

func Md2Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data, err, s := content.InputHttpBytes(r)

	if err == nil {
		content.OutputHttpByte(w, r, Md2(data))
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
