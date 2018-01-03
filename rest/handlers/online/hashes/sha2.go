package hashes

import (
	"net/http"
	"hash"
	"crypto/sha256"
	"crypto/sha512"

	"rest/handlers"
	"rest/handlers/online"
)

const SHA2Path string = online.HashesPath + "/sha2"

var SHA2Bits = map[string]func() hash.Hash{
	"256": sha256.New,
	"384": sha512.New384,
	"512": sha512.New,
}

func SHA2(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	bit := handlers.Path2Bit(r)
	data := []byte("data")
	b := SHA2Bits[bit]()

	b.Write(data)
	handlers.WriteHash(w, b.Sum(nil))
}
