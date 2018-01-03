package hashes

import (
	"net/http"
	"io"
)

func Base64(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `Base64`)
}
