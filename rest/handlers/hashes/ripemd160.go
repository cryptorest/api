package hashes

import (
	"net/http"
	"io"
)

func RIPEMD160(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `RIPEMD160`)
}
