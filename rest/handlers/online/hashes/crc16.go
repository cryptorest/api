package hashes

import (
	"net/http"
	"github.com/cryptorest/crc"

	"rest/handlers"
	"rest/handlers/online"
)

const Crc16Path string = online.HashesPath + "/crc16"

var Crc16Types = [4]string{
	"arc",
	"ccitt",
	"x25",
	"xmodem",
}

func CRC16(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	bit := handlers.Path2Bit(r)
	data := []byte("data")

	switch bit {
	case Crc16Types[0]:
		handlers.WriteUInt64(w, crc.CalculateCRC(crc.CRC16, data))
	case Crc16Types[1]:
		handlers.WriteUInt64(w, crc.CalculateCRC(crc.CCITT, data))
	case Crc16Types[2]:
		handlers.WriteUInt64(w, crc.CalculateCRC(crc.X25, data))
	case Crc16Types[3]:
		handlers.WriteUInt64(w, crc.CalculateCRC(crc.XMODEM, data))
	}
}
