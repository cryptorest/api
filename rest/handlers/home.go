package handlers

import (
	"net/http"
	"io"
)

func Home(w http.ResponseWriter, r *http.Request) {
//	if r.URL.Path != "/" {
//		http.NotFound(w, r)

		//return
//	}

	w.Header().Set("Content-Type", "text/plain")
	r.Header.Write(w)

	io.WriteString(w, `HOME`)
}
