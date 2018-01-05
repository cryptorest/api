package hashes

import (
	"hash"
	"net/http"
	"github.com/cryptorest/keccakc"

	"rest/utils"
	"rest/errors"
	"rest/handlers/online"
)

const KeccakPath = online.HashesPath + "/keccak"

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

	bit := utils.Path2Bit(r)
	data := []byte("data")
	b := KeccakBits[bit]()

	b.Write(data)
	utils.WriteHash(w, b.Sum(nil))
}
