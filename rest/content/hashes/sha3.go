package hashes

import (
	"fmt"
	e "errors"
	"net/http"
	"golang.org/x/crypto/sha3"

	"rest/errors"
	"rest/content"
)

const Sha3Path = content.HashesPath + "/sha3"

const errorSha3Message = "invalid bit size %s for SHA3"

var Sha3Bits = []string {
	"224",
	"256",
	"384",
	"512",
}

func Sha3b224(data []byte) []byte {
	b := sha3.New224()

	return hashSum(data, b)
}

func Sha3b256(data []byte) []byte {
	b := sha3.New256()

	return hashSum(data, b)
}

func Sha3b384(data []byte) []byte {
	b := sha3.New384()

	return hashSum(data, b)
}

func Sha3b512(data []byte) []byte {
	b := sha3.New512()

	return hashSum(data, b)
}

func Sha3Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, &*r) {
		return
	}

	bit          := content.Path2Bit(&*r)
	data, err, s := content.InputHttpBytes(&*r, false, false)

	if err == nil {
		var b []byte

		switch bit {
		// 224
		case Sha3Bits[0]:
			b = Sha3b224(data)
		// 256
		case Sha3Bits[1]:
			b = Sha3b256(data)
		// 384
		case Sha3Bits[2]:
			b = Sha3b384(data)
		// 512
		case Sha3Bits[3]:
			b = Sha3b512(data)
		default:
			err = e.New(fmt.Sprintf(errorSha3Message, bit))
		}

		if err == nil {
			content.OutputHttpHash(w, &*r, b)
		} else {
			content.OutputHttpError(w, &*r, err, http.StatusNotAcceptable)
		}
	} else {
		content.OutputHttpError(w, &*r, err, s)
	}
}
