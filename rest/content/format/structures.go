package format

import (
	"io"
)

const HumanReadablePrefix = ""
const HumanReadableIndent = "  "

type Structure struct {
	MimeTypes            *[]string
	FileExtensions       *[]string
	InputFormatFileFunc  func(s *InputStructure) error
	InputFormatFunc      func(w io.Reader, s *InputStructure, hr bool) error
	OutputFormatFunc     func(w io.Writer, s *OutputStructure, hr bool) error
}

type InputStructure struct {
	File           string `yaml:"File"`
	Date           string `yaml:"Date"`
	Host           string `yaml:"Host"`
	Port           int    `yaml:"Port"`
	Content        []byte `yaml:"Content"`
	ContentSize    int64  `yaml:"ContentSize"`
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
