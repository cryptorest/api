package content

import (
	"io"
	"os"
	"fmt"
	"time"
	"errors"
	"strconv"
	"strings"
	"net/http"
	"io/ioutil"
	"path/filepath"
	"mime/multipart"

	"github.com/dustin/randbo"

	"rest/content/format"
)

const messageErrorContentSize0 = "content size is 0 bytes"
const messageErrorFileUpload   = "upload file error on system"

var initTime = fmt.Sprint(time.Now().UnixNano())

func InputHttpMimeType(r *http.Request) string {
	return r.Header.Get(HttpMimeTypeInputKey)
}

func DefaultInputHttpFormat(i *Input) {
	i.Format       = &format.Text
	i.HttpMimeType = format.TextHttpMimeTypes[0]
}

type Input struct {
	Deploy        bool
	Parsing       bool
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

func (i *Input) FileSize() error {
	var s   int
	var err error

	size := i.Reader.Header.Get(HttpMimeTypeInputSize)

	if size == EmptyString {
		i.Structure.ContentSize = 0
		err                     = errors.New(messageErrorContentSize0)
		i.Structure.Status      = http.StatusLengthRequired
	} else {
		s, err = strconv.Atoi(size)

		if err != nil {
			i.Structure.ContentSize = 0
		} else {
			i.Structure.ContentSize = int64(s)
		}

		if i.Structure.ContentSize > Config.FileSizeLimit {
			err                = http.ErrContentLength
			i.Structure.Status = http.StatusLengthRequired
		}
	}

	return err
}

func (i *Input) BufferRead(r multipart.File, w io.Writer) error {
	var n   int
	var err error

	buf := make([]byte, Config.BufferSize)

	for {
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

		_, err = w.Write(buf[:n])

		if err != nil {
			return err
		}
	}

	buf = nil

	return err
}

func (i *Input) RandomName() string {
	r := randbo.New()
	b := []byte(initTime)

	r.Read(b)

	return fmt.Sprintf(formatHex, b)
}

func (i *Input) FilePut(fileHeader *multipart.FileHeader) error {
	var inputFile  multipart.File
	var outputFile *os.File
	var fileName   string
	var dirName    string
	var err        error

	if Config.TemporaryUpload {
		fileName = i.RandomName()
		dirName  = *Config.TmpDir
	} else {
		fileName = fileHeader.Filename
		dirName  = *Config.UploadDir
	}

	i.Structure.File = filepath.Join(dirName, fileName)
	fileName = EmptyString
	dirName  = EmptyString

	inputFile, err = fileHeader.Open()
	defer inputFile.Close()

	if err != nil {
		inputFile.Close()

		i.Structure.Status = http.StatusNotAcceptable
		i.Structure.Error  = err.Error()

		return err
	}

	outputFile, err = os.Create(i.Structure.File)
	defer outputFile.Close()

	if err != nil {
		i.Structure.Status = http.StatusInternalServerError
		i.Structure.Error  = messageErrorFileUpload

		return errors.New(i.Structure.Error)
	}

	err = i.BufferRead(inputFile, outputFile)

	if err != nil {
		i.Structure.Status = http.StatusInternalServerError
		i.Structure.Error  = err.Error()

		return err
	}

	return err
}

func (i *Input) FileContentRead() error {
	var err error

	content, err := ioutil.ReadFile(i.Structure.File)

	if err == nil {
		i.Structure.Content = content

		if Config.TemporaryUpload {
			err = os.Remove(i.Structure.File)

			if err != nil {
				i.Structure.Status = http.StatusInternalServerError
				i.Structure.Error  = err.Error()

				return err
			}
		}
	} else {
		i.Structure.Status = http.StatusInternalServerError
		i.Structure.Error  = err.Error()

		return err
	}

	return err
}

func (i *Input) FileToBufferContentRead(fileHeader *multipart.FileHeader, err error) error {
	// TODO: File read to buffer

	return err
}

func (i *Input) FileRead() error {
	var err error

	defer func() error {
		return err
	}()

	err = i.Reader.ParseMultipartForm(i.Structure.ContentSize)

	if err != nil {
		i.Structure.Status = http.StatusNotAcceptable
		i.Structure.Error  = err.Error()

		return err
	}

	for _, fileHeaders := range i.Reader.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			if int(i.Structure.ContentSize) > Config.BufferSize {
				err = i.FilePut(&*fileHeader)

				if err == nil {
					err = i.FileContentRead()
				}
			} else {
				err = i.FileToBufferContentRead(&*fileHeader, err)
			}

			if err != nil {
				return err
			}
		}
	}

	return err
}

func (i *Input) BodySize() error {
	var s   int
	var err error

	size := i.Reader.Header.Get(HttpMimeTypeInputSize)

	if size == EmptyString {
		i.Structure.ContentSize = 0
		err                     = errors.New(messageErrorContentSize0)
		i.Structure.Status      = http.StatusLengthRequired
	} else {
		s, err = strconv.Atoi(size)

		if err != nil {
			i.Structure.ContentSize = 0
		} else {
			i.Structure.ContentSize = int64(s)
		}

		if i.Structure.ContentSize > Config.BodySizeLimit {
			err                = http.ErrContentLength
			i.Structure.Status = http.StatusRequestEntityTooLarge
		}
	}

	return err
}

func (i *Input) BodyRead() error {
	part, err := ioutil.ReadAll(io.LimitReader(i.Reader.Body, i.Structure.ContentSize))
	defer i.Reader.Body.Close()

	if err != nil {
		return err
	}

	i.Structure.Content = part
//	i.Reader.Body       = ioutil.NopCloser(io.MultiReader(bytes.NewReader(part), i.Reader.Body))

	return nil
}

func (i *Input) Clean() {
	i.Structure.File        = EmptyString
	i.Structure.Error       = EmptyString
	i.Structure.Content     = []byte(EmptyString)
	i.Structure.Status      = 0
	i.Structure.ContentSize = 0

	i.Structure = nil
	i.Format    = nil
	i.Reader    = nil

	i.Deploy       = false
	i.Parsing      = false
	i.HttpMimeType = EmptyString
}

func (i *Input) Build() ([]byte, error, int) {
	var err error

	if i.HttpMimeType == HttpMimeTypeInputFile {
		err = i.FileSize()

		if err == nil {
			err = i.FileRead()
			// TODO: 1) parsing define
		}
	} else {
		err = i.BodySize()

		if err == nil {
			err = i.BodyRead()
			// TODO: 1) parsing define
		}
	}

	if err == nil {
		// TODO: parsing
	}

	c := i.Structure.Content
	s := i.Structure.Status

	i.Clean()

	return c, err, s
}

var InputHttpExecute = func(r *http.Request, upload bool, parsing bool) ([]byte, error, int) {
	var input Input

	input.Reader       = &*r
	input.HttpMimeType = InputHttpMimeType(&*r)
	input.Structure    = &format.InputStructure{}

	input.FormatFind()

	return input.Build()
}

func InputHttpBytes(r *http.Request, upload bool, parsing bool) ([]byte, error, int) {
	return InputHttpExecute(&*r, upload, parsing)
}

func InputHttpString(r *http.Request, upload bool, parsing bool) (string, error, int) {
	i, err, s := InputHttpExecute(&*r, upload, parsing)

	return string(i), err, s
}
