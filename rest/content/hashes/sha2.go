package hashes

import (
	"fmt"
	e "errors"
	"net/http"
	"crypto/sha256"
	"crypto/sha512"

	"rest/errors"
	"rest/content"
)

const Sha2Path = content.HashesPath + "/sha2"

const errorSha2Message = "invalid bit size %s for SHA2"

var Sha2Bits = []string {
	"224",
	"256",
	"384",
	"512",
	"512-224",
	"512-256",
}

func Sha2b224(data []byte) []byte {
	b := sha256.New224()

	return hashSum(data, b)
}

func Sha2b256(data []byte) []byte {
	b := sha256.New()

	return hashSum(data, b)
}

func Sha2b384(data []byte) []byte {
	b := sha512.New384()

	return hashSum(data, b)
}

func Sha2b512(data []byte) []byte {
	b :=  sha512.New()

	return hashSum(data, b)
}

func Sha2b512b224(data []byte) []byte {
	b := sha512.New512_224()

	return hashSum(data, b)
}

func Sha2b512b256(data []byte) []byte {
	b := sha512.New512_256()

	return hashSum(data, b)
}

func Sha2Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, &*r) {
		return
	}

	bit          := content.Path2Bit(&*r)
	data, err, s := content.InputHttpBytes(&*r)

	if err == nil {
		var b []byte

		switch bit {
		// 224
		case Sha2Bits[0]:
			b = Sha2b224(data)
		// 256
		case Sha2Bits[1]:
			b = Sha2b256(data)
		// 384
		case Sha2Bits[2]:
			b = Sha2b384(data)
		// 512
		case Sha2Bits[3]:
			b = Sha2b512(data)
		// 512_224
		case Sha2Bits[4]:
			b = Sha2b512b224(data)
		// 512_256
		case Sha2Bits[5]:
			b = Sha2b512b256(data)
		default:
			err = e.New(fmt.Sprintf(errorSha2Message, bit))
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
