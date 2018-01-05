package data

import (
	"net/http"
	"encoding/base64"

	"rest/data"
	"rest/errors"
)

const Base64Path = data.DataPath + "/base64"

var Base64Actions = []string{
	"encode",
	"decode",
}

func Base64(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	action := data.Path2Action(r)

	switch action {
	case Base64Actions[0]:
		bData := []byte("data")
		str := base64.StdEncoding.EncodeToString(bData)

		data.WriteString(w, str)
	case Base64Actions[1]:
		str := "ZGF0YQ=="
		bData, err := base64.StdEncoding.DecodeString(str)

		if err != nil {
			data.WriteError(w, err)

			return
		}

		data.WriteBytes(w, bData)
	}
}
