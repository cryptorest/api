package hashes

import (
	"fmt"
	e "errors"
	"net/http"

	"github.com/cryptorest/keccakc"

	"rest/errors"
	"rest/content"
)

const KeccakPath = content.HashesPath + "/keccak"

var KeccakBits = []string {
	"224",
	"256",
	"384",
	"512",
}

func Keccak224(data []byte) []byte {
	b := keccak.New224()

	return hashSum(data, b)
}

func Keccak256(data []byte) []byte {
	b := keccak.New256()

	return hashSum(data, b)
}

func Keccak384(data []byte) []byte {
	b := keccak.New384()

	return hashSum(data, b)
}

func Keccak512(data []byte) []byte {
	b := keccak.New512()

	return hashSum(data, b)
}

func KeccakHttp(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, &*r) {
		return
	}

	bit          := content.Path2Bit(&*r)
	data, err, s := content.InputHttpBytes(&*r, false, false)

	if err == nil {
		var b []byte

		switch bit {
		// 224
		case KeccakBits[0]:
			b = Keccak224(data)
		// 256
		case KeccakBits[1]:
			b = Keccak256(data)
		// 384
		case KeccakBits[2]:
			b = Keccak384(data)
		// 512
		case KeccakBits[3]:
			b = Keccak512(data)
		default:
			err = e.New(fmt.Sprintf("invalid bit size %s for Keccak", bit))
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
