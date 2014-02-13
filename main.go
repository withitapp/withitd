package main

import (
	"encoding/json"
	"github.com/withitapp/withitd/dbase"
	"net/http"
	"net/url"
	"strconv"
	"time"
)

func main() {
	userTable := &dbase.UserTable{}
	userController := &UserController{userTable}
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

type Controller interface {
	Index() interface{}
	Show(ID int) (interface{}, error)
	Create(values url.Values) (int, error)
	Update(ID int, values url.Values) error
	Delete(ID int) error
}

func NewControllerHandler(controller Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idString := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idString)
		if err != nil {
			panic(err)
		}

		var response interface{}

		switch r.Method {
		case "GET":
			if id != 0 {
				response, err = controller.Show(id)
			} else {
				response = controller.Index()
			}
		case "PUT", "POST":
			id, err = controller.Create(r.PostForm)
		case "UPDATE":
			err = controller.Update(id, r.PostForm)
		case "DELETE":
			err = controller.Delete(id)
		default:
		}

		if err != nil {
			panic(err)
		}

		jsonBlob, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBlob)
	}
}
