package hashes

import (
	"net/http"
	"github.com/cryptorest/crc8"

	"rest/handlers"
	"rest/handlers/online"
)

const Crc8Path string = online.HashesPath + "/crc8"

var Crc8Types = [9]string{
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
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	bit := handlers.Path2Bit(r)
	data := []byte("data")

	switch bit {
	case Crc8Types[0]:
		handlers.WriteUInt8(w, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8)))
	case Crc8Types[1]:
		handlers.WriteUInt8(w, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_DARC)))
	case Crc8Types[2]:
		handlers.WriteUInt8(w, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_DVB_S2)))
	case Crc8Types[3]:
		handlers.WriteUInt8(w, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_EBU)))
	case Crc8Types[4]:
		handlers.WriteUInt8(w, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_I_CODE)))
	case Crc8Types[5]:
		handlers.WriteUInt8(w, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_ITU)))
	case Crc8Types[6]:
		handlers.WriteUInt8(w, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_MAXIM)))
	case Crc8Types[7]:
		handlers.WriteUInt8(w, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_ROHC)))
	case Crc8Types[8]:
		handlers.WriteUInt8(w, crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_WCDMA)))
	}
}
