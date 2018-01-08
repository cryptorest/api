package hashes

import (
	"net/http"
	"github.com/cryptorest/crc"

	"rest/content"
	"rest/errors"
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

	bit := content.Path2Bit(r)
	bData := []byte("data")

	switch bit {
	case Crc16Types[0]:
		content.OutputUInt64(w, r, crc.CalculateCRC(crc.CRC16, bData))
	case Crc16Types[1]:
		content.OutputUInt64(w, r, crc.CalculateCRC(crc.CCITT, bData))
	case Crc16Types[2]:
		content.OutputUInt64(w, r, crc.CalculateCRC(crc.X25, bData))
	case Crc16Types[3]:
		content.OutputUInt64(w, r, crc.CalculateCRC(crc.XMODEM, bData))
	}
}
