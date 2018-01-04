package main

import (
	"flag"
	"path"
	"strings"
	"net/http"

	"rest/handlers"
	"rest/handlers/online/hashes"
)

var (
	httpsAddr = flag.String("https_addr", uriBuild(DefaultHost, DefaultPort), "TLS address to listen on ('host:port'). Required.")

	hostHTTPS = flag.String("https_host", "", "Optional host or host:port to use for https:// links to this service. By default, this is implied from -https_addr.")
)

func httpsHost() string {
	if *hostHTTPS != "" {
		return *hostHTTPS
	}

	if v := *httpsAddr; strings.HasPrefix(v, ":") {
		return DefaultHost + v
	} else {
		return v
	}
}

func initOnlineHashesHandlers(mux *http.ServeMux) {
	mux.HandleFunc(hashes.Sha1Path, hashes.SHA1)
	for bit, _ := range hashes.Sha2Bits {
		mux.HandleFunc(path.Join(hashes.Sha2Path, bit), hashes.SHA2)
	}
	for bit, _ := range hashes.Sha3Bits {
		mux.HandleFunc(path.Join(hashes.Sha3Path, bit), hashes.SHA3)
	}
	for _, bit := range hashes.Blake2sBits {
		mux.HandleFunc(path.Join(hashes.Blake2sPath, bit), hashes.BLAKE2s)
	}
	for _, bit := range hashes.Blake2bBits {
		mux.HandleFunc(path.Join(hashes.Blake2bPath, bit), hashes.BLAKE2b)
	}
	for _, bit := range hashes.ShakeBits {
		mux.HandleFunc(path.Join(hashes.ShakePath, bit), hashes.SHAKE)
	}
	for _, action := range hashes.Base32Actions {
		mux.HandleFunc(path.Join(hashes.Base32Path, action), hashes.Base32)
	}
	for _, action := range hashes.Base64Actions {
		mux.HandleFunc(path.Join(hashes.Base64Path, action), hashes.Base64)
	}
	for bit, _ := range hashes.KeccakBits {
		mux.HandleFunc(path.Join(hashes.KeccakPath, bit), hashes.KECCAK)
	}
	for _, typ := range hashes.Crc8Types {
		mux.HandleFunc(path.Join(hashes.Crc8Path, typ), hashes.CRC8)
	}
	for _, typ := range hashes.Crc16Types {
		mux.HandleFunc(path.Join(hashes.Crc16Path, typ), hashes.CRC16)
	}
	for _, typ := range hashes.Crc32Types {
		mux.HandleFunc(path.Join(hashes.Crc32Path, typ), hashes.CRC32)
	}
	for _, typ := range hashes.Crc64Types {
		mux.HandleFunc(path.Join(hashes.Crc64Path, typ), hashes.CRC64)
	}
	mux.HandleFunc(hashes.Ripemd160Path, hashes.RIPEMD160)
	mux.HandleFunc(hashes.Md2Path, hashes.MD2)
	mux.HandleFunc(hashes.Md4Path, hashes.MD4)
	mux.HandleFunc(hashes.Md5Path, hashes.MD5)
}

func initHandlers() {
	mux := http.NewServeMux()

	http.HandleFunc(handlers.HomePath, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.TLS == nil:
			http.Redirect(w, r, URISchema+httpsHost()+handlers.HomePath, http.StatusFound)

			return
		}

		mux.ServeHTTP(w, r)
	})

	mux.HandleFunc(handlers.HomePath, handlers.Home)

	initOnlineHashesHandlers(&*mux)
}
