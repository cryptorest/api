package hashes

import (
	"net/http"
	"hash/crc32"

	"rest/content"
	"rest/errors"
)

const Crc32Path = content.HashesPath + "/crc32"

var Crc32Types = []string{
	"ieee",
	"koopman",
	"castagnoli",
}

func CRC32(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	pType := content.Path2Type(r)
	data  := content.InputBytes(r)

	switch pType {
	case Crc32Types[0]:
		content.OutputUInt32(w, r, crc32.Checksum(data, crc32.MakeTable(crc32.IEEE)))
	case Crc32Types[1]:
		content.OutputUInt32(w, r, crc32.Checksum(data, crc32.MakeTable(crc32.Koopman)))
	case Crc32Types[2]:
		content.OutputUInt32(w, r, crc32.Checksum(data, crc32.MakeTable(crc32.Castagnoli)))
	}
}
