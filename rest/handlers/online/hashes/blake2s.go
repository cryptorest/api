package hashes

import (
	"fmt"
	"net/http"
	"golang.org/x/crypto/blake2s"

	"rest/handlers"
	"rest/handlers/online"
)

const Blake2sPath string = online.HashesPath + "/blake2s"

var Blake2sBits = [1]string{
	"256",
}

func BLAKE2s(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	bit := handlers.Path2Bit(r)
	data := []byte("data")

	switch bit {
	case Blake2sBits[0]:
		fmt.Fprintf(w, "%x", blake2s.Sum256(data))
	}
}
