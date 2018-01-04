package hashes

import (
	"net/http"
	"hash/crc64"

	"rest/handlers"
	"rest/handlers/online"
)

const Crc64Path string = online.HashesPath + "/crc64"

var Crc64Types = [2]string{
	"iso",
	"ecma",
}

func CRC64(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	pType := handlers.Path2Type(r)
	data := []byte("data")

	switch pType {
	case Crc64Types[0]:
		handlers.WriteUInt64(w, crc64.Checksum(data, crc64.MakeTable(crc64.ISO)))
	case Crc64Types[1]:
		handlers.WriteUInt64(w, crc64.Checksum(data, crc64.MakeTable(crc64.ECMA)))
	}
}
