package hashes

import (
	"net/http"
	"golang.org/x/crypto/blake2s"

	"rest/utils"
	"rest/errors"
	"rest/handlers/online"
)

const Blake2sPath = online.HashesPath + "/blake2s"

var Blake2sBits = [1]string{
	"256",
}

func BLAKE2s(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit := utils.Path2Bit(r)
	data := []byte("data")

	switch bit {
	case Blake2sBits[0]:
		utils.Write32Byte(w, blake2s.Sum256(data))
	}
}
