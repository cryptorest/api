package errors

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

func Path(w http.ResponseWriter, r *http.Request, p string) bool {
	if r.URL.Path != p {
		http.Error(w,
			ePathMessage(http.StatusNotFound, r.RequestURI),
			http.StatusNotFound)

		return true
	}

	return false
}
