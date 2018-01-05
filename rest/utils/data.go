package utils

import (
	"fmt"
	"net/http"
)

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

func Write32Byte(w http.ResponseWriter, b [32]byte) {
	fmt.Fprintf(w, "%x", b)
}

func Write48Byte(w http.ResponseWriter, b [48]byte) {
	fmt.Fprintf(w, "%x", b)
}

func Write64Byte(w http.ResponseWriter, b [64]byte) {
	fmt.Fprintf(w, "%x", b)
}

func WriteError(w http.ResponseWriter, e error) {
	fmt.Fprintf(w, "Error: %s", e)
}
