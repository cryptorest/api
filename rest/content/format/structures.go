package format

const HumanReadablePrefix = ""
const HumanReadableIndent = "  "

type InputStructure struct {
	FileExtensions []byte `yaml:"FileExtension"`
	File           string `yaml:"File"`
	Date           string `yaml:"Date"`
	Host           string `yaml:"Host"`
	Port           int    `yaml:"Port"`
	Content        []byte `yaml:"Content"`
	Error          string `yaml:"Error"`
	Status         int    `yaml:"Status"`
}

type OutputStructure struct {
	Date       string `yaml:"Date"`
	Timestamp  int64  `yaml:"Timestamp"`
	Host       string `yaml:"Host"`
	Port       int    `yaml:"Port"`
	Content    string `yaml:"Content"`
	Error      string `yaml:"Error"`
	Status     int    `yaml:"Status"`
}
