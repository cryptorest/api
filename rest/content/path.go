package content

import (
	"strconv"
	"net/http"
	"path/filepath"
)

func Path2Bits(r *http.Request) int {
	strByte := filepath.Base(r.URL.Path)

	i, err := strconv.Atoi(strByte)

	if err == nil {
		return i
	}

	return 0
}

func Path2Bit(r *http.Request) string {
	return filepath.Base(r.URL.Path)
}

func Path2Action(r *http.Request) string {
	return filepath.Base(r.URL.Path)
}

func Path2Type(r *http.Request) string {
	return filepath.Base(r.URL.Path)
}
