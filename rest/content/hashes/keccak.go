package hashes

import (
	"hash"
	"net/http"
	"github.com/cryptorest/keccakc"

	"rest/content"
	"rest/errors"
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

	bit := content.Path2Bit(r)
	bData := []byte("data")
	b := KeccakBits[bit]()

	b.Write(bData)
	content.OutputHash(w, r, b.Sum(nil))
}
