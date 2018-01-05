package hashes

import (
	"net/http"
	"encoding/hex"
	"golang.org/x/crypto/sha3"

	"rest/utils"
	"rest/errors"
	"rest/handlers/online"
)

const ShakePath = online.HashesPath + "/shake"

var ShakeBits = [2]string{
	"128",
	"256",
}

var minBits = 256
//var maxBits = 1024 * 1024

func SHAKE(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit := utils.Path2Bit(r)
	data := []byte("data")

	switch bit {
	case ShakeBits[0]:
		hash := make([]byte, minBits / 8)
		sha3.ShakeSum128(hash, data)

		utils.WriteString(w, hex.EncodeToString(hash))
	case ShakeBits[1]:
		hash := make([]byte, minBits / 4)
		sha3.ShakeSum256(hash, data)

		utils.WriteString(w, hex.EncodeToString(hash))
	}
}
