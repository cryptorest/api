package handlers

import (
	"fmt"
	"strconv"
	"path/filepath"
	"net/http"
)

func Path2Bits(r *http.Request) int {
	str_bit := filepath.Base(r.URL.Path)

	i, err := strconv.Atoi(str_bit)

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

func WriteHash(w http.ResponseWriter, b []byte) {
	fmt.Fprintf(w, "%x", b)
}

func WriteBytes(w http.ResponseWriter, b []byte) {
	fmt.Fprintf(w, "%s", b)
}

func WriteString(w http.ResponseWriter, s string) {
	fmt.Fprintf(w, "%s", s)
}

func WriteUInt8(w http.ResponseWriter, i uint8) {
	fmt.Fprintf(w, "%x", i)
}

func WriteUInt32(w http.ResponseWriter, i uint32) {
	fmt.Fprintf(w, "%x", i)
}

func WriteUInt64(w http.ResponseWriter, i uint64) {
	fmt.Fprintf(w, "%x", i)
}

func WriteError(w http.ResponseWriter, e error) {
	fmt.Fprintf(w, "Error: %s", e)
}
