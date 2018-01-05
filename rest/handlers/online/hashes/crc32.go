package hashes

import (
	"net/http"
	"hash/crc32"

	"rest/utils"
	"rest/errors"
	"rest/handlers/online"
)

const Crc32Path = online.HashesPath + "/crc32"

var Crc32Types = [3]string{
	"ieee",
	"koopman",
	"castagnoli",
}

func CRC32(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	pType := utils.Path2Type(r)
	data := []byte("data")

	switch pType {
	case Crc32Types[0]:
		utils.WriteUInt32(w, crc32.Checksum(data, crc32.MakeTable(crc32.IEEE)))
	case Crc32Types[1]:
		utils.WriteUInt32(w, crc32.Checksum(data, crc32.MakeTable(crc32.Koopman)))
	case Crc32Types[2]:
		utils.WriteUInt32(w, crc32.Checksum(data, crc32.MakeTable(crc32.Castagnoli)))
	}
}
