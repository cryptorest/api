package hashes

import (
	"net/http"
	"hash/crc32"

	"rest/data"
	"rest/errors"
)

const Crc32Path = data.HashesPath + "/crc32"

var Crc32Types = []string{
	"ieee",
	"koopman",
	"castagnoli",
}

func CRC32(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	pType := data.Path2Type(r)
	bData := []byte("data")

	switch pType {
	case Crc32Types[0]:
		data.WriteUInt32(w, crc32.Checksum(bData, crc32.MakeTable(crc32.IEEE)))
	case Crc32Types[1]:
		data.WriteUInt32(w, crc32.Checksum(bData, crc32.MakeTable(crc32.Koopman)))
	case Crc32Types[2]:
		data.WriteUInt32(w, crc32.Checksum(bData, crc32.MakeTable(crc32.Castagnoli)))
	}
}
