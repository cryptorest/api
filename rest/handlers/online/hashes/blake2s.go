package hashes

import (
	"net/http"
	"fmt"
	"golang.org/x/crypto/blake2s"

	"rest/handlers"
	"rest/handlers/online"
)

const BLAKE2sPath string = online.HashesPath + "/blake2s"

var BLAKE2s_Bits = [1]string{
	"256",
}

func BLAKE2s(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	bit := handlers.Path2Bit(r)
	data := []byte("data")

	switch{
	case bit == BLAKE2s_Bits[0]:
		fmt.Fprintf(w, "%x", blake2s.Sum256(data))
	}
}
