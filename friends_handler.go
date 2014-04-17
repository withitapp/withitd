package main

import (
	"encoding/json"
	"github.com/withitapp/withitd/dbase"
	"net/http"
)

func NewFriendsHandler(db *dbase.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := r.FormValue("user_id")

		friends, err := db.SelectAllFriendsOfUser(userID)
		if err != nil {
			panic(err)
		}

		jsonBlob, err := json.Marshal(friends)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBlob)
	}
}
