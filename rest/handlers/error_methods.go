package handlers

import (
	"net/http"
)

func ErrorMethodGET(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodGet {
		http.Error(w,
			"ERROR: Method GET required",
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}

func ErrorMethodPOST(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPost {
		http.Error(w,
			"ERROR: Method POST required",
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}

func ErrorMethodPATCH(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodPatch {
		http.Error(w,
			"ERROR: Method PATCH required",
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}

func ErrorMethodDELETE(w http.ResponseWriter, r *http.Request) bool {
	if r.Method != http.MethodDelete {
		http.Error(w,
			"ERROR: Method DELETE required",
			http.StatusMethodNotAllowed)

		return true
	}

	return false
}
