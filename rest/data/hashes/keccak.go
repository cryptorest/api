package hashes

import (
	"hash"
	"net/http"
	"github.com/cryptorest/keccakc"

	"rest/data"
	"rest/errors"
)

const KeccakPath = data.HashesPath + "/keccak"

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

	bit := data.Path2Bit(r)
	bData := []byte("data")
	b := KeccakBits[bit]()

	b.Write(bData)
	data.WriteHash(w, b.Sum(nil))
}
