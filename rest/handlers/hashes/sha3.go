package hashes

import (
	"net/http"
	"fmt"
	"hash"
	"golang.org/x/crypto/sha3"

	"rest/handlers"
)

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
	data := ""
	b := SHA3_Bits[bit]()

	b.Write([]byte(data))
	fmt.Fprintf(w, "%x", b.Sum(nil))
}
