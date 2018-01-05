package hashes

import (
	"net/http"
	"encoding/base32"

	"rest/utils"
	"rest/errors"
	"rest/handlers/online"
)

const Base32Path = online.HashesPath + "/base32"

var Base32Actions = [2]string{
	"encode",
	"decode",
}

func Base32(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	action := utils.Path2Action(r)

	switch action {
	case Base32Actions[0]:
		data := []byte("data")
		str := base32.StdEncoding.EncodeToString(data)

		utils.WriteString(w, str)
	case Base32Actions[1]:
		str := "MRQXIYI="
		data, err := base32.StdEncoding.DecodeString(str)

		if err != nil {
			utils.WriteError(w, err)

			return
		}

		utils.WriteBytes(w, data)
	}
}
