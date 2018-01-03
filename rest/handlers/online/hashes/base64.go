package hashes

import (
	"net/http"
	"encoding/base64"

	"rest/handlers"
	"rest/handlers/online"
)

const Base64Path string = online.HashesPath + "/base64"

var Base64_Actions = [2]string{
	"encode",
	"decode",
}

func Base64(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	action := handlers.Path2Action(r)

	switch{
	case action == Base64_Actions[0]:
		data := []byte("data")
		str := base64.StdEncoding.EncodeToString(data)

		handlers.WriteString(w, []byte(str))
	case action == Base64_Actions[1]:
		str := "ZGF0YQ=="
		data, err := base64.StdEncoding.DecodeString(str)

		if err != nil {
			handlers.WriteError(w, err)

			return
		}

		handlers.WriteString(w, data)
	}
}