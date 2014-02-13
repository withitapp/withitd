package main

import (
	"net/http"
)

type APPHandler struct {
	Handlers map[string]http.Handler
}

func (ah *APPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for path, handler := range ah.Handlers {
		if path == r.URL.Path {
			handler.ServeHTTP(w, r)
		}
	}
}
