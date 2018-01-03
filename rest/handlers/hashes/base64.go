package hashes

import (
	"net/http"
	"io"
	"fmt"

	"rest/handlers"
)

func Base64(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPOST(w, r) {
		return
	}

	action := handlers.Path2Action(r)

	io.WriteString(w, `Base64`)
	fmt.Fprintf(w, ": %s", action)
}
