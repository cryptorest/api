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

func Blake2b256(data []byte) [32]byte {
	return blake2b.Sum256(data)
}

func Blake2b384(data []byte) [48]byte {
	return blake2b.Sum384(data)
}

func Blake2b512(data []byte) [64]byte {
	return blake2b.Sum512(data)
}

func Blake2bHttp(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, &*r) {
		return
	}

	bit          := content.Path2Bit(&*r)
	data, err, s := content.InputHttpBytes(&*r)

	if err == nil {
		switch bit {
		// 256
		case Blake2bBits[0]:
			content.OutputHttp32Byte(w, &*r, Blake2b256(data))
		// 384
		case Blake2bBits[1]:
			content.OutputHttp48Byte(w, &*r, Blake2b384(data))
		// 512
		case Blake2bBits[2]:
			content.OutputHttp64Byte(w, &*r, Blake2b512(data))
		}
	} else {
		content.OutputHttpError(w, &*r, err, s)
	}
}
