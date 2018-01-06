package config

import (
	"flag"
)

func InitCommand() bool {
	ServerCertFile = *flag.String("certFile", DefaultCertFile, "TLS server Cert file.")
	ServerKeyFile  = *flag.String("keyFile", DefaultKeyFile, "TLS server Key file.")
	ServerHost     = *flag.String("host", DefaultHost, "IP address to listen on ('host:port'). Required.")
	ServerPort     = *flag.Int("port", DefaultPort, "Port to listen on ('host:port'). Required.")
	ServerVerbose  = *flag.Bool("verbose", false, "Verbose HTTP/2 debugging. Required.")

	flag.Parse()

	return true
}
