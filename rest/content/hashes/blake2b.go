package hashes

import (
	"net/http"
	"github.com/golang/crypto/blake2b"

	"rest/errors"
	"rest/content"
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

	bit          := content.Path2Bit(r)
	data, err, s := content.InputHttpBytes(r)

	if err == nil {
		switch bit {
		// 256
		case Blake2bBits[0]:
			content.OutputHttp32Byte(w, r, blake2b.Sum256(data))
		// 384
		case Blake2bBits[1]:
			content.OutputHttp48Byte(w, r, blake2b.Sum384(data))
		// 512
		case Blake2bBits[2]:
			content.OutputHttp64Byte(w, r, blake2b.Sum512(data))
		}
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
