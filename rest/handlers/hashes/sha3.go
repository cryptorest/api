package hashes

import (
	"net/http"
	"io"
	"fmt"

	"rest/handlers"
)

func SHA3(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPOST(w, r) {
		return
	}

	bit := handlers.Path2Bit(r)

	io.WriteString(w, `SHA3`)
	fmt.Fprintf(w, ": %d", bit)
}
