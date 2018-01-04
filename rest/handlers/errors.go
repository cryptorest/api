package handlers

import (
	"fmt"
	"net/http"
)

func ePathMessage(eID int, eURI string) string {
	return fmt.Sprintf("Error %d: Path %s %s",
		eID,
		eURI,
		http.StatusText(eID))
}

func ErrorPath(w http.ResponseWriter, r *http.Request) bool {
	if r.URL.Path != HomePath {
		http.Error(w,
			ePathMessage(http.StatusNotFound, r.RequestURI),
			http.StatusNotFound)

		return true
	}

	return false
}

// Methods

func eMethodMessage(eID int, eCurrentMethod string, eMethod string) string {
	return fmt.Sprintf("Error %d: %s %s for %s",
			eID,
			eCurrentMethod,
			http.StatusText(eID),
			eMethod)
}

func ErrorMethodGet(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodGet {
		http.Error(w,
			eMethodMessage(http.StatusMethodNotAllowed, r.Method, http.MethodGet),
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}

func ErrorMethodPost(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		http.Error(w,
			eMethodMessage(http.StatusMethodNotAllowed, r.Method, http.MethodPost),
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}

func ErrorMethodPatch(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPatch {
		http.Error(w,
			eMethodMessage(http.StatusMethodNotAllowed, r.Method, http.MethodPatch),
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}

func ErrorMethodDelete(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodDelete {
		http.Error(w,
			eMethodMessage(http.StatusMethodNotAllowed, r.Method, http.MethodDelete),
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}
