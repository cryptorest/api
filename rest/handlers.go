package main

import (
	"net/http"

	"rest/handlers"
	"rest/handlers/hashes"
)

func initHandlers() {
	mux2 := http.NewServeMux()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
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
