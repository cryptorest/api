package main

import (
	"path"
	"strings"
	"flag"
	"net/http"

	"rest/handlers"
	"rest/handlers/online/hashes"
)

var (
	httpsAddr = flag.String("https_addr", "localhost:64430", "TLS address to listen on ('host:port' or ':port'). Required.")
	httpAddr  = flag.String("http_addr", "", "Plain HTTP address to listen on ('host:port', or ':port'). Empty means no HTTP.")

	hostHTTP  = flag.String("http_host", "", "Optional host or host:port to use for http:// links to this service. By default, this is implied from -http_addr.")
	hostHTTPS = flag.String("https_host", "", "Optional host or host:port to use for https:// links to this service. By default, this is implied from -https_addr.")
)

func httpsHost() string {
	if *hostHTTPS != "" {
		return *hostHTTPS
	}

	if v := *httpsAddr; strings.HasPrefix(v, ":") {
		return "localhost" + v
	} else {
		return v
	}
}

func httpHost() string {
	if *hostHTTP != "" {
		return *hostHTTP
	}

	if v := *httpAddr; strings.HasPrefix(v, ":") {
		return "localhost" + v
	} else {
		return v
	}
}

func initHandlers() {
	mux2 := http.NewServeMux()

	http.HandleFunc(handlers.HomePath, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.TLS == nil:
			http.Redirect(w, r, "https://"+httpsHost()+handlers.HomePath, http.StatusFound)

			return
		}

		mux2.ServeHTTP(w, r)
	})

	mux2.HandleFunc(handlers.HomePath, handlers.Home)

	// Online Hashes
	mux2.HandleFunc(hashes.Sha1Path, hashes.SHA1)
	for key := range hashes.Sha2Bits {
		mux2.HandleFunc(path.Join(hashes.Sha2Path, key), hashes.SHA2)
	}
	for key := range hashes.Sha3Bits {
		mux2.HandleFunc(path.Join(hashes.Sha3Path, key), hashes.SHA3)
	}
	for i := range hashes.Blake2sBits {
		mux2.HandleFunc(path.Join(hashes.Blake2sPath, hashes.Blake2sBits[i]), hashes.BLAKE2s)
	}
	for i := range hashes.Blake2bBits {
		mux2.HandleFunc(path.Join(hashes.Blake2bPath, hashes.Blake2bBits[i]), hashes.BLAKE2b)
	}
	for i := range hashes.ShakeBits {
		mux2.HandleFunc(path.Join(hashes.ShakePath, hashes.ShakeBits[i]), hashes.SHAKE)
	}
	for i := range hashes.Base32Actions {
		mux2.HandleFunc(path.Join(hashes.Base32Path, hashes.Base64Actions[i]), hashes.Base32)
	}
	for i := range hashes.Base64Actions {
		mux2.HandleFunc(path.Join(hashes.Base64Path, hashes.Base64Actions[i]), hashes.Base64)
	}
	for key := range hashes.KeccakBits {
		mux2.HandleFunc(path.Join(hashes.KeccakPath, key), hashes.KECCAK)
	}
	for i := range hashes.Crc8Types {
		mux2.HandleFunc(path.Join(hashes.Crc8Path, hashes.Crc8Types[i]), hashes.CRC8)
	}
	for i := range hashes.Crc16Types {
		mux2.HandleFunc(path.Join(hashes.Crc16Path, hashes.Crc16Types[i]), hashes.CRC16)
	}
	for i := range hashes.Crc32Types {
		mux2.HandleFunc(path.Join(hashes.Crc32Path, hashes.Crc32Types[i]), hashes.CRC32)
	}
	for i := range hashes.Crc64Types {
		mux2.HandleFunc(path.Join(hashes.Crc64Path, hashes.Crc64Types[i]), hashes.CRC64)
	}
	mux2.HandleFunc(hashes.Ripemd160Path, hashes.RIPEMD160)
	mux2.HandleFunc(hashes.Md4Path, hashes.MD4)
	mux2.HandleFunc(hashes.Md5Path, hashes.MD5)
}
