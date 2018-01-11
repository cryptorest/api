package hashes

import (
	"net/http"
	"crypto/sha256"
	"crypto/sha512"

	"rest/errors"
	"rest/content"
)

const Sha2Path = content.HashesPath + "/sha2"

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

	b.Write(data)

	return b.Sum(nil)
}

func Sha2b256(data []byte) []byte {
	b := sha256.New()

	b.Write(data)

	return b.Sum(nil)
}

func Sha2b384(data []byte) []byte {
	b := sha512.New384()

	b.Write(data)

	return b.Sum(nil)
}

func Sha2b512(data []byte) []byte {
	b :=  sha512.New()

	b.Write(data)

	return b.Sum(nil)
}

func Sha2b512b224(data []byte) []byte {
	b := sha512.New512_224()

	b.Write(data)

	return b.Sum(nil)
}

func Sha2b512b256(data []byte) []byte {
	b := sha512.New512_256()

	b.Write(data)

	return b.Sum(nil)
}

func Sha2Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit          := content.Path2Bit(r)
	data, err, s := content.InputHttpBytes(r)

	if err == nil {
		switch bit {
		// 224
		case Sha2Bits[0]:
			content.OutputHttpByte(w, r, Sha2b224(data))
		// 256
		case Sha2Bits[1]:
			content.OutputHttpByte(w, r, Sha2b256(data))
		// 384
		case Sha2Bits[2]:
			content.OutputHttpByte(w, r, Sha2b384(data))
		// 512
		case Sha2Bits[3]:
			content.OutputHttpByte(w, r, Sha2b512(data))
		// 512_224
		case Sha2Bits[4]:
			content.OutputHttpByte(w, r, Sha2b512b224(data))
		// 512_256
		case Sha2Bits[5]:
			content.OutputHttpByte(w, r, Sha2b512b256(data))
		}
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
