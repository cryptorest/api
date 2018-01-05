package hashes

import (
	"net/http"
	"encoding/base32"

	"rest/data"
	"rest/errors"
	"rest/handlers/online"
)

const Base32Path = online.HashesPath + "/base32"

var Base32Actions = []string{
	"encode",
	"decode",
}

func Base32(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	action := data.Path2Action(r)

	switch action {
	case Base32Actions[0]:
		bData := []byte("data")
		str := base32.StdEncoding.EncodeToString(bData)

		data.WriteString(w, str)
	case Base32Actions[1]:
		str := "MRQXIYI="
		bData, err := base32.StdEncoding.DecodeString(str)

		if err != nil {
			data.WriteError(w, err)

			return
		}

		data.WriteBytes(w, bData)
	}
}
