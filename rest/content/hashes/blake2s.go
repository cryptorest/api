package hashes

import (
	"net/http"
	"golang.org/x/crypto/blake2s"

	"rest/errors"
	"rest/content"
)

const Blake2sPath = content.HashesPath + "/blake2s"

var Blake2sBits = []string{
	"256",
}

func Blake2s256(data []byte) [32]byte {
	return blake2s.Sum256(data)
}

func Blake2sHttp(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit          := content.Path2Bit(r)
	data, err, s := content.InputHttpBytes(r)

	if err == nil {
		switch bit {
		// 256
		case Blake2sBits[0]:
			content.OutputHttp32Byte(w, r, Blake2s256(data))
		}
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
