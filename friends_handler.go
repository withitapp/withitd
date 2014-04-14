package main

import (
	"encoding/json"
	"github.com/withitapp/withitd/dbase"
	"net/http"
)

func NewFriendsHandler(db *dbase.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		userID := r.Form.Get("user_id")

		query := "select users.* "
		query += "from users, friendships "
		query += "where " + userID + "=friendships.alpha_id AND friendships.beta_id = users.id"

		friends, err := db.UserTable.Query(query)
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
