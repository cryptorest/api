package hashes

import (
	"net/http"
	"encoding/hex"
	"golang.org/x/crypto/sha3"

	"rest/content"
	"rest/errors"
)

const ShakePath = content.HashesPath + "/shake"

var ShakeBits = []string{
	"128",
	"256",
}

var minBits = 256
//var maxBits = 1024 * 1024

func SHAKE(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit := content.Path2Bit(r)
	bData := []byte("data")

	switch bit {
	case ShakeBits[0]:
		hash := make([]byte, minBits / 8)
		sha3.ShakeSum128(hash, bData)

		content.WriteString(w, hex.EncodeToString(hash))
	case ShakeBits[1]:
		hash := make([]byte, minBits / 4)
		sha3.ShakeSum256(hash, bData)

		content.WriteString(w, hex.EncodeToString(hash))
	}
}
