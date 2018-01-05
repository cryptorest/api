package hashes

import (
	"net/http"
	"github.com/cryptorest/crc"

	"rest/data"
	"rest/errors"
	"rest/handlers/online"
)

const Crc16Path = online.HashesPath + "/crc16"

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

	bit := data.Path2Bit(r)
	bData := []byte("data")

	switch bit {
	case Crc16Types[0]:
		data.WriteUInt64(w, crc.CalculateCRC(crc.CRC16, bData))
	case Crc16Types[1]:
		data.WriteUInt64(w, crc.CalculateCRC(crc.CCITT, bData))
	case Crc16Types[2]:
		data.WriteUInt64(w, crc.CalculateCRC(crc.X25, bData))
	case Crc16Types[3]:
		data.WriteUInt64(w, crc.CalculateCRC(crc.XMODEM, bData))
	}
}
