package content

import (
	"io"
	"os"
//	"strconv"
	"strings"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"mime/multipart"

	"rest/config"
	"rest/content/format"
	"errors"
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
					i.HttpMimeType   = httpMimeType
					i.Format         = InputFormatFuncs[f]
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

func (i *Input) Read() error {
	var err error

	defer func() error {
		i.Structure.Status = http.StatusInternalServerError

		return err
	}()

	err = i.Reader.ParseMultipartForm(Size24K)
	if err != nil {
		i.Structure.Status = http.StatusInternalServerError
		i.Structure.Error  = err.Error()

		return err
	}

	for _, fileHeaders := range i.Reader.MultipartForm.File {
		for _, header := range fileHeaders {
			// open uploaded
			var inFile multipart.File

			inFile, err = header.Open()
			if err != nil {
				i.Structure.Error = err.Error()

				return err
			}

			// open destination
			var outFile *os.File

			i.Structure.File = filepath.Join(config.Server.UploadDir, header.Filename)
			outFile, err = os.Create(i.Structure.File)
			if err != nil {
				i.Structure.Error = "Upload file error on system"

				return errors.New(i.Structure.Error)
			}

			// 32K buffer copy
//			var written int64

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				i.Structure.Error = err.Error()

				return err
			}

//			i.Structure.Content = []byte(strconv.Itoa(int(written)))
//			w.Write([]byte("uploaded file:" + hdr.Filename + ";length:" + strconv.Itoa(int(written))))
		}
	}

//	err = format.InputJsonFile(&i.Structure)
//	if err != nil {
//		i.Structure.Status = http.StatusInternalServerError
//		i.Structure.Error  = err.Error()
//
//		return
//	}

	content, err := ioutil.ReadFile(i.Structure.File)
	if err == nil {
		i.Structure.Content = content

		err = os.Remove(i.Structure.File)
		if err != nil {
			i.Structure.Status = http.StatusInternalServerError
			i.Structure.Error  = err.Error()

			return err
		}
	} else {
		i.Structure.Status = http.StatusInternalServerError
		i.Structure.Error  = err.Error()

		return err
	}

	return nil
}

func (i *Input) Build() ([]byte, error, int) {
	return i.Structure.Content, nil, i.Structure.Status
}

var InputHttpExecute = func(r *http.Request) ([]byte, error, int) {
	var input Input
	var err   error

	input.Reader       = r
	input.HttpMimeType = InputHttpMimeType(r)
	input.Structure    = format.InputStructure{}

	input.FormatFind()
	err = input.Read()

	if err == nil {
		return input.Build()
	} else {
		return input.Structure.Content, err, input.Structure.Status
	}
}

func InputHttpBytes(r *http.Request) ([]byte, error, int) {
	return InputHttpExecute(r)
}

func InputHttpString(r *http.Request) (string, error, int) {
	i, err, s := InputHttpExecute(r)

	return string(i), err, s
}
