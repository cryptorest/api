package hashes

import (
	"net/http"
	"io"
)

func SHA3(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `SHA3`)
}
