package handlers

import (
	"strconv"
	"path/filepath"
	"net/http"
)

func Path2Bit (r *http.Request) int {
	str_bit := filepath.Base(r.URL.Path)

	i, err := strconv.Atoi(str_bit)

	if err == nil {
		return i
	}

	return 0
}

func Path2Action (r *http.Request) string {
	return filepath.Base(r.URL.Path)
}
