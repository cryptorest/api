package handlers

import (
	"net/http"
	"io"
)

const HomePath string = "/"

func Home(w http.ResponseWriter, r *http.Request) {
	if ErrorPath(w, r) {
		return
	}

	if ErrorMethodGet(w, r) {
		return
	}

	w.Header().Set("Content-Type", "text/plain")
	r.Header.Write(w)

	io.WriteString(w, `HOME`)
}
