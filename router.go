package main

import (
	"net/http"
)

type Router struct {
	Handlers map[string]http.HandlerFunc
}

func NewRouter() *Router {
	return &Router{map[string]http.HandlerFunc{}}
}

func (r *Router) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	for path, handler := range r.Handlers {
		if path == req.URL.Path {
			handler.ServeHTTP(w, req)
			return
		}
	}
}

func (r *Router) AddHandler(path string, handler http.HandlerFunc) {
	r.Handlers[path] = handler
}
