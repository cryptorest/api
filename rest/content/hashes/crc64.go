package hashes

import (
	"net/http"
	"hash/crc64"

	"rest/content"
	"rest/errors"
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

	pType := content.Path2Type(r)
	bData := []byte("data")

	switch pType {
	case Crc64Types[0]:
		content.WriteUInt64(w, crc64.Checksum(bData, crc64.MakeTable(crc64.ISO)))
	case Crc64Types[1]:
		content.WriteUInt64(w, crc64.Checksum(bData, crc64.MakeTable(crc64.ECMA)))
	}
}
