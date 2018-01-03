package handlers

import (
	"net/http"
)

func ErrorMethodPOST(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != "POST" {
		http.Error(w, "ERROR: Method POST required", 400)

		return true
	}

	return false
}
