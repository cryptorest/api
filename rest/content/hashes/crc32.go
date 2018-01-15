package hashes

import (
	"fmt"
	e "errors"
	"net/http"
	"hash/crc32"

	"rest/errors"
	"rest/content"
)

const Crc32Path = content.HashesPath + "/crc32"

const errorCrc32Message = "invalid type %s for CRC32"

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
	data, err, s := content.InputHttpBytes(&*r, false, false)

	if err == nil {
		var i uint32

		switch pType {
		// IEEE
		case Crc32Types[0]:
			i = Crc32Ieee(data)
		// Koopman
		case Crc32Types[1]:
			i = Crc32Koopman(data)
		// Castagnoli
		case Crc32Types[2]:
			i = Crc32Castagnoli(data)
		default:
			err = e.New(fmt.Sprintf(errorCrc32Message, pType))
		}

		if err == nil {
			content.OutputHttpUInt32(w, &*r, i)
		} else {
			content.OutputHttpError(w, &*r, err, http.StatusNotAcceptable)
		}
	} else {
		content.OutputHttpError(w, &*r, err, s)
	}
}
