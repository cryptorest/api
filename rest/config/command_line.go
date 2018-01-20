package config

import (
	"fmt"
	"flag"
)

func InitCommand(c *Structure) {
	flag.StringVar(&c.ConfigFile, "configFile", Default.ConfigFile, "Server configuration file(TOML, JSON, YAML, XML).")
	flag.StringVar(&c.UploadDir, "uploadDir", Default.UploadDir, "Upload directory. Required.")
	flag.StringVar(&c.CertFile, "certFile", Default.CertFile, "TLS server Cert file.")
	flag.StringVar(&c.KeyFile, "keyFile", Default.KeyFile, "TLS server Key file.")
	flag.StringVar(&c.Host, "host", Default.Host, "IP address to listen on ('host:port'). Required.")
	flag.IntVar(&c.Port, "port", Default.Port, "Port to listen on ('host:port'). Required.")
	flag.BoolVar(&c.Verbose, "verbose", Default.Verbose, "Verbose HTTP/2 debugging. Required.")
	flag.StringVar(&c.TmpDir, "tmpDir", Default.TmpDir, "Temporary directory. Required.")
	flag.BoolVar(&c.TemporaryUpload, "temporaryUpload", Default.TemporaryUpload, "Temporary upload file and save or not save on system. Required.")
	flag.IntVar(&c.BufferSize, "bufferSize", Default.BufferSize, fmt.Sprintf("Buffer size(in kilobytes) for reading (Min: %d, Max: %d). Required.", BufferSizeMin, BufferSizeMax))
	flag.IntVar(&c.FileSizeLimit, "fileSizeLimit", Default.FileSizeLimit, fmt.Sprintf("File size(in kilobytes) limit for upload (Min: %d, Max: unlimited). Required.", FileSizeLimitMin))
	flag.IntVar(&c.BodySizeLimit, "bodySizeLimit", Default.BodySizeLimit, fmt.Sprintf("Body size(in kilobytes) limit for upload (Min: %d, Max: unlimited). Required.", BodySizeLimitMin))

	flag.Parse()
}
