package hashes

import (
	"net/http"
	"github.com/cryptorest/crc8"

	"rest/errors"
	"rest/content"
)

const Crc8Path = content.HashesPath + "/crc8"

var Crc8Types = []string{
	"darc",
	"dvb-s2",
	"ebu",
	"i-code",
	"itu",
	"maxim",
	"rohc",
	"wcdma",
}

func Crc8(data []byte) uint8 {
	return crc8.Checksum(data, crc8.MakeTable(crc8.CRC8))
}

func Crc8Darc(data []byte) uint8 {
	return crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_DARC))
}

func Crc8Dvbs2(data []byte) uint8 {
	return crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_DVB_S2))
}

func Crc8Ebu(data []byte) uint8 {
	return crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_EBU))
}

func Crc8Icode(data []byte) uint8 {
	return crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_I_CODE))
}

func Crc8Itu(data []byte) uint8 {
	return crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_ITU))
}

func Crc8Maxim(data []byte) uint8 {
	return crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_MAXIM))
}

func Crc8Rohc(data []byte) uint8 {
	return crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_ROHC))
}

func Crc8Wcdma(data []byte) uint8 {
	return crc8.Checksum(data, crc8.MakeTable(crc8.CRC8_WCDMA))
}

func Crc8Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, &*r) {
		return
	}

	bit          := content.Path2Bit(&*r)
	data, err, s := content.InputHttpBytes(&*r)

	if err == nil {
		switch bit {
		// DARC
		case Crc8Types[1]:
			content.OutputHttpUInt8(w, &*r, Crc8Darc(data))
		// DVB S2
		case Crc8Types[2]:
			content.OutputHttpUInt8(w, &*r, Crc8Dvbs2(data))
		// EBU
		case Crc8Types[3]:
			content.OutputHttpUInt8(w, &*r, Crc8Ebu(data))
		// iCODE
		case Crc8Types[4]:
			content.OutputHttpUInt8(w, &*r, Crc8Icode(data))
		// ITU
		case Crc8Types[5]:
			content.OutputHttpUInt8(w, &*r, Crc8Itu(data))
		// MAXIM
		case Crc8Types[6]:
			content.OutputHttpUInt8(w, &*r, Crc8Maxim(data))
		// ROHC
		case Crc8Types[7]:
			content.OutputHttpUInt8(w, &*r, Crc8Rohc(data))
		// WCDMA
		case Crc8Types[8]:
			content.OutputHttpUInt8(w, &*r, Crc8Wcdma(data))
		default:
			content.OutputHttpUInt8(w, &*r, Crc8(data))
		}
	} else {
		content.OutputHttpError(w, &*r, err, s)
	}
}
