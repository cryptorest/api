package hashes

import (
	"net/http"
	"github.com/cryptorest/crc8"

	"rest/content"
	"rest/errors"
)

const Crc8Path = content.HashesPath + "/crc8"

var Crc8Types = []string{
	"arc",
	"darc",
	"dvb-s2",
	"ebu",
	"i-code",
	"itu",
	"maxim",
	"rohc",
	"wcdma",
}

func CRC8(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit  := content.Path2Bit(r)
	data := content.InputBytes(r)

	switch bit {
	case Crc8Types[0]:
		content.OutputUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8)))
	case Crc8Types[1]:
		content.OutputUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_DARC)))
	case Crc8Types[2]:
		content.OutputUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_DVB_S2)))
	case Crc8Types[3]:
		content.OutputUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_EBU)))
	case Crc8Types[4]:
		content.OutputUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_I_CODE)))
	case Crc8Types[5]:
		content.OutputUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_ITU)))
	case Crc8Types[6]:
		content.OutputUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_MAXIM)))
	case Crc8Types[7]:
		content.OutputUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_ROHC)))
	case Crc8Types[8]:
		content.OutputUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_WCDMA)))
	}
}
