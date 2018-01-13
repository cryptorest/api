package main

import (
	"fmt"
	"log"
	"runtime"
	"net/http"
	"golang.org/x/net/http2"

	"rest/config"
)

func prcSetup() {
	n := int(runtime.NumCPU() / 2)

	if n < 1 {
		n = 1
	}

	runtime.GOMAXPROCS(n)
}

func main() {
	prcSetup()

	config.Init()

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
