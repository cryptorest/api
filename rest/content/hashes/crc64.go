package hashes

import (
	"net/http"
	"hash/crc64"

	"rest/errors"
	"rest/content"
)

const Crc64Path = content.HashesPath + "/crc64"

var Crc64Types = []string{
	"iso",
	"ecma",
}

func CRC64(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	pType        := content.Path2Type(r)
	data, err, s := content.InputHttpBytes(r)

	if err == nil {
		switch pType {
		// ISO
		case Crc64Types[0]:
			content.OutputHttpUInt64(w, r, crc64.Checksum(data, crc64.MakeTable(crc64.ISO)))
		// ECMA
		case Crc64Types[1]:
			content.OutputHttpUInt64(w, r, crc64.Checksum(data, crc64.MakeTable(crc64.ECMA)))
		}
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
