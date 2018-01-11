package hashes

import (
	"net/http"
	"github.com/cryptorest/crc8"

	"rest/errors"
	"rest/content"
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

	bit          := content.Path2Bit(r)
	data, err, s := content.InputHttpBytes(r)

	if err == nil {
		switch bit {
		// ARC
		case Crc8Types[0]:
			content.OutputHttpUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8)))
		// DARC
		case Crc8Types[1]:
			content.OutputHttpUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_DARC)))
		// DVB S2
		case Crc8Types[2]:
			content.OutputHttpUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_DVB_S2)))
		// EBU
		case Crc8Types[3]:
			content.OutputHttpUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_EBU)))
		// iCODE
		case Crc8Types[4]:
			content.OutputHttpUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_I_CODE)))
		// ITU
		case Crc8Types[5]:
			content.OutputHttpUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_ITU)))
		// MAXIM
		case Crc8Types[6]:
			content.OutputHttpUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_MAXIM)))
		// ROHC
		case Crc8Types[7]:
			content.OutputHttpUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_ROHC)))
		// WCDMA
		case Crc8Types[8]:
			content.OutputHttpUInt8(w, r, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_WCDMA)))
		}
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
