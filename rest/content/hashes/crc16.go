package hashes

import (
	"net/http"
	"github.com/cryptorest/crc"

	"rest/errors"
	"rest/content"
)

const Crc16Path = content.HashesPath + "/crc16"

var Crc16Types = []string{
	"ccitt",
	"x25",
	"xmodem",
}

func Crc16(data []byte) uint64 {
	return crc.CalculateCRC(crc.CRC16, data)
}

func Crc16Ccitt(data []byte) uint64 {
	return crc.CalculateCRC(crc.CCITT, data)
}

func Crc16X25(data []byte) uint64 {
	return crc.CalculateCRC(crc.X25, data)
}

func Crc16Xmodem(data []byte) uint64 {
	return crc.CalculateCRC(crc.XMODEM, data)
}

func Crc16Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, &*r) {
		return
	}

	bit          := content.Path2Bit(&*r)
	data, err, s := content.InputHttpBytes(&*r, false, false)

	if err == nil {
		var i uint64

		switch bit {
		// CCITT
		case Crc16Types[0]:
			i = Crc16Ccitt(data)
		// X25
		case Crc16Types[1]:
			i = Crc16X25(data)
		// XMODEM
		case Crc16Types[2]:
			i = Crc16Xmodem(data)
		default:
			i = Crc16(data)
		}

		content.OutputHttpUInt64(w, &*r, i)
	} else {
		content.OutputHttpError(w, &*r, err, s)
	}
}
