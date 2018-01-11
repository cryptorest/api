package hashes

import (
	"net/http"
	"golang.org/x/crypto/sha3"

	"rest/errors"
	"rest/content"
)

const Sha3Path = content.HashesPath + "/sha3"

var Sha3Bits = []string {
	"224",
	"256",
	"384",
	"512",
}

func Sha3b224(data []byte) []byte {
	b := sha3.New224()

	b.Write(data)

	return b.Sum(nil)
}

func Sha3b256(data []byte) []byte {
	b := sha3.New256()

	b.Write(data)

	return b.Sum(nil)
}

func Sha3b384(data []byte) []byte {
	b := sha3.New384()

	b.Write(data)

	return b.Sum(nil)
}

func Sha3b512(data []byte) []byte {
	b := sha3.New512()

	b.Write(data)

	return b.Sum(nil)
}

func Sha3Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit          := content.Path2Bit(r)
	data, err, s := content.InputHttpBytes(r)

	if err == nil {
		switch bit {
		// 224
		case Sha3Bits[0]:
			content.OutputHttpByte(w, r, Sha3b224(data))
			// 256
		case Sha3Bits[1]:
			content.OutputHttpByte(w, r, Sha3b256(data))
			// 384
		case Sha3Bits[2]:
			content.OutputHttpByte(w, r, Sha3b384(data))
			// 512
		case Sha3Bits[3]:
			content.OutputHttpByte(w, r, Sha3b512(data))
		}
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
