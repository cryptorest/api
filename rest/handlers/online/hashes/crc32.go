package hashes

import (
	"net/http"
	"hash/crc32"

	"rest/handlers"
	"rest/handlers/online"
)

const Crc32Path string = online.HashesPath + "/crc32"

var Crc32Types = [3]string{
	"ieee",
	"koopman",
	"castagnoli",
}

func CRC32(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	pType := handlers.Path2Type(r)
	data := []byte("data")

	switch pType {
	case Crc32Types[0]:
		handlers.WriteUInt32(w, crc32.Checksum(data, crc32.MakeTable(crc32.IEEE)))
	case Crc32Types[1]:
		handlers.WriteUInt32(w, crc32.Checksum(data, crc32.MakeTable(crc32.Koopman)))
	case Crc32Types[2]:
		handlers.WriteUInt32(w, crc32.Checksum(data, crc32.MakeTable(crc32.Castagnoli)))
	}
}
