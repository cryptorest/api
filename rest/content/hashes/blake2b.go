package hashes

import (
	"net/http"
	"github.com/golang/crypto/blake2b"

	"rest/content"
	"rest/errors"
)

const Blake2bPath = content.HashesPath + "/blake2b"

var Blake2bBits = []string{
	"256",
	"384",
	"512",
}

func BLAKE2b(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit  := content.Path2Bit(r)
	data := content.InputBytes(r)

	switch bit {
	case Blake2bBits[0]:
		content.Output32Byte(w, r, blake2b.Sum256(data))
	case Blake2bBits[1]:
		content.Output48Byte(w, r, blake2b.Sum384(data))
	case Blake2bBits[2]:
		content.Output64Byte(w, r, blake2b.Sum512(data))
	}
}
