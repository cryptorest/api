package hashes

import (
	"net/http"
	"encoding/base64"

	"rest/utils"
	"rest/errors"
	"rest/handlers/online"
)

const Base64Path = online.HashesPath + "/base64"

var Base64Actions = [2]string{
	"encode",
	"decode",
}

func Base64(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	action := utils.Path2Action(r)

	switch action {
	case Base64Actions[0]:
		data := []byte("data")
		str := base64.StdEncoding.EncodeToString(data)

		utils.WriteString(w, str)
	case Base64Actions[1]:
		str := "ZGF0YQ=="
		data, err := base64.StdEncoding.DecodeString(str)

		if err != nil {
			utils.WriteError(w, err)

			return
		}

		utils.WriteBytes(w, data)
	}
}
