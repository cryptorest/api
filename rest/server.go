package main

import (
	"fmt"
	"log"
	"net/http"
	"golang.org/x/net/http2"

	"rest/config"
)

func main() {
	config.Init()

	var server http.Server
	server.Addr       = fmt.Sprintf("%s:%d", config.ServerHost, config.ServerPort)
	http2.VerboseLogs = config.ServerVerbose

	initHandlers()

	http2.ConfigureServer(&server, &http2.Server{})

	go func() {
		log.Fatal(server.ListenAndServeTLS(config.ServerCertFile, config.ServerKeyFile))
	}()

	select {}
}
