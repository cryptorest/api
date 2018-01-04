package hashes

import (
	"net/http"
	"encoding/base32"

	"rest/handlers"
	"rest/handlers/online"
)

const Base32Path string = online.HashesPath + "/base32"

var Base32Actions = [2]string{
	"encode",
	"decode",
}

func Base32(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	action := handlers.Path2Action(r)

	switch action {
	case Base32Actions[0]:
		data := []byte("data")
		str := base32.StdEncoding.EncodeToString(data)

		handlers.WriteString(w, str)
	case Base32Actions[1]:
		str := "MRQXIYI="
		data, err := base32.StdEncoding.DecodeString(str)

		if err != nil {
			handlers.WriteError(w, err)

			return
		}

		handlers.WriteBytes(w, data)
	}
}
