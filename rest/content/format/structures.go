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
	Content    []byte `yaml:"Content"`
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
	Status     string `yaml:"Status"`
}
