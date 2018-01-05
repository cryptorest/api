package hashes

import (
	"net/http"
	"github.com/cryptorest/go-md2"

	"rest/data"
	"rest/errors"
)

const Md2Path = data.HashesPath + "/md2"

func MD2(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bData := []byte("data")
	md := md2.New()

	md.Write(bData)
	data.WriteHash(w, md.Sum(nil))
}
