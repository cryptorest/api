package hashes

import (
	"hash"
	"net/http"

	"rest/errors"
	"rest/content"
	"github.com/cryptorest/keccakc"
)

const KeccakPath = content.HashesPath + "/keccak"

var KeccakBits = map[string]func() hash.Hash{
	"224": keccak.New224,
	"256": keccak.New256,
	"384": keccak.New384,
	"512": keccak.New512,
}

func KECCAK(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit          := content.Path2Bit(r)
	data, err, s := content.InputHttpBytes(r)
	b            := KeccakBits[bit]()

	if err == nil {
		b.Write(data)
		content.OutputHttpHash(w, r, b.Sum(nil))
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
