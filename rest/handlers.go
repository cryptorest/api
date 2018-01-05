package main

import (
	"fmt"
	"path"
	"hash"
	"net/http"

	"rest/errors"
	"rest/handlers"
	"rest/data/data"
	"rest/data/hashes"
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
	initHandlerAsString(&*mux, handlers.OnlinePath + hashes.Md2Path, hashes.MD2)
	initHandlerAsString(&*mux, handlers.OnlinePath + hashes.Md4Path, hashes.MD4)
	initHandlerAsString(&*mux, handlers.OnlinePath + hashes.Md5Path, hashes.MD5)
	initHandlerAsString(&*mux, handlers.OnlinePath + hashes.Ripemd160Path, hashes.RIPEMD160)
	initHandlerAsString(&*mux, handlers.OnlinePath + hashes.Base32Path, hashes.Base32)
	initHandlerAsString(&*mux, handlers.OnlinePath + hashes.Base64Path, hashes.Base64)
	initHandlerAsString(&*mux, handlers.OnlinePath + hashes.Sha1Path, hashes.SHA1)
	initHandlerAsHash(&*mux, handlers.OnlinePath + hashes.Sha2Path, hashes.SHA2, hashes.Sha2Bits)
	initHandlerAsHash(&*mux, handlers.OnlinePath + hashes.Sha3Path, hashes.SHA3, hashes.Sha3Bits)
	initHandlerAsHash(&*mux, handlers.OnlinePath + hashes.KeccakPath, hashes.KECCAK, hashes.KeccakBits)
	initHandlerAsArray(&*mux, handlers.OnlinePath + hashes.Blake2sPath, hashes.BLAKE2s, hashes.Blake2sBits)
	initHandlerAsArray(&*mux, handlers.OnlinePath + hashes.Blake2bPath, hashes.BLAKE2b, hashes.Blake2bBits)
	initHandlerAsArray(&*mux, handlers.OnlinePath + hashes.ShakePath, hashes.SHAKE, hashes.ShakeBits)
	initHandlerAsArray(&*mux, handlers.OnlinePath + hashes.Crc8Path, hashes.CRC8, hashes.Crc8Types)
	initHandlerAsArray(&*mux, handlers.OnlinePath + hashes.Crc16Path, hashes.CRC16, hashes.Crc16Types)
	initHandlerAsArray(&*mux, handlers.OnlinePath + hashes.Crc32Path, hashes.CRC32, hashes.Crc32Types)
	initHandlerAsArray(&*mux, handlers.OnlinePath + hashes.Crc64Path, hashes.CRC64, hashes.Crc64Types)
}

func initOnlineDataHandlers(mux *http.ServeMux) {
	initHandlerAsArray(&*mux, handlers.OnlinePath + data.Base32Path, data.Base32, data.Base32Actions)
	initHandlerAsArray(&*mux, handlers.OnlinePath + data.Base64Path, data.Base64, data.Base64Actions)
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

	// Root
	mux.HandleFunc(RootPath, Root)

	// Online
	initOnlineHashesHandlers(&*mux)
	initOnlineDataHandlers(&*mux)
}
