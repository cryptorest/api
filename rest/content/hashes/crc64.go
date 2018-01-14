package hashes

import (
	"fmt"
	e "errors"
	"net/http"
	"hash/crc64"

	"rest/errors"
	"rest/content"
)

const Crc64Path = content.HashesPath + "/crc64"

const errorCrc64Message = "invalid type %s for CRC64"

var Crc64Types = []string{
	"iso",
	"ecma",
}

func Crc64Iso(data []byte) uint64 {
	return crc64.Checksum(data, crc64.MakeTable(crc64.ISO))
}

func Crc64Ecma(data []byte) uint64 {
	return crc64.Checksum(data, crc64.MakeTable(crc64.ISO))
}

func Crc64Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, &*r) {
		return
	}

	pType        := content.Path2Type(&*r)
	data, err, s := content.InputHttpBytes(&*r)

	if err == nil {
		var i uint64

		switch pType {
		// ISO
		case Crc64Types[0]:
			i = Crc64Iso(data)
		// ECMA
		case Crc64Types[1]:
			i = Crc64Ecma(data)
		default:
			err = e.New(fmt.Sprintf(errorCrc64Message, pType))
		}

		if err == nil {
			content.OutputHttpUInt64(w, &*r, i)
		} else {
			content.OutputHttpError(w, &*r, err, http.StatusNotAcceptable)
		}
	} else {
		content.OutputHttpError(w, &*r, err, s)
	}
}
