package hashes

import (
	"net/http"
	"encoding/hex"
	"golang.org/x/crypto/sha3"

	"rest/handlers"
	"rest/handlers/online"
)

const ShakePath string = online.HashesPath + "/shake"

var ShakeBits = [2]string{
	"128",
	"256",
}

var minBits = 256
//var maxBits = 1024 * 1024

func SHAKE(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	bit := handlers.Path2Bit(r)
	data := []byte("data")

	switch bit {
	case ShakeBits[0]:
		hash := make([]byte, minBits / 8)
		sha3.ShakeSum128(hash, data)

		handlers.WriteString(w, hex.EncodeToString(hash))
	case ShakeBits[1]:
		hash := make([]byte, minBits / 4)
		sha3.ShakeSum256(hash, data)

		handlers.WriteString(w, hex.EncodeToString(hash))
	}
}
