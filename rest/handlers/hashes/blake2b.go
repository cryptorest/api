package hashes

import (
	"net/http"
	"io"
	"fmt"

	"rest/handlers"
)

func BLAKE2b(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPOST(w, r) {
		return
	}

	bit := handlers.Path2Bit(r)

	io.WriteString(w, `BLAKE2b`)
	fmt.Fprintf(w, ": %d", bit)
}
