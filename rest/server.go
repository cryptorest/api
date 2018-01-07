package main

import (
	"fmt"
	"log"
	"runtime"
	"net/http"
	"golang.org/x/net/http2"

	"rest/config"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	initConfig()

	var server http.Server
	server.Addr       = fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	http2.VerboseLogs = config.Server.Verbose

	initHandlers()

	http2.ConfigureServer(&server, &http2.Server{})

	go func() {
		log.Fatal(server.ListenAndServeTLS(config.Server.CertFile, config.Server.KeyFile))
	}()

	select {}
}
