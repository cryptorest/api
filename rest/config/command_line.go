package config

import (
	"flag"
)

func InitCommand(c *configuration) bool {
	c.CertFile = *flag.String("certFile", Default.CertFile, "TLS server Cert file.")
	c.KeyFile  = *flag.String("keyFile", Default.KeyFile, "TLS server Key file.")
	c.Host     = *flag.String("host", Default.Host, "IP address to listen on ('host:port'). Required.")
	c.Port     = *flag.Int("port", Default.Port, "Port to listen on ('host:port'). Required.")
	c.Verbose  = *flag.Bool("verbose", false, "Verbose HTTP/2 debugging. Required.")

	flag.Parse()

	return true
}
