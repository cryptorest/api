package hashes

import (
	"hash"
	"net/http"
	"golang.org/x/crypto/sha3"

	"rest/errors"
	"rest/content"
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

	bit          := content.Path2Bit(r)
	data, err, s := content.InputHttpBytes(r)
	b            := Sha3Bits[bit]()

	if err == nil {
		b.Write(data)
		content.OutputHttpHash(w, r, b.Sum(nil))
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
