package handlers

import (
	"fmt"
	"net/http"
)

func errorMessage(eID int, eCurrentMethod string, eMethod string) string {
	return fmt.Sprintf("Error %d: %s %s for %s",
			eID,
			eCurrentMethod,
			http.StatusText(eID),
			eMethod)
}

func ErrorMethodGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodGet {
		http.Error(w,
			errorMessage(http.StatusMethodNotAllowed, r.Method, http.MethodGet),
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}

func ErrorMethodPost(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		http.Error(w,
			errorMessage(http.StatusMethodNotAllowed, r.Method, http.MethodPost),
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}

func ErrorMethodPatch(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPatch {
		http.Error(w,
			errorMessage(http.StatusMethodNotAllowed, r.Method, http.MethodPatch),
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}

func ErrorMethodDelete(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodDelete {
		http.Error(w,
			errorMessage(http.StatusMethodNotAllowed, r.Method, http.MethodDelete),
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}
