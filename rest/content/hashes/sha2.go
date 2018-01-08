package hashes

import (
	"hash"
	"net/http"
	"crypto/sha256"
	"crypto/sha512"

	"rest/content"
	"rest/errors"
)

const Sha2Path = content.HashesPath + "/sha2"

var Sha2Bits = map[string]func() hash.Hash{
	"224": sha256.New224,
	"256": sha256.New,
	"384": sha512.New384,
	"512": sha512.New,
	"512-224": sha512.New512_224,
	"512-256": sha512.New512_256,
}

func SHA2(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit := content.Path2Bit(r)
	bData := []byte("data")
	b := Sha2Bits[bit]()

	b.Write(bData)
	content.OutputHash(w, r, b.Sum(nil))
}