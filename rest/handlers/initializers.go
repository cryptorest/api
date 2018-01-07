package handlers

import (
	"path"
	"hash"
	"net/http"
)

var AllPathes []string

func initHandlerAsString(mux *http.ServeMux, httpPath string, httpFunc func(w http.ResponseWriter, r *http.Request)) {
	mux.HandleFunc(httpPath, httpFunc)
	AllPathes = append(AllPathes, httpPath)
}

func initHandlerAsArray(mux *http.ServeMux, httpPath string, httpFunc func(w http.ResponseWriter, r *http.Request), values []string) {
	for _, value := range values {
		p := path.Join(httpPath, value)

		mux.HandleFunc(p, httpFunc)
		AllPathes = append(AllPathes, p)
	}
}

func initHandlerAsHash(mux *http.ServeMux, httpPath string, httpFunc func(w http.ResponseWriter, r *http.Request), values map[string]func() hash.Hash) {
	for value := range values {
		p := path.Join(httpPath, value)

		mux.HandleFunc(p, httpFunc)
		AllPathes = append(AllPathes, p)
	}
}
