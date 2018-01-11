package config

import (
	"os"
	"fmt"
	"log"
)

const DirectoryPermisson = 0700

type Structure struct {
	ConfigFile string `yaml:"ConfigFile"`
	URISchema  string
	CertFile   string `yaml:"CertFile"`
	KeyFile    string `yaml:"KeyFile"`
	Host       string `yaml:"Host"`
	Port       int    `yaml:"Port"`
	UploadDir  string `yaml:"UploadDir"`
	GlobalPort int
	Verbose    bool   `yaml:"Verbose"`
}

var Server  Structure
var Default Structure

func InitDefault(c *Structure) {
	c.ConfigFile = ""
	c.URISchema  = "https://"
	c.CertFile   = "server.crt"
	c.KeyFile    = "server.key"
	c.Host       = "localhost"
	c.Port       = 64443
	c.UploadDir  = "./upload"
	c.GlobalPort = 443
	c.Verbose    = false
}

func DirectoryCreate(pathDir string, title string) {
	_, err := os.Stat(pathDir)

	if os.IsNotExist(err) {
		err = os.MkdirAll(pathDir, 0000)
		if err != nil {
			log.Fatalf("%s: %v", title, err)
		}
	}
}

func ServerURI(uriPath string) string {
	var schema string

	if Default.GlobalPort == Server.Port {
		schema = Server.Host
	} else {
		schema = fmt.Sprintf("%s:%d", Server.Host, Server.Port)
	}

	return Server.URISchema + schema + uriPath
}

func Init() {
	InitDefault(&Default)

	InitCommand(&Server)
	Server.URISchema  = Default.URISchema
	Server.GlobalPort = Default.GlobalPort
	Server.UploadDir  = Default.UploadDir

	if Server == Default {
		InitEnvironment(&Server)
	}
	InitFile(&Server)

	Server.URISchema  = Default.URISchema
	Server.GlobalPort = Default.GlobalPort

	DirectoryCreate(Server.UploadDir, "Upload directory")
}
