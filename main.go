package main

import (
	"github.com/withitapp/withitd/cntrl"
	"github.com/withitapp/withitd/dbase"
	"net/http"
	"time"
)

func main() {
	userTable := &dbase.UserTable{}
	userController := &cntrl.UserController{userTable}
	userHandler := NewControllerHandler(userController)

	appHandler := &APPHandler{
		map[string]http.Handler{
			"/users": userHandler,
		},
	}

	server := &http.Server{
		Addr:           ":3000",
		Handler:        appHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	println("Starting http server at port 3000 ...")
	server.ListenAndServe()
}

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
