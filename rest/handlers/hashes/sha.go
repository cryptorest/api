package hashes

import (
	"net/http"
	"fmt"
	"hash"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"

	"rest/handlers"
)

var SHA_Bits = map[string]func() hash.Hash{
	"1":   sha1.New,
	"256": sha256.New,
	"384": sha512.New384,
	"512": sha512.New,
}

func SHA(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	bit := handlers.Path2Bit(r)
	data := ""
	b := SHA_Bits[bit]()

	b.Write([]byte(data))
	fmt.Fprintf(w, "%x", b.Sum(nil))
}
