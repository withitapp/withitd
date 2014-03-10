package main

import (
	"net/http"
)

type APPHandler struct {
	Handlers map[string]http.HandlerFunc
}

func NewAPPHandler() *APPHandler {
	return &APPHandler{map[string]http.HandlerFunc{}}
}

func (ah *APPHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for path, handler := range ah.Handlers {
		if path == r.URL.Path {
			handler.ServeHTTP(w, r)
		}
	}
}

func (ah *APPHandler) AddHandler(path string, handler http.HandlerFunc) {
	ah.Handlers[path] = handler
}
