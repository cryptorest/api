package config

import (
	"fmt"
)

const (
	ServerUriSchema   = "https://"
	DefaultHost       = "localhost"
	DefaultPort       = 64443
	DefaultGlobalPort = 443
	DefaultCertFile   = "server.crt"
	DefaultKeyFile    = "server.key"
)

var (
	ServerCertFile string
	ServerKeyFile  string
	ServerHost     string
	ServerPort     int
	ServerVerbose  bool
)

func ServerURI(uriPath string) string {
	var scheme string

	if DefaultGlobalPort == ServerPort {
		scheme = ServerHost
	} else {
		scheme = fmt.Sprintf("%s:%d", ServerHost, ServerPort)
	}

	return ServerUriSchema + scheme + uriPath
}

func Init() {
	if !InitEnvironment() {
		if !InitFile() {
			InitCommand()
		}
	}
}
