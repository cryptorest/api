package errors

import (
	"fmt"
	"net/http"
)

const errorActionFormatString = "Error %d: Path %s %s"

func ePathMessage(eID int, eURI string) string {
	return fmt.Sprintf(errorActionFormatString,
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
