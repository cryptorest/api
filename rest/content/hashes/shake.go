package hashes

import (
	e "errors"
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

var MinBit = 256
var MaxBit = 1024 * 1024 * 1024

func checkBit(b int) error {
	b *= 8
	if b < MinBit || b > MaxBit || b%16 != 0 {
		return e.New("invalid bit size")
	}

	return nil
}

func Shake128(data []byte, b int) ([]byte, error) {
	err := checkBit(b)
	if err != nil {
		return nil, err
	}

	hash := make([]byte, b)

	sha3.ShakeSum128(hash, data)

	return hash, nil
}

func Shake256(data []byte, b int) ([]byte, error) {
	err := checkBit(b)
	if err != nil {
		return nil, err
	}

	hash := make([]byte, b)

	sha3.ShakeSum256(hash, data)

	return hash, nil
}

func ShakeHttp(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit          := content.Path2Bit(r)
	data, err, s := content.InputHttpBytes(r)

	if err == nil {
		switch bit {
		// 128
		case ShakeBits[0]:
			hash, err := Shake128(data, MinBit / 8)

			if err == nil {
				content.OutputHttpString(w, r, hex.EncodeToString(hash))
			} else {
				content.OutputHttpError(w, r, err, http.StatusUnprocessableEntity)
			}
		// 256
		case ShakeBits[1]:
			hash, err := Shake256(data, MinBit / 4)

			if err == nil {
				content.OutputHttpString(w, r, hex.EncodeToString(hash))
			} else {
				content.OutputHttpError(w, r, err, http.StatusUnprocessableEntity)
			}
		}
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
