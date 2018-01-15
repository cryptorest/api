package data

import (
	"fmt"
	e "errors"
	"net/http"
	"encoding/base32"

	"rest/errors"
	"rest/content"
)

const Base32Path = content.DataPath + "/base32"

const errorBase32Message = "invalid action %s for Base32"

var Base32Actions = []string{
	"encode",
	"decode",
}

func Base32Encode(data []byte) (string, error) {
	return base32.StdEncoding.EncodeToString(data), nil
}

func Base32Decode(str string) ([]byte, error) {
	return base32.StdEncoding.DecodeString(str)
}

func Base32Http(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, &*r) {
		return
	}

	var err error
	action := content.Path2Action(&*r)

	switch action {
	// Encode
	case Base32Actions[0]:
		data, err, s := content.InputHttpBytes(&*r, false, false)

		if err == nil {
			str, _ := Base32Encode(data)

			content.OutputHttpString(w, r, str)
		} else {
			content.OutputHttpError(w, &*r, err, s)
		}
	// Decode
	case Base32Actions[1]:
		str, err, s := content.InputHttpString(&*r, false, false)

		if err == nil {
			data, err := Base32Decode(str)

			if err == nil {
				content.OutputHttpByte(w, &*r, data)
			} else {
				content.OutputHttpError(w, &*r, err, http.StatusUnprocessableEntity)
			}
		} else {
			content.OutputHttpError(w, &*r, err, s)
		}
	default:
		err = e.New(fmt.Sprintf(errorBase32Message, action))
	}

	if err != nil {
		content.OutputHttpError(w, &*r, err, http.StatusNotAcceptable)
	}
}
