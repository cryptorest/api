package main

import (
	"fmt"
	"sort"
	"net/http"

	"rest/errors"
	"rest/config"
	"rest/handlers"
)

const RootPath = "/"

func showRequestInfo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")

	fmt.Fprintf(w, "Method: %s\n", r.Method)
	fmt.Fprintf(w, "Protocol: %s\n", r.Proto)
	fmt.Fprintf(w, "Host: %s\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr: %s\n", r.RemoteAddr)
	fmt.Fprintf(w, "RequestURI: %q\n", r.RequestURI)
	fmt.Fprintf(w, "URL: %#v\n", r.URL)
	fmt.Fprintf(w, "Body.ContentLength: %d (-1 means unknown)\n", r.ContentLength)
	fmt.Fprintf(w, "Close: %v (relevant for HTTP/1 only)\n", r.Close)
	fmt.Fprintf(w, "TLS: %#v\n", r.TLS)
	fmt.Fprint(w, "Pathes:\n")
	sort.Strings(handlers.AllPathes)
	for _, p := range handlers.AllPathes {
		fmt.Fprintf(w, "    %s\n", p)
	}

//	r.Header.Write(w)
}

func Root(w http.ResponseWriter, r *http.Request) {
	if errors.Path(w, r, RootPath) {
		return
	}

	if errors.IsMethodGet(r) && errors.IsMethodHead(r) {
		errors.MethodGet(w, r)

		return
	}

	showRequestInfo(w, r)
}

func initHandlers() {
	mux := http.NewServeMux()

	http.HandleFunc(RootPath, func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.TLS == nil:
			http.Redirect(w, r, config.ServerURI(RootPath), http.StatusFound)

			return
		}

		mux.ServeHTTP(w, r)
	})

	mux.HandleFunc(RootPath, Root)
	handlers.InitOnline(&*mux)
}
