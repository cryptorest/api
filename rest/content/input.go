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

const messageErrorContentSize0    = "content size is 0 bytes"
const messageErrorFileUpload      = "file upload error on system"
const messageErrorFileUploadNo    = "file upload does not support"
const messageErrorFileUploadTmpNo = "temporary file upload does not support"

var initTime = fmt.Sprint(time.Now().UnixNano())

func InputHttpMimeType(r *http.Request) string {
	return r.Header.Get(HttpMimeTypeInputKey)
}

func DefaultInputHttpFormat(i *Input) {
	i.Format       = &format.Text
	i.HttpMimeType = format.TextHttpMimeTypes[0]
}

type Input struct {
	Deploy       bool
	Parsing      bool
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

func (i *Input) FormatFindByExtention() {
	fileExt := filepath.Ext(i.Structure.File)

	for _, f := range &Formats {
		for _, fileExtension := range *f.FileExtensions {
			if *&fileExt == *&fileExtension {
				//i.HttpMimeType = fileExtension
				i.Format       = &f

				return
			}
		}
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
	b := []byte(initTime)

	randbo.New().Read(b)

	return fmt.Sprintf(formatHex, b)
}

func (i *Input) FilePut(fileHeader *multipart.FileHeader, isTemporaryFile bool) error {
	var inputFile  multipart.File
	var outputFile *os.File
	var fileName   string
	var dirName    string
	var err        error

	if isTemporaryFile {
		if Config.TemporaryUpload {
			fileName = i.RandomName()
			dirName  = *Config.TmpDir
		} else {
			i.Structure.Status = http.StatusNotAcceptable
			i.Structure.Error  = messageErrorFileUploadTmpNo

			return errors.New(i.Structure.Error)
		}
	} else {
		if Config.FilesUpload {
			fileName = fileHeader.Filename
			dirName  = *Config.UploadDir
		} else {
			i.Structure.Status = http.StatusNotAcceptable
			i.Structure.Error  = messageErrorFileUploadNo

			return errors.New(i.Structure.Error)
		}
	}

	i.Structure.File = filepath.Join(dirName, fileName)
	fileName = EmptyString
	dirName  = EmptyString

	inputFile, err = fileHeader.Open()
	defer inputFile.Close()

	if err != nil {
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

func (i *Input) FileToBufferContentRead(fileHeader *multipart.FileHeader) error {
	var inputFile multipart.File
	var err       error

	i.Structure.Content = make([]byte, i.Structure.ContentSize)

	inputFile, err = fileHeader.Open()
	defer inputFile.Close()

	if err != nil {
		i.Structure.Status = http.StatusNotAcceptable
		i.Structure.Error  = err.Error()

		return err
	}

	_, err = inputFile.Read(i.Structure.Content)

	return err
}

func (i *Input) FileRead(isTemporaryFile bool) error {
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
			i.Structure.File = fileHeader.Filename

			if int(i.Structure.ContentSize) > Config.BufferSize {
				err = i.FilePut(&*fileHeader, isTemporaryFile)

				if err == nil {
					err = i.FileContentRead()
				}
			} else {
				err = i.FileToBufferContentRead(&*fileHeader)
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
	var err error

	i.Structure.Content, err = ioutil.ReadAll(io.LimitReader(i.Reader.Body, i.Structure.ContentSize))
	defer i.Reader.Body.Close()

	return err
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

func (i *Input) Build(isTemporaryFile bool, isFormatParse bool) ([]byte, error, int) {
	var err error

	if i.HttpMimeType == HttpMimeTypeInputFile {
		err = i.FileSize()

		if err == nil {
			err = i.FileRead(isTemporaryFile)

			if err == nil && isFormatParse {
				i.FormatFindByExtention()
			}
		}
	} else {
		err = i.BodySize()

		if err == nil {
			err = i.BodyRead()
		}
	}

	if err == nil && isFormatParse {
		i.Format.InputFormatFunc(*&i.Structure.Content, nil)
	}

	defer i.Clean()

	return i.Structure.Content, err, i.Structure.Status
}

var InputHttpExecute = func(r *http.Request, isTemporaryFile bool, isFormatParse bool) ([]byte, error, int) {
	var input Input

	input.Reader       = &*r
	input.HttpMimeType = InputHttpMimeType(&*r)
	input.Structure    = &format.InputStructure{}

	input.FormatFind()

	return input.Build(isTemporaryFile, isFormatParse)
}

func InputHttpBytes(r *http.Request, isTemporaryFile bool, isFormatParse bool) ([]byte, error, int) {
	return InputHttpExecute(&*r, isTemporaryFile, isFormatParse)
}

func InputHttpString(r *http.Request, isTemporaryFile bool, isFormatParse bool) (string, error, int) {
	i, err, s := InputHttpExecute(&*r, isTemporaryFile, isFormatParse)

	return string(i), err, s
}
