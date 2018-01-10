package content

import (
	"io"
	"os"
	"strconv"
	"strings"
	"net/http"
	"mime/multipart"

	"rest/content/format"
	"io/ioutil"
	"log"
)

const Size24K = (1 << 10) * 24

var InputFormatFuncs = [5]func(w io.Reader, s *format.InputStructure, hr bool) error {
	format.InputText,
	format.InputJson,
	format.InputYaml,
	format.InputToml,
	format.InputXml,
}

func InputHttpMimeType(r *http.Request) string {
	return r.Header.Get(MimeKeyRequest)
}

type Input struct {
	FileExtensions []string
	HttpMimeType   string
	Reader         *http.Request
	Structure      format.InputStructure
	Format         func(w io.Reader, s *format.InputStructure, hr bool) error
}

func DefaultInputHttpFormat(i *Input) {
	i.Format       = format.InputText
	i.HttpMimeType = format.TextHttpMimeTypes[0]
}

func (i *Input) FormatFind() {
	inputHttpMimeType := i.HttpMimeType
	i.HttpMimeType     = EmptyString

	for _, mimeType := range strings.Split(inputHttpMimeType, ";") {
		for f, formatHttpMimeType := range HttpMimeTypes {
			for _, httpMimeType := range formatHttpMimeType {
				if mimeType == httpMimeType {
					i.HttpMimeType = httpMimeType
					i.Format = InputFormatFuncs[f]
					i.FileExtensions = FileExtensions[f]

					break
				}
			}
		}
	}

	if i.HttpMimeType == EmptyString {
		DefaultInputHttpFormat(&*i)
	}
}

func (i *Input) Build() []byte {
	var err error

	defer func() {
		if err != nil {
//			http.Error(w, err.Error(), status)
		}
	}()

	err = i.Reader.ParseMultipartForm(Size24K)
	if err != nil {
		i.Structure.Status = http.StatusInternalServerError

		return i.Structure.Content
	}

	for _, fheaders := range i.Reader.MultipartForm.File {
		for _, hdr := range fheaders {
			// open uploaded
			var infile multipart.File

			infile, err = hdr.Open()
			if err != nil {
				i.Structure.Status = http.StatusInternalServerError

				return i.Structure.Content
			}

			// open destination
			var outfile *os.File

			i.Structure.File = "./uploaded/" + hdr.Filename
			outfile, err = os.Create(i.Structure.File)
			if err != nil {
				i.Structure.Status = http.StatusInternalServerError

				return i.Structure.Content
			}

			// 32K buffer copy
			var written int64

			written, err = io.Copy(outfile, infile)
			if err != nil {
				i.Structure.Status = http.StatusInternalServerError

				return i.Structure.Content
			}

			i.Structure.Content = []byte(strconv.Itoa(int(written)))
//			w.Write([]byte("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written))))
		}
	}


//	err = format.InputJsonFile(&i.Structure)
//	if err != nil {
//		i.Structure.Status = http.StatusInternalServerError
//
//		return i.Structure.Content
//	}

	content, err := ioutil.ReadFile(i.Structure.File)
	if err != nil {
		log.Fatal(err)
	}
	i.Structure.Content = content
	err = os.Remove(i.Structure.File)

	return i.Structure.Content
}

var InputHttpExecute = func(r *http.Request) []byte {
	var input Input

	input.Reader            = r
	input.HttpMimeType      = InputHttpMimeType(r)
	input.Structure         = format.InputStructure{}

	input.FormatFind()

	return input.Build()
}

func InputBytes(r *http.Request) []byte {
	return InputHttpExecute(r)
}

func InputString(r *http.Request) string {
	return string(InputHttpExecute(r))
}
