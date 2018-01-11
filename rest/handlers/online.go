package handlers

import (
	"net/http"

	"rest/content/data"
	"rest/content/hashes"
)

const OnlinePath = "/online"

func InitOnline(mux *http.ServeMux) {
	// Hashes
	initHandlerAsString(&*mux, OnlinePath + hashes.Md2Path, hashes.Md2Http)
	initHandlerAsString(&*mux, OnlinePath + hashes.Md4Path, hashes.Md4Http)
	initHandlerAsString(&*mux, OnlinePath + hashes.Md5Path, hashes.Md5Http)
	initHandlerAsString(&*mux, OnlinePath + hashes.Ripemd160Path, hashes.Ripemd160Http)
	initHandlerAsString(&*mux, OnlinePath + hashes.Base32Path, hashes.Base32Http)
	initHandlerAsString(&*mux, OnlinePath + hashes.Base64Path, hashes.Base64Http)
	initHandlerAsString(&*mux, OnlinePath + hashes.Sha1Path, hashes.Sha1Http)
	initHandlerAsArray(&*mux, OnlinePath + hashes.Sha2Path, hashes.Sha2Http, hashes.Sha2Bits)
	initHandlerAsArray(&*mux, OnlinePath + hashes.Sha3Path, hashes.Sha3Http, hashes.Sha3Bits)
	initHandlerAsArray(&*mux, OnlinePath + hashes.KeccakPath, hashes.KeccakHttp, hashes.KeccakBits)
	initHandlerAsArray(&*mux, OnlinePath + hashes.Blake2bPath, hashes.Blake2bHttp, hashes.Blake2bBits)
	initHandlerAsArray(&*mux, OnlinePath + hashes.Blake2sPath, hashes.Blake2sHttp, hashes.Blake2sBits)
	initHandlerAsArray(&*mux, OnlinePath + hashes.ShakePath, hashes.ShakeHttp, hashes.ShakeBits)
	initHandlerAsString(&*mux, OnlinePath + hashes.Crc8Path, hashes.Crc8Http)
	initHandlerAsArray(&*mux, OnlinePath + hashes.Crc8Path, hashes.Crc8Http, hashes.Crc8Types)
	initHandlerAsString(&*mux, OnlinePath + hashes.Crc16Path, hashes.Crc16Http)
	initHandlerAsArray(&*mux, OnlinePath + hashes.Crc16Path, hashes.Crc16Http, hashes.Crc16Types)
	initHandlerAsArray(&*mux, OnlinePath + hashes.Crc32Path, hashes.Crc32Http, hashes.Crc32Types)
	initHandlerAsArray(&*mux, OnlinePath + hashes.Crc64Path, hashes.Crc64Http, hashes.Crc64Types)

	// Data
	initHandlerAsArray(&*mux, OnlinePath + data.Base32Path, data.Base32Http, data.Base32Actions)
	initHandlerAsArray(&*mux, OnlinePath + data.Base64Path, data.Base64Http, data.Base64Actions)
}
