package hashes

import (
	"net/http"
	"hash"
	"golang.org/x/crypto/sha3"

	"rest/handlers"
	"rest/handlers/online"
)

const SHA3Path string = online.HashesPath + "/sha3"

var SHA3_Bits = map[string]func() hash.Hash{
	"256": sha3.New256,
	"384": sha3.New384,
	"512": sha3.New512,
}

func SHA3(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	bit := handlers.Path2Bit(r)
	data := []byte("data")
	b := SHA3_Bits[bit]()

	b.Write(data)
	handlers.WriteHash(w, b.Sum(nil))
}