package hashes

import (
	"net/http"
	"github.com/cryptorest/crc"

	"rest/utils"
	"rest/errors"
	"rest/handlers/online"
)

const Crc16Path = online.HashesPath + "/crc16"

var Crc16Types = [4]string{
	"arc",
	"ccitt",
	"x25",
	"xmodem",
}

func CRC16(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	bit := utils.Path2Bit(r)
	data := []byte("data")

	switch bit {
	case Crc16Types[0]:
		utils.WriteUInt64(w, crc.CalculateCRC(crc.CRC16, data))
	case Crc16Types[1]:
		utils.WriteUInt64(w, crc.CalculateCRC(crc.CCITT, data))
	case Crc16Types[2]:
		utils.WriteUInt64(w, crc.CalculateCRC(crc.X25, data))
	case Crc16Types[3]:
		utils.WriteUInt64(w, crc.CalculateCRC(crc.XMODEM, data))
	}
}
