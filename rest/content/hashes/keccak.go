package hashes

import (
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
	if errors.MethodPost(w, r) {
		return
	}

	bit          := content.Path2Bit(r)
	data, err, s := content.InputHttpBytes(r)

	if err == nil {
		switch bit {
		// IEEE
		case KeccakBits[0]:
			content.OutputHttpHash(w, r, Keccak224(data))
		case KeccakBits[1]:
			content.OutputHttpHash(w, r, Keccak256(data))
		case KeccakBits[2]:
			content.OutputHttpHash(w, r, Keccak384(data))
		case KeccakBits[3]:
			content.OutputHttpHash(w, r, Keccak512(data))
		}
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
