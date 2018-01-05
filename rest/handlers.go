package main

import (
	"fmt"
	"path"
	"net/http"

	"rest/errors"
	"rest/handlers/online/hashes"
)

const RootPath = "/"

func serverURI(path string) string {
	return fmt.Sprintf("%s%s:%d%s", ServerUriSchema, *serverHost, *serverPort, path)
}

func showRequestInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RequestURI: %q\n", r.RequestURI)
	fmt.Fprintf(w, "URL: %#v\n", r.URL)
	fmt.Fprintf(w, "Body.ContentLength: %d (-1 means unknown)\n", r.ContentLength)
	fmt.Fprintf(w, "Close: %v (relevant for HTTP/1 only)\n", r.Close)
	fmt.Fprintf(w, "TLS: %#v\n", r.TLS)

	r.Header.Write(w)
}

func Root(w http.ResponseWriter, r *http.Request) {
	if errors.Path(w, r, RootPath) {
		return
	}

	if errors.IsMethodGet(r) && errors.IsMethodHead(r) {
		errors.MethodGet(w, r)

		return
	}

	showRequestInfo(w, r)
}

func initOnlineHashesHandlers(mux *http.ServeMux) {
	mux.HandleFunc(hashes.Sha1Path, hashes.SHA1)
	for bit := range hashes.Sha2Bits {
		mux.HandleFunc(path.Join(hashes.Sha2Path, bit), hashes.SHA2)
	}
	for bit := range hashes.Sha3Bits {
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
	for bit := range hashes.KeccakBits {
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

	http.HandleFunc(RootPath, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.TLS == nil:
			http.Redirect(w, r, serverURI(RootPath), http.StatusFound)

			return
		}

		mux.ServeHTTP(w, r)
	})

	mux.HandleFunc(RootPath, Root)
	initOnlineHashesHandlers(&*mux)
}
