package hashes

import (
	"net/http"
	"github.com/cryptorest/go-md2"

	"rest/handlers"
	"rest/handlers/online"
)

const Md2Path string = online.HashesPath + "/md2"

func MD2(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	data := []byte("data")
	md := md2.New()

	md.Write(data)
	handlers.WriteHash(w, md.Sum(nil))
}
