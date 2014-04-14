package main

import (
	"encoding/json"
	"fmt"
	"github.com/withitapp/withitd/dbase"
	"github.com/withitapp/withitd/model"
	"net/http"
)

// Using poll_id paramater, returns members of poll
func NewMembersHandler(db *dbase.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()

		pollID := r.Form.Get("poll_id")

		query := "select users.* "
		query += "from users, memberships "
		query += "where " + pollID + "=memberships.poll_id AND memberships.user_id = users.id"

		members, err := db.UserTable.Query(query)
		if err != nil {
			panic(err)
		}

		jsonBlob, err := json.Marshal(members)
		if err != nil {
			panic(err)
		}

		fmt.Printf("Returning %v members from poll %v\n", len(members.([]*model.User)), pollID)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBlob)
	}
}
