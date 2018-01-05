package hashes

import (
	"net/http"
	"github.com/cryptorest/go-md2"

	"rest/utils"
	"rest/errors"
	"rest/handlers/online"
)

const Md2Path = online.HashesPath + "/md2"

func MD2(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	data := []byte("data")
	md := md2.New()

	md.Write(data)
	utils.WriteHash(w, md.Sum(nil))
}
