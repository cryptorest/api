package handlers

import (
	"net/http"

	"rest/content/data"
	"rest/content/hashes"
)

const OnlinePath = "/online"

func InitOnline(mux *http.ServeMux) {
	// Hashes
	initHandlerAsString(&*mux, OnlinePath + hashes.Md2Path, hashes.MD2)
	initHandlerAsString(&*mux, OnlinePath + hashes.Md4Path, hashes.MD4)
	initHandlerAsString(&*mux, OnlinePath + hashes.Md5Path, hashes.MD5)
	initHandlerAsString(&*mux, OnlinePath + hashes.Ripemd160Path, hashes.RIPEMD160)
	initHandlerAsString(&*mux, OnlinePath + hashes.Base32Path, hashes.Base32)
	initHandlerAsString(&*mux, OnlinePath + hashes.Base64Path, hashes.Base64)
	initHandlerAsString(&*mux, OnlinePath + hashes.Sha1Path, hashes.SHA1)
	initHandlerAsHash(&*mux, OnlinePath + hashes.Sha2Path, hashes.SHA2, hashes.Sha2Bits)
	initHandlerAsHash(&*mux, OnlinePath + hashes.Sha3Path, hashes.SHA3, hashes.Sha3Bits)
	initHandlerAsHash(&*mux, OnlinePath + hashes.KeccakPath, hashes.KECCAK, hashes.KeccakBits)
	initHandlerAsArray(&*mux, OnlinePath + hashes.Blake2sPath, hashes.BLAKE2s, hashes.Blake2sBits)
	initHandlerAsArray(&*mux, OnlinePath + hashes.Blake2bPath, hashes.BLAKE2b, hashes.Blake2bBits)
	initHandlerAsArray(&*mux, OnlinePath + hashes.ShakePath, hashes.SHAKE, hashes.ShakeBits)
	initHandlerAsArray(&*mux, OnlinePath + hashes.Crc8Path, hashes.CRC8, hashes.Crc8Types)
	initHandlerAsArray(&*mux, OnlinePath + hashes.Crc16Path, hashes.CRC16, hashes.Crc16Types)
	initHandlerAsArray(&*mux, OnlinePath + hashes.Crc32Path, hashes.CRC32, hashes.Crc32Types)
	initHandlerAsArray(&*mux, OnlinePath + hashes.Crc64Path, hashes.CRC64, hashes.Crc64Types)

	// Data
	initHandlerAsArray(&*mux, OnlinePath + data.Base32Path, data.Base32, data.Base32Actions)
	initHandlerAsArray(&*mux, OnlinePath + data.Base64Path, data.Base64, data.Base64Actions)
}
