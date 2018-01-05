package hashes

import (
	"net/http"
	"hash/crc64"

	"rest/utils"
	"rest/errors"
	"rest/handlers/online"
)

const Crc64Path = online.HashesPath + "/crc64"

var Crc64Types = [2]string{
	"iso",
	"ecma",
}

func CRC64(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	pType := utils.Path2Type(r)
	data := []byte("data")

	switch pType {
	case Crc64Types[0]:
		utils.WriteUInt64(w, crc64.Checksum(data, crc64.MakeTable(crc64.ISO)))
	case Crc64Types[1]:
		utils.WriteUInt64(w, crc64.Checksum(data, crc64.MakeTable(crc64.ECMA)))
	}
}
