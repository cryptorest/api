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

	bit := content.Path2Bit(r)
	bData := []byte("data")

	switch bit {
	case Blake2bBits[0]:
		content.Write32Byte(w, blake2b.Sum256(bData))
	case Blake2bBits[1]:
		content.Write48Byte(w, blake2b.Sum384(bData))
	case Blake2bBits[2]:
		content.Write64Byte(w, blake2b.Sum512(bData))
	}
}
