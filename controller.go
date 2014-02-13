package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

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
