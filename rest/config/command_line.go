package config

import (
	"flag"
)

func InitCommand(c *Configuration) {
	flag.StringVar(&c.ConfigFile, "configFile", Default.ConfigFile, "Server configuration file(TOML, JSON, YaML).")
	flag.StringVar(&c.CertFile, "certFile", Default.CertFile, "TLS server Cert file.")
	flag.StringVar(&c.KeyFile, "keyFile", Default.KeyFile, "TLS server Key file.")
	flag.StringVar(&c.Host, "host", Default.Host, "IP address to listen on ('host:port'). Required.")
	flag.IntVar(&c.Port, "port", Default.Port, "Port to listen on ('host:port'). Required.")
	flag.BoolVar(&c.Verbose, "verbose", Default.Verbose, "Verbose HTTP/2 debugging. Required.")

	flag.Parse()
}
