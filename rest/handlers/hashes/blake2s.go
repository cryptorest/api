package hashes

import (
	"net/http"
	"io"
)

func BLAKE2s(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `BLAKE2s`)
}
