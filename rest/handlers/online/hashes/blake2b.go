package hashes

import (
	"fmt"
	"net/http"
	"golang.org/x/crypto/blake2b"

	"rest/handlers"
	"rest/handlers/online"
)

const Blake2bPath string = online.HashesPath + "/blake2b"

var Blake2bBits = [3]string{
	"256",
	"384",
	"512",
}

func BLAKE2b(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	bit := handlers.Path2Bit(r)
	data := []byte("data")

	switch bit {
	case Blake2bBits[0]:
		fmt.Fprintf(w, "%x", blake2b.Sum256(data))
	case Blake2bBits[1]:
		fmt.Fprintf(w, "%x", blake2b.Sum384(data))
	case Blake2bBits[2]:
		fmt.Fprintf(w, "%x", blake2b.Sum512(data))
	}
}
