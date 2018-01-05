package hashes

import (
	"net/http"
	"golang.org/x/crypto/blake2s"

	"rest/data"
	"rest/errors"
)

const Blake2sPath = data.HashesPath + "/blake2s"

var Blake2sBits = []string{
	"256",
}

func BLAKE2s(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit := data.Path2Bit(r)
	bData := []byte("data")

	switch bit {
	case Blake2sBits[0]:
		data.Write32Byte(w, blake2s.Sum256(bData))
	}
}
