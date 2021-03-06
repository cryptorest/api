package data

import (
	"fmt"
	e "errors"
	"net/http"
	"encoding/base64"

	"rest/errors"
	"rest/content"

)

const Base64Path = content.DataPath + "/base64"

const errorBase64Message = "invalid action %s for Base64"

var Base64Actions = []string{
	"encode",
	"decode",
}

func Base64Encode(data []byte) (string, error) {
	return base64.StdEncoding.EncodeToString(data), nil
}

func Base64Decode(str string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(str)
}

func Base64Http(w http.ResponseWriter, r *http.Request) {
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
			str, _ := Base64Encode(data)

			content.OutputHttpString(w, &*r, str)
		} else {
			content.OutputHttpError(w, &*r, err, s)
		}
	// Decode
	case Base32Actions[1]:
		str, err, s := content.InputHttpString(&*r, false, false)

		if err == nil {
			data, err := Base64Decode(str)

			if err == nil {
				content.OutputHttpByte(w, &*r, data)
			} else {
				content.OutputHttpError(w, &*r, err, http.StatusUnprocessableEntity)
			}
		} else {
			content.OutputHttpError(w, &*r, err, s)
		}
	default:
		err = e.New(fmt.Sprintf(errorBase64Message, action))
	}

	if err != nil {
		content.OutputHttpError(w, &*r, err, http.StatusNotAcceptable)
	}
}
