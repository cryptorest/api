package config

import (
	"os"
	"fmt"
	"log"
)

const DirectoryPermisson = 0700

const BufferSize1K      = 1024
const BufferSizeMin     = 1
const BufferSizeMax     = BufferSize1K
const DefaultBufferSize = 16

const FileSizeLimitMin     = 0
const DefaultFileSizeLimit = BufferSize1K * BufferSize1K

type Structure struct {
	GlobalPort    int
	URISchema     string
	ConfigFile    string `yaml:"ConfigFile"`
	CertFile      string `yaml:"CertFile"`
	KeyFile       string `yaml:"KeyFile"`
	Host          string `yaml:"Host"`
	Port          int    `yaml:"Port"`
	UploadDir     string `yaml:"UploadDir"`
	Verbose       bool   `yaml:"Verbose"`
	BufferSize    int    `yaml:"BufferSize"`
	FileSizeLimit int    `yaml:"FileSizeLimit"`
}

var Server  Structure
var Default Structure

func InitDefault(c *Structure) {
	c.ConfigFile    = ""
	c.URISchema     = "https://"
	c.CertFile      = "server.crt"
	c.KeyFile       = "server.key"
	c.Host          = "localhost"
	c.Port          = 64443
	c.UploadDir     = "./upload"
	c.GlobalPort    = 443
	c.Verbose       = false
	c.BufferSize    = DefaultBufferSize
	c.FileSizeLimit = DefaultFileSizeLimit
}

func DirectoryCreate(pathDir string, title string) {
	_, err := os.Stat(pathDir)

	if os.IsNotExist(err) {
		err = os.MkdirAll(pathDir, DirectoryPermisson)
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

	if Server.BufferSize < BufferSizeMin || Server.BufferSize > BufferSizeMax {
		log.Fatalf("invalid size buffer %dK", Server.BufferSize)
	}

	if Server.FileSizeLimit < FileSizeLimitMin {
		log.Fatalf("invalid limit file size %dK", Server.BufferSize)
	}

	Server.URISchema  = Default.URISchema
	Server.GlobalPort = Default.GlobalPort

	DirectoryCreate(Server.UploadDir, "Upload directory")
}
