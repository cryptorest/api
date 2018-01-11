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

	b.Write(data)

	return b.Sum(nil)
}

func Keccak256(data []byte) []byte {
	b := keccak.New256()

	b.Write(data)

	return b.Sum(nil)
}

func Keccak384(data []byte) []byte {
	b := keccak.New384()

	b.Write(data)

	return b.Sum(nil)
}

func Keccak512(data []byte) []byte {
	b := keccak.New512()

	b.Write(data)

	return b.Sum(nil)
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
			content.OutputHttpByte(w, r, Keccak224(data))
		case KeccakBits[1]:
			content.OutputHttpByte(w, r, Keccak256(data))
		case KeccakBits[2]:
			content.OutputHttpByte(w, r, Keccak384(data))
		case KeccakBits[3]:
			content.OutputHttpByte(w, r, Keccak512(data))
		}
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
