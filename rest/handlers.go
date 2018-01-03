package main

import (
	"strings"
	"flag"
	"net/http"

	"rest/handlers"
	"rest/handlers/hashes"
)

var (
	httpsAddr = flag.String("https_addr", "localhost:64430", "TLS address to listen on ('host:port' or ':port'). Required.")
	httpAddr  = flag.String("http_addr", "", "Plain HTTP address to listen on ('host:port', or ':port'). Empty means no HTTP.")

	hostHTTP  = flag.String("http_host", "", "Optional host or host:port to use for http:// links to this service. By default, this is implied from -http_addr.")
	hostHTTPS = flag.String("https_host", "", "Optional host or host:port to use for http:// links to this service. By default, this is implied from -https_addr.")
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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.TLS == nil:
			http.Redirect(w, r, "https://"+httpsHost()+"/", http.StatusFound)

			return
		}

		mux2.ServeHTTP(w, r)
	})

	mux2.HandleFunc("/", handlers.Home)
// Online Hashes
	mux2.HandleFunc("/online/hashes/sha/1", hashes.SHA)
	mux2.HandleFunc("/online/hashes/sha/256", hashes.SHA)
	mux2.HandleFunc("/online/hashes/sha/384", hashes.SHA)
	mux2.HandleFunc("/online/hashes/sha/512", hashes.SHA)
	mux2.HandleFunc("/online/hashes/sha3/256", hashes.SHA3)
	mux2.HandleFunc("/online/hashes/sha3/384", hashes.SHA3)
	mux2.HandleFunc("/online/hashes/sha3/512", hashes.SHA3)
	mux2.HandleFunc("/online/hashes/ripemd160", hashes.RIPEMD160)
	mux2.HandleFunc("/online/hashes/blake2s/256", hashes.BLAKE2s)
	mux2.HandleFunc("/online/hashes/blake2b/256", hashes.BLAKE2b)
	mux2.HandleFunc("/online/hashes/blake2b/384", hashes.BLAKE2b)
	mux2.HandleFunc("/online/hashes/blake2b/512", hashes.BLAKE2b)
	mux2.HandleFunc("/online/hashes/base64/encode", hashes.Base64)
	mux2.HandleFunc("/online/hashes/base64/decode", hashes.Base64)
}
