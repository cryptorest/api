package hashes

import (
	"net/http"
	"github.com/golang/crypto/blake2b"

	"rest/data"
	"rest/errors"
)

const Blake2bPath = data.HashesPath + "/blake2b"

var Blake2bBits = []string{
	"256",
	"384",
	"512",
}

func BLAKE2b(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit := data.Path2Bit(r)
	bData := []byte("data")

	switch bit {
	case Blake2bBits[0]:
		data.Write32Byte(w, blake2b.Sum256(bData))
	case Blake2bBits[1]:
		data.Write48Byte(w, blake2b.Sum384(bData))
	case Blake2bBits[2]:
		data.Write64Byte(w, blake2b.Sum512(bData))
	}
}
