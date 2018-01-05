package main

import (
	"fmt"
	"path"
	"hash"
	"net/http"

	"rest/errors"
	"rest/handlers/online/hashes"
)

const RootPath = "/"

var AllPathes []string

func serverURI(uriPath string) string {
	var scheme string

	if DefaultGlobalPort == *serverPort {
		scheme = *serverHost
	} else {
		scheme = fmt.Sprintf("%s:%d", *serverHost, *serverPort)
	}

	return ServerUriSchema + scheme + uriPath
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
	fmt.Fprint(w, "Pathes:\n")
	for _, p := range AllPathes {
		fmt.Fprintf(w, "    %s\n", p)
	}

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

func initHandlerAsString(mux *http.ServeMux, httpPath string, httpFunc func(w http.ResponseWriter, r *http.Request)) {
	mux.HandleFunc(httpPath, httpFunc)
	AllPathes = append(AllPathes, httpPath)
}

func initHandlerAsArray(mux *http.ServeMux, httpPath string, httpFunc func(w http.ResponseWriter, r *http.Request), values []string) {
	for _, value := range values {
		p := path.Join(httpPath, value)

		mux.HandleFunc(p, httpFunc)
		AllPathes = append(AllPathes, p)
	}
}

func initHandlerAsHash(mux *http.ServeMux, httpPath string, httpFunc func(w http.ResponseWriter, r *http.Request), values map[string]func() hash.Hash) {
	for value := range values {
		p := path.Join(httpPath, value)

		mux.HandleFunc(p, httpFunc)
		AllPathes = append(AllPathes, p)
	}
}

func initOnlineHashesHandlers(mux *http.ServeMux) {
	initHandlerAsString(&*mux, hashes.Md2Path, hashes.MD2)
	initHandlerAsString(&*mux, hashes.Md4Path, hashes.MD4)
	initHandlerAsString(&*mux, hashes.Md5Path, hashes.MD5)
	initHandlerAsString(&*mux, hashes.Ripemd160Path, hashes.RIPEMD160)
	initHandlerAsString(&*mux, hashes.Sha1Path, hashes.SHA1)
	initHandlerAsHash(&*mux, hashes.Sha2Path, hashes.SHA2, hashes.Sha2Bits)
	initHandlerAsHash(&*mux, hashes.Sha3Path, hashes.SHA3, hashes.Sha3Bits)
	initHandlerAsHash(&*mux, hashes.KeccakPath, hashes.KECCAK, hashes.KeccakBits)
	initHandlerAsArray(&*mux, hashes.Blake2sPath, hashes.BLAKE2s, hashes.Blake2sBits)
	initHandlerAsArray(&*mux, hashes.Blake2bPath, hashes.BLAKE2b, hashes.Blake2bBits)
	initHandlerAsArray(&*mux, hashes.ShakePath, hashes.SHAKE, hashes.ShakeBits)
	initHandlerAsArray(&*mux, hashes.Base32Path, hashes.Base32, hashes.Base32Actions)
	initHandlerAsArray(&*mux, hashes.Base64Path, hashes.Base64, hashes.Base64Actions)
	initHandlerAsArray(&*mux, hashes.Crc8Path, hashes.CRC8, hashes.Crc8Types)
	initHandlerAsArray(&*mux, hashes.Crc16Path, hashes.CRC16, hashes.Crc16Types)
	initHandlerAsArray(&*mux, hashes.Crc32Path, hashes.CRC32, hashes.Crc32Types)
	initHandlerAsArray(&*mux, hashes.Crc64Path, hashes.CRC64, hashes.Crc64Types)
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
