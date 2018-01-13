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

const messageErrorContentSize0 = "content size is 0 bytes"
const messageErrorFileUpload   = "upload file error on system"

func InputHttpMimeType(r *http.Request) string {
	return r.Header.Get(HttpMimeTypeInputKey)
}

func DefaultInputHttpFormat(i *Input) {
	i.Format       = &format.Text
	i.HttpMimeType = format.TextHttpMimeTypes[0]
}

type Input struct {
	BufferSize    int
	FileSizeLimit int64
	BodySizeLimit int64
	HttpMimeType  string
	UploadDir     *string
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

		if i.Structure.ContentSize > i.FileSizeLimit {
			err                = http.ErrContentLength
			i.Structure.Status = http.StatusLengthRequired
		}
	}

	return err
}

func (i *Input) BufferRead(r multipart.File, w io.Writer) error {
	var n   int
	var err error

	buf := make([]byte, i.BufferSize)

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

func (i *Input) FilePut(fileHeader *multipart.FileHeader, err error) error {
	var inputFile  multipart.File
	var outputFile *os.File

	i.Structure.File = filepath.Join(*i.UploadDir, fileHeader.Filename)

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

func (i *Input) FileRead() error {
	var err error

	defer func() error {
		return err
	}()

	err = i.Reader.ParseMultipartForm(i.FileSizeLimit)
	if err != nil {
		i.Structure.Status = http.StatusNotAcceptable
		i.Structure.Error  = err.Error()

		return err
	}

	for _, fileHeaders := range i.Reader.MultipartForm.File {
		for _, fileHeader := range fileHeaders {
			err = i.FilePut(&*fileHeader, err)

			if err != nil {
				return err
			}
		}
	}

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

		if i.Structure.ContentSize > i.BodySizeLimit {
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
	i.Structure = nil
	i.Format    = nil
	i.Reader    = nil

	i.BufferSize    = 0
	i.FileSizeLimit = 0
	i.HttpMimeType  = EmptyString
}

func (i *Input) Build() ([]byte, error, int) {
	var err error

	if i.HttpMimeType == HttpMimeTypeInputFile {
		err = i.FileSize()

		if err == nil {
			err = i.FileRead()
			// TODO: 1) parsing or 2) upload or 3) upload and remove
		}
	} else {
		err = i.BodySize()

		if err == nil {
			err = i.BodyRead()
			// TODO: 1) parsing or 2) not parsing
		}
	}

	c := i.Structure.Content
	s := i.Structure.Status

	i.Clean()

	return c, err, s
}

var InputHttpExecute = func(r *http.Request) ([]byte, error, int) {
	var input Input

	input.BufferSize    = config.Server.BufferSize * config.BufferSizeBlock
	input.FileSizeLimit = int64(config.Server.FileSizeLimit * config.BufferSizeBlock)
	input.BodySizeLimit = int64(config.Server.BodySizeLimit * config.BufferSizeBlock)
	input.UploadDir     = &config.Server.UploadDir
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
