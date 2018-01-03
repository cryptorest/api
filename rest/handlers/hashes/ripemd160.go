package hashes

import (
	"net/http"
	"io"
	"fmt"

	"rest/handlers"
)

func RIPEMD160(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPOST(w, r) {
		return
	}

	bit := handlers.Path2Bit(r)

	io.WriteString(w, `RIPEMD160`)
	fmt.Fprintf(w, ": %d", bit)
}
