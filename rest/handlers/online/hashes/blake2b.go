package hashes

import (
	"net/http"
	"github.com/golang/crypto/blake2b"

	"rest/utils"
	"rest/errors"
	"rest/handlers/online"
)

const Blake2bPath = online.HashesPath + "/blake2b"

var Blake2bBits = [3]string{
	"256",
	"384",
	"512",
}

func BLAKE2b(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit := utils.Path2Bit(r)
	data := []byte("data")

	switch bit {
	case Blake2bBits[0]:
		utils.Write32Byte(w, blake2b.Sum256(data))
	case Blake2bBits[1]:
		utils.Write48Byte(w, blake2b.Sum384(data))
	case Blake2bBits[2]:
		utils.Write64Byte(w, blake2b.Sum512(data))
	}
}
