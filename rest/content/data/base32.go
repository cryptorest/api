package data

import (
	"net/http"
	"encoding/base32"

	"rest/content"
	"rest/errors"
)

const Base32Path = content.DataPath + "/base32"

var Base32Actions = []string{
	"encode",
	"decode",
}

func Base32(w http.ResponseWriter, r *http.Request) {
	if errors.MethodPost(w, r) {
		return
	}

	action := content.Path2Action(r)

	switch action {
	case Base32Actions[0]:
		bData := []byte("data")
		str := base32.StdEncoding.EncodeToString(bData)

		content.OutputString(w, r, str)
	case Base32Actions[1]:
		str := "MRQXIYI="
		bData, err := base32.StdEncoding.DecodeString(str)

		if err != nil {
			content.OutputError(w, r, err)

			return
		}

		content.OutputBytes(w, r, bData)
	}
}
