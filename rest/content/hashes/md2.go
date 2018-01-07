package hashes

import (
	"net/http"
	"github.com/cryptorest/go-md2"

	"rest/content"
	"rest/errors"
)

const Md2Path = content.HashesPath + "/md2"

func MD2(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bData := []byte("data")
	md := md2.New()

	md.Write(bData)
	content.WriteHash(w, md.Sum(nil))
}
