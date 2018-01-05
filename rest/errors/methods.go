package errors

import (
	"fmt"
	"net/http"
)

func eMethodMessage(eID int, eCurrentMethod string, eMethod string) string {
	return fmt.Sprintf("Error %d: %s %s for %s",
			eID,
			eCurrentMethod,
			http.StatusText(eID),
			eMethod)
}

func IsMethodOptions(r *http.Request) bool {
	return r.Method != http.MethodOptions
}

func MethodOptions(w http.ResponseWriter, r *http.Request) bool {
	if IsMethodOptions(r) {
		http.Error(w,
			eMethodMessage(http.StatusMethodNotAllowed, r.Method, http.MethodOptions),
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}

func IsMethodHead(r *http.Request) bool {
	return r.Method != http.MethodHead
}

func MethodHead(w http.ResponseWriter, r *http.Request) bool {
	if IsMethodHead(r) {
		http.Error(w,
			eMethodMessage(http.StatusMethodNotAllowed, r.Method, http.MethodHead),
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}

func IsMethodGet(r *http.Request) bool {
	return r.Method != http.MethodGet
}

func MethodGet(w http.ResponseWriter, r *http.Request) bool {
	if IsMethodGet(r) {
		http.Error(w,
			eMethodMessage(http.StatusMethodNotAllowed, r.Method, http.MethodGet),
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}

func IsMethodPost(r *http.Request) bool {
	return r.Method != http.MethodPost
}

func MethodPost(w http.ResponseWriter, r *http.Request) bool {
	if IsMethodPost(r) {
		http.Error(w,
			eMethodMessage(http.StatusMethodNotAllowed, r.Method, http.MethodPost),
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}

func IsMethodPatch(r *http.Request) bool {
	return r.Method != http.MethodPatch
}

func MethodPatch(w http.ResponseWriter, r *http.Request) bool {
	if IsMethodPatch(r) {
		http.Error(w,
			eMethodMessage(http.StatusMethodNotAllowed, r.Method, http.MethodPatch),
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}

func IsMethodDelete(r *http.Request) bool {
	return r.Method != http.MethodDelete
}

func MethodDelete(w http.ResponseWriter, r *http.Request) bool {
	if IsMethodDelete(r) {
		http.Error(w,
			eMethodMessage(http.StatusMethodNotAllowed, r.Method, http.MethodDelete),
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}
