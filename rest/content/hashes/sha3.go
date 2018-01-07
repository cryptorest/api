package hashes

import (
	"hash"
	"net/http"
	"golang.org/x/crypto/sha3"

	"rest/content"
	"rest/errors"
)

const Sha3Path = content.HashesPath + "/sha3"

var Sha3Bits = map[string]func() hash.Hash{
	"224": sha3.New224,
	"256": sha3.New256,
	"384": sha3.New384,
	"512": sha3.New512,
}

func SHA3(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit := content.Path2Bit(r)
	bData := []byte("data")
	b := Sha3Bits[bit]()

	b.Write(bData)
	content.WriteHash(w, b.Sum(nil))
}
