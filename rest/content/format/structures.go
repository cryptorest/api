package format

type InputStructure struct {
	ConfigFile string `yaml:"ConfigFile"`
	Date       string `yaml:"Date"`
	ID         string `yaml:"ID"`
	UserID     string `yaml:"UserID"`
	Host       string `yaml:"Host"`
	Port       int    `yaml:"Port"`
	Content    string `yaml:"Content"`
}

func DefaultInputStructure(c *InputStructure) {
	c.ConfigFile = ""
	c.Date       = "server.crt"
	c.ID         = ""
	c.UserID     = ""
	c.Host       = "localhost"
	c.Port       = 64443
	c.Content    = ""
}

type OutputStructure struct {
	ConfigFile string `yaml:"ConfigFile"`
	Date       string `yaml:"Date"`
	ID         string `yaml:"ID"`
	UserID     string `yaml:"UserID"`
	Host       string `yaml:"Host"`
	Port       int    `yaml:"Port"`
	Content    string `yaml:"Content"`
}

func DefaultOutputStructure(c *OutputStructure) {
	c.ConfigFile = ""
	c.Date       = "server.crt"
	c.ID         = ""
	c.UserID     = ""
	c.Host       = "localhost"
	c.Port       = 64443
	c.Content    = ""
}
