package content

import (
	"rest/content/format"
)

var MimeTypes = [][5]string {
	format.TextMimeTypes,
	format.JsonMimeTypes,
	format.YamlMimeTypes,
	format.TomlMimeTypes,
}

func Init() {
	format.DefaultInputStructure(&format.InputStructure{})
	format.DefaultOutputStructure(&format.OutputStructure{})
}
