package hashes

import (
	"hash"
	"net/http"
	"github.com/cryptorest/keccakc"

	"rest/handlers"
	"rest/handlers/online"
)

const KeccakPath string = online.HashesPath + "/keccak"

var KeccakBits = map[string]func() hash.Hash{
	"224": keccak.New224,
	"256": keccak.New256,
	"384": keccak.New384,
	"512": keccak.New512,
}

func KECCAK(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	bit := handlers.Path2Bit(r)
	data := []byte("data")
	b := KeccakBits[bit]()

	b.Write(data)
	handlers.WriteHash(w, b.Sum(nil))
}
