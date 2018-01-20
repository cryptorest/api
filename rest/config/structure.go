package config

import (
	"os"
	"fmt"
	"log"

	"rest/content"
)

const DirectoryPermisson = 0700

const BufferSizeBlock   = 1024
const BufferSizeMin     = 1
const BufferSizeMax     = BufferSizeBlock
const DefaultBufferSize = 4

const FileSizeLimitMin     = 0
const DefaultFileSizeLimit = BufferSizeBlock * BufferSizeBlock

const BodySizeLimitMin     = 0
const DefaultBodySizeLimit = 64 * BufferSizeBlock

type Structure struct {
	GlobalPort      int
	URISchema       string
	ConfigFile      string `yaml:"ConfigFile"`
	CertFile        string `yaml:"CertFile"`
	KeyFile         string `yaml:"KeyFile"`
	Host            string `yaml:"Host"`
	Port            int    `yaml:"Port"`
	UploadDir       string `yaml:"UploadDir"`
	TmpDir          string `yaml:"TmpDir"`
	Verbose         bool   `yaml:"Verbose"`
	TemporaryUpload bool   `yaml:"TemporaryUpload"`
	BufferSize      int    `yaml:"BufferSize"`
	FileSizeLimit   int    `yaml:"FileSizeLimit"`
	BodySizeLimit   int    `yaml:"BobySizeLimit"`
}

var Server  Structure
var Default Structure

func InitDefault(c *Structure) {
	c.ConfigFile      = ""
	c.URISchema       = "https://"
	c.CertFile        = "server.crt"
	c.KeyFile         = "server.key"
	c.Host            = "localhost"
	c.Port            = 64443
	c.UploadDir       = "./upload"
	c.TmpDir          = "./tmp"
	c.GlobalPort      = 443
	c.Verbose         = false
	c.TemporaryUpload = true
	c.BufferSize      = DefaultBufferSize
	c.FileSizeLimit   = DefaultFileSizeLimit
	c.BodySizeLimit   = DefaultBodySizeLimit
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
	Server.TmpDir     = Default.TmpDir

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

	if Server.TemporaryUpload {
		DirectoryCreate(Server.TmpDir, "Temporary upload directory")
	}

	content.Config.TemporaryUpload = Server.TemporaryUpload
	content.Config.UploadDir       = &Server.UploadDir
	content.Config.TmpDir          = &Server.TmpDir
	content.Config.BufferSize      = Server.BufferSize * BufferSizeBlock
	content.Config.FileSizeLimit   = int64(Server.FileSizeLimit * BufferSizeBlock)
	content.Config.BodySizeLimit   = int64(Server.BodySizeLimit * BufferSizeBlock)
}
