package config

import (
	"fmt"
)

type configuration struct {
	UriSchema  string
	CertFile   string
	KeyFile    string
	Host       string
	Port       int
	GlobalPort int
	Verbose    bool
}

var Server  configuration
var Default configuration

func initDefault(c *configuration) {
	c.UriSchema  = "https://"
	c.CertFile   = "server.crt"
	c.KeyFile    = "server.key"
	c.Host       = "localhost"
	c.Port       = 64443
	c.GlobalPort = 443
	c.Verbose    = false
}

func ServerURI(uriPath string) string {
	var scheme string

	if Default.GlobalPort == Server.Port {
		scheme = Server.Host
	} else {
		scheme = fmt.Sprintf("%s:%d", Server.Host, Server.Port)
	}

	return Server.UriSchema + scheme + uriPath
}

func Init() {
	initDefault(&Default)

	if !InitEnvironment(&Server) {
		if !InitFile(&Server) {
			InitCommand(&Server)
		}
	}

	Server.UriSchema = Default.UriSchema
	Server.GlobalPort = Default.GlobalPort
}
