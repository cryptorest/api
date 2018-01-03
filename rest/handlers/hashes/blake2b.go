package hashes

import (
	"net/http"
	"io"
)

func BLAKE2b(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `BLAKE2b`)
}
