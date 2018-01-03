package main

import (
	"fmt"
	"log"
	"flag"
	"html"
	"net/http"

	"golang.org/x/net/http2"
)

func main() {
	var srv http.Server
//	http2.VerboseLogs = true
	srv.Addr = ":64443"
	flag.BoolVar(&http2.VerboseLogs, "verbose", false, "Verbose HTTP/2 debugging.")
	flag.Parse()

	initHandlers()

	http2.ConfigureServer(&srv, &http2.Server{})

	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hey client %q\n", html.EscapeString(r.URL.Path))
		ShowRequestInfoHandler(w, r)
	})

	go func() {
		log.Fatal(srv.ListenAndServeTLS("server.crt", "server.key"))
	}()

	select {}
}

func ShowRequestInfoHandler(w http.ResponseWriter, r *http.Request) {
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
	fmt.Fprintf(w, "\nHeaders:\n")
	r.Header.Write(w)
}
