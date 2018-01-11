package hashes

import (
	"net/http"
	"encoding/hex"
	"golang.org/x/crypto/sha3"

	"rest/errors"
	"rest/content"
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

	bit          := content.Path2Bit(r)
	data, err, s := content.InputHttpBytes(r)

	if err == nil {
		switch bit {
		// 128
		case ShakeBits[0]:
			hash := make([]byte, minBits / 8)

			sha3.ShakeSum128(hash, data)

			content.OutputHttpString(w, r, hex.EncodeToString(hash))
		// 256
		case ShakeBits[1]:
			hash := make([]byte, minBits / 4)

			sha3.ShakeSum256(hash, data)

			content.OutputHttpString(w, r, hex.EncodeToString(hash))
		}
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
