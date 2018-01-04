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

var defaultBits = [2]int{
	256,
	512,
}

func SHAKE(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	bit := handlers.Path2Bit(r)
	data := []byte("data")

	switch {
	case bit == ShakeBits[0]:
		hash := make([]byte, defaultBits[0] / 8)
		sha3.ShakeSum128(hash, data)

		handlers.WriteBytes(w, []byte(hex.EncodeToString(hash)))
	case bit == ShakeBits[1]:
		hash := make([]byte, defaultBits[1] / 8)
		sha3.ShakeSum256(hash, data)

		handlers.WriteBytes(w, []byte(hex.EncodeToString(hash)))
	}
}
