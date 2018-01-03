package hashes

import (
	"net/http"
	"fmt"
	"encoding/base64"

	"rest/handlers"
)

func Base64(w http.ResponseWriter, r *http.Request) {
	if handlers.ErrorMethodPost(w, r) {
		return
	}

	action := handlers.Path2Action(r)

	switch{
	case action == "encode":
		data := []byte("data")
		str := base64.StdEncoding.EncodeToString(data)

		fmt.Fprintf(w, "%s", str)
	case action == "decode":
		str := "ZGF0YQ=="
		data, err := base64.StdEncoding.DecodeString(str)

		if err != nil {
			fmt.Println("error:", err)
			return
		}

		fmt.Fprintf(w, "%s", data)
	}
}
