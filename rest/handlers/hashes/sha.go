package hashes

import (
	"net/http"
	"io"
)

func SHA(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, `SHA`)
}
