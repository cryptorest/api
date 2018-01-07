package main

import (
	"fmt"

	"rest/config"
)

func ServerURI(uriPath string) string {
	var scheme string

	if config.Default.GlobalPort == config.Server.Port {
		scheme = config.Server.Host
	} else {
		scheme = fmt.Sprintf("%s:%d", config.Server.Host, config.Server.Port)
	}

	return config.Server.URISchema + scheme + uriPath
}

func initConfig() {
	config.InitDefault(&config.Default)

	config.InitCommand(&config.Server)
	config.Server.URISchema = config.Default.URISchema
	config.Server.GlobalPort = config.Default.GlobalPort

	if config.Server == config.Default {
		config.InitEnvironment(&config.Server)
	}
	config.InitFile(&config.Server)

	config.Server.URISchema = config.Default.URISchema
	config.Server.GlobalPort = config.Default.GlobalPort
}
