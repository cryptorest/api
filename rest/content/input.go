package content

import (
	"io"
	"os"
//	"bytes"
	"errors"
	"strconv"
	"strings"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"mime/multipart"

	"rest/config"
	"rest/content/format"
)

const Size24K = (1 << 10) * 24
const BufferSizeMin = 1024
const BufferSizeMax = 256 * BufferSizeMin

func InputHttpMimeType(r *http.Request) string {
	return r.Header.Get(HttpMimeTypeInputKey)
}

func DefaultInputHttpFormat(i *Input) {
	i.Format       = &format.Text
	i.HttpMimeType = format.TextHttpMimeTypes[0]
}

type Input struct {
	HttpMimeType string
	Reader       *http.Request
	Structure    *format.InputStructure
	Format       *format.Structure
}

func (i *Input) FormatFind() {
	inputHttpMimeType := i.HttpMimeType
	i.HttpMimeType     = EmptyString

	for _, mimeType := range strings.Split(inputHttpMimeType, HttpMimeTypeSeparator) {
		if mimeType == HttpMimeTypeInputFile {
			i.HttpMimeType = mimeType
			i.Format       = &format.Text

			return
		}

		for _, f := range &Formats {
			for _, httpMimeType := range *f.MimeTypes {
				if mimeType == httpMimeType {
					i.HttpMimeType = httpMimeType
					i.Format       = nil

					return
				}
			}
		}
	}

	if i.HttpMimeType == EmptyString {
		DefaultInputHttpFormat(&*i)
	}
}

func (i *Input) Size() error {
	var s   int
	var err error

	size := i.Reader.Header.Get(HttpMimeTypeInputSize)

	if size == EmptyString {
		i.Structure.ContentSize = 0

		return errors.New("content size 0")
	}

	s, err = strconv.Atoi(size)

	if err != nil {
		i.Structure.ContentSize = 0
	} else {
		i.Structure.ContentSize = int64(s)
	}

	return err
}

func (i *Input) FileRead() error {
	var err error

	defer func() error {
		return err
	}()

	err = i.Reader.ParseMultipartForm(Size24K)
	if err != nil {
		i.Structure.Status = http.StatusNotAcceptable
		i.Structure.Error  = err.Error()

		return err
	}

	for _, fileHeaders := range i.Reader.MultipartForm.File {
		for _, header := range fileHeaders {
			// open uploaded
			var inFile multipart.File

			inFile, err = header.Open()
//			defer inFile.Close()
			if err != nil {
				i.Structure.Status = http.StatusInternalServerError
				i.Structure.Error  = err.Error()

				return err
			}

			// open destination
			var outFile *os.File

			i.Structure.File = filepath.Join(config.Server.UploadDir, header.Filename)
			outFile, err = os.Create(i.Structure.File)
//			defer outFile.Close()
			if err != nil {
				i.Structure.Status = http.StatusInternalServerError
				i.Structure.Error  = "upload file error on system"

				return errors.New(i.Structure.Error)
			}

			// 32K buffer copy
//			var written int64

			_, err = io.Copy(outFile, inFile)
			if err != nil {
				i.Structure.Status = http.StatusInternalServerError
				i.Structure.Error  = err.Error()

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

//		err = os.Remove(i.Structure.File)
//		if err != nil {
//			i.Structure.Status = http.StatusInternalServerError
//			i.Structure.Error  = err.Error()
//
//			return err
//		}
	} else {
		i.Structure.Status = http.StatusInternalServerError
		i.Structure.Error  = err.Error()

		return err
	}

	return nil
}

func (i *Input) BodyRead() error {
	defer i.Reader.Body.Close()

	part, err := ioutil.ReadAll(io.LimitReader(i.Reader.Body, i.Structure.ContentSize))
	if err != nil {
		return err
	}

	i.Structure.Content = part
//	i.Reader.Body       = ioutil.NopCloser(io.MultiReader(bytes.NewReader(part), i.Reader.Body))

	return nil
}

func (i *Input) Build() ([]byte, error, int) {
	err := i.Size()

	if err == nil {
		if i.HttpMimeType == HttpMimeTypeInputFile {
			err = i.FileRead()
		} else {
			err = i.BodyRead()
		}
	}

	return i.Structure.Content, err, i.Structure.Status
}

var InputHttpExecute = func(r *http.Request) ([]byte, error, int) {
	var input Input

	input.Reader       = &*r
	input.HttpMimeType = InputHttpMimeType(&*r)
	input.Structure    = &format.InputStructure{}

	input.FormatFind()

	return input.Build()
}

func InputHttpBytes(r *http.Request) ([]byte, error, int) {
	return InputHttpExecute(&*r)
}

func InputHttpString(r *http.Request) (string, error, int) {
	i, err, s := InputHttpExecute(&*r)

	return string(i), err, s
}
