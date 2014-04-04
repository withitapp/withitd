package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func NewAuthHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		data := r.PostForm

		id := data["fb_id"][0]
		token := data["fb_token"][0]

		user := FetchUser(id, token)
		jsonBlob, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}

		fmt.Printf("User logged in: id: %s, %s", user.ID, user.FirstName)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBlob)
	}
}
