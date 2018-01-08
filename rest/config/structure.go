package config

type Structure struct {
	ConfigFile string `yaml:"ConfigFile"`
	URISchema  string
	CertFile   string `yaml:"CertFile"`
	KeyFile    string `yaml:"KeyFile"`
	Host       string `yaml:"Host"`
	Port       int    `yaml:"Port"`
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
	c.GlobalPort = 443
	c.Verbose    = false
}