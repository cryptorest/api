package hashes

import (
	"fmt"
	e "errors"
	"net/http"
	"encoding/hex"
	"golang.org/x/crypto/sha3"

	"rest/errors"
	"rest/content"
)

const ShakePath = content.HashesPath + "/shake"

const errorShakeMessage = "invalid bit size %s for SHAKE"

var ShakeBits = []string{
	"128",
	"256",
}

var MinBit = 256
var MaxBit = 1024 * 1024 * 1024

func checkBit(b int) error {
	b *= 8
	if b < MinBit || b > MaxBit || b%16 != 0 {
		return e.New(fmt.Sprintf(errorShakeMessage, b))
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
	if errors.MethodPost(w, &*r) {
		return
	}

	bit          := content.Path2Bit(&*r)
	data, err, s := content.InputHttpBytes(&*r, false, false)

	if err == nil {
		switch bit {
		// 128
		case ShakeBits[0]:
			hash, err := Shake128(data, MinBit / 8)

			if err == nil {
				content.OutputHttpString(w, &*r, hex.EncodeToString(hash))
			} else {
				content.OutputHttpError(w, &*r, err, http.StatusUnprocessableEntity)
			}
		// 256
		case ShakeBits[1]:
			hash, err := Shake256(data, MinBit / 4)

			if err == nil {
				content.OutputHttpString(w, &*r, hex.EncodeToString(hash))
			} else {
				content.OutputHttpError(w, &*r, err, http.StatusUnprocessableEntity)
			}
		default:
			err = e.New(fmt.Sprintf(errorShakeMessage, bit))
		}

		if err != nil {
			content.OutputHttpError(w, &*r, err, http.StatusNotAcceptable)
		}
	} else {
		content.OutputHttpError(w, &*r, err, s)
	}
}
