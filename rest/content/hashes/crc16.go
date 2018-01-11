package hashes

import (
	"net/http"
	"github.com/cryptorest/crc"

	"rest/errors"
	"rest/content"
)

const Crc16Path = content.HashesPath + "/crc16"

var Crc16Types = []string{
	"arc",
	"ccitt",
	"x25",
	"xmodem",
}

func CRC16(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit          := content.Path2Bit(r)
	data, err, s := content.InputHttpBytes(r)

	if err == nil {
		switch bit {
		// ARC
		case Crc16Types[0]:
			content.OutputHttpUInt64(w, r, crc.CalculateCRC(crc.CRC16, data))
		// CCITT
		case Crc16Types[1]:
			content.OutputHttpUInt64(w, r, crc.CalculateCRC(crc.CCITT, data))
		// X25
		case Crc16Types[2]:
			content.OutputHttpUInt64(w, r, crc.CalculateCRC(crc.X25, data))
		// XMODEM
		case Crc16Types[3]:
			content.OutputHttpUInt64(w, r, crc.CalculateCRC(crc.XMODEM, data))
		}
	} else {
		content.OutputHttpError(w, r, err, s)
	}
}
