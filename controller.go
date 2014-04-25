package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strconv"
)

type Controller interface {
	Index(values url.Values) (interface{}, error)
	Show(ID int) (interface{}, error)
	Create(values url.Values) (int, error)
	Update(ID int, values url.Values) error
	Delete(ID int) error
}

func NewControllerHandler(controller Controller) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			response interface{}
			err      error
			id       int
		)

		idString := r.FormValue("id")
		if idString != "" {
			id, err = strconv.Atoi(idString)
		}

		switch r.Method {
		case "GET":
			if id != 0 {
				response, err = controller.Show(id)
			} else {
				response, err = controller.Index(r.Form)
			}
		case "PUT", "POST", "PATCH":
			if idString == "" {
				id, err = controller.Create(r.Form)
				response = map[string]int{"id": id}
			} else {
				err = controller.Update(id, r.Form)
			}
		case "DELETE":
			err = controller.Delete(id)
		default:
		}

		if err != nil {
			response = err.Error()
		}

		jsonBlob, err := json.Marshal(response)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBlob)
	}
}
