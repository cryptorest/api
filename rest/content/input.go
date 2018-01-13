package content

import (
	"io"
	"os"
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

const BufferSizeBlock = 1024

const Size24K = (1 << 10) * 24

func InputHttpMimeType(r *http.Request) string {
	return r.Header.Get(HttpMimeTypeInputKey)
}

func DefaultInputHttpFormat(i *Input) {
	i.Format       = &format.Text
	i.HttpMimeType = format.TextHttpMimeTypes[0]
}

type Input struct {
	BufferSize    int
	FileSizeLimit int
	UploadDir     string
	HttpMimeType  string
	Reader        *http.Request
	Structure     *format.InputStructure
	Format        *format.Structure
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

		return errors.New("content size is 0")
	}

	s, err = strconv.Atoi(size)

	if err != nil {
		i.Structure.ContentSize = 0
	} else {
		i.Structure.ContentSize = int64(s)
	}

	return err
}

func (i *Input) ReadBuffer(r multipart.File, w io.Writer) error {
	var n   int
	var err error

	buf := make([]byte, i.BufferSize * BufferSizeBlock)

	for {
		// read a chunk
		n, err = r.Read(buf)
		if err != nil {
			if err != io.EOF {
				return err
			} else {
				err = nil
			}
		}
		if n == 0 {
			break
		}

		// write a chunk
		_, err = w.Write(buf[:n])
		if err != nil {
			return err
		}
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
			var inFile   multipart.File
			// open destination
			var outFile *os.File

			i.Structure.File = filepath.Join(config.Server.UploadDir, header.Filename)

			inFile, err = header.Open()
			if err != nil {
				inFile.Close()
				i.Structure.Status = http.StatusInternalServerError
				i.Structure.Error  = err.Error()

				return err
			}

			outFile, err = os.Create(i.Structure.File)
			if err != nil {
				inFile.Close()
				outFile.Close()
				i.Structure.Status = http.StatusInternalServerError
				i.Structure.Error  = "upload file error on system"

				return errors.New(i.Structure.Error)
			}

			err = i.ReadBuffer(inFile, outFile)
			if err != nil {
				inFile.Close()
				outFile.Close()

				i.Structure.Status = http.StatusInternalServerError
				i.Structure.Error  = err.Error()

				return err
			}

			inFile.Close()
			outFile.Close()
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

func (i *Input) Clean() {
	i.Structure = nil
	i.Format    = nil
	i.Reader    = nil

	i.BufferSize    = 0
	i.FileSizeLimit = 0
	i.UploadDir     = EmptyString
	i.HttpMimeType  = EmptyString
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

	c := i.Structure.Content
	s := i.Structure.Status

	i.Clean()

	return c, err, s
}

var InputHttpExecute = func(r *http.Request) ([]byte, error, int) {
	var input Input

	input.BufferSize    = config.Server.BufferSize
	input.FileSizeLimit = config.Server.FileSizeLimit
	input.UploadDir     = config.Server.UploadDir
	input.Reader        = &*r
	input.HttpMimeType  = InputHttpMimeType(&*r)
	input.Structure     = &format.InputStructure{}

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
