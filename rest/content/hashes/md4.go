package hashes

import (
	"net/http"
	"golang.org/x/crypto/md4"

	"rest/content"
	"rest/errors"
)

const Md4Path = content.HashesPath + "/md4"

func MD4(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data := content.InputBytes(r)
	md   := md4.New()

	md.Write(data)
	content.OutputHash(w, r, md.Sum(nil))
}
