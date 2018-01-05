package main

import (
	"fmt"
	"log"
	"flag"
	"net/http"

	"golang.org/x/net/http2"
)

const (
	ServerUriSchema = "https://"
	DefaultHost     = "localhost"
	DefaultPort     = 64443
	DefaultCertFile = "server.crt"
	DefaultKeyFile  = "server.key"
)

var (
	serverCertFile = flag.String("certFile", DefaultCertFile, "TLS server Cert file.")
	serverKeyFile  = flag.String("keyFile", DefaultKeyFile, "TLS server Key file.")
	serverHost     = flag.String("host", DefaultHost, "IP address to listen on ('host:port'). Required.")
	serverPort     = flag.Int("port", DefaultPort, "Port to listen on ('host:port'). Required.")
	serverVerbose  = flag.Bool("verbose", false, "Verbose HTTP/2 debugging. Required.")
)

func main() {
	flag.Parse()

	var server http.Server
	server.Addr = fmt.Sprintf("%s:%d", *serverHost, *serverPort)
	http2.VerboseLogs = *serverVerbose

	initHandlers()

	http2.ConfigureServer(&server, &http2.Server{})

	go func() {
		log.Fatal(server.ListenAndServeTLS(*serverCertFile, *serverKeyFile))
	}()

	select {}
}
