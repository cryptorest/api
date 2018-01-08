package format

const HumanReadablePrefix = ""
const HumanReadableIndent = "  "

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
	Date       string `yaml:"Date"`
	Timestamp  int64  `yaml:"Timestamp"`
	ID         string `yaml:"ID"`
	UserID     string `yaml:"UserID"`
	Host       string `yaml:"Host"`
	Port       int    `yaml:"Port"`
	Content    string `yaml:"Content"`
	Error      string `yaml:"Error"`
}

func DefaultOutputStructure(c *OutputStructure) {
	c.Date       = ""
	c.Timestamp  = 0
	c.ID         = ""
	c.UserID     = ""
	c.Host       = "localhost"
	c.Port       = 64443
	c.Content    = ""
	c.Error      = ""
}
