package hashes

import (
	"net/http"
	"hash/crc32"

	"rest/errors"
	"rest/content"
)

const Crc32Path = content.HashesPath + "/crc32"

var Crc32Types = []string{
	"ieee",
	"koopman",
	"castagnoli",
}

func Crc32Ieee(data []byte) uint32 {
	return crc32.Checksum(data, crc32.MakeTable(crc32.IEEE))
}

func Crc32Koopman(data []byte) uint32 {
	return crc32.Checksum(data, crc32.MakeTable(crc32.Koopman))
}

func Crc32Castagnoli(data []byte) uint32 {
	return crc32.Checksum(data, crc32.MakeTable(crc32.Castagnoli))
}

func Crc32Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, &*r) {
		return
	}

	pType        := content.Path2Type(&*r)
	data, err, s := content.InputHttpBytes(&*r)

	if err == nil {
		switch pType {
		// IEEE
		case Crc32Types[0]:
			content.OutputHttpUInt32(w, &*r, Crc32Ieee(data))
		// Koopman
		case Crc32Types[1]:
			content.OutputHttpUInt32(w, &*r, Crc32Koopman(data))
		// Castagnoli
		case Crc32Types[2]:
			content.OutputHttpUInt32(w, &*r, Crc32Castagnoli(data))
		}
	} else {
		content.OutputHttpError(w, &*r, err, s)
	}
}
