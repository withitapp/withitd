package main

import (
	"encoding/json"
	"fmt"
	"github.com/withitapp/withitd/dbase"
	"net/http"
)

// Using poll_id paramater, returns members of poll
func NewMembersHandler(db *dbase.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		data := r.PostForm

		//id := data["id"][0]
		//userID := data["user_id"][0]
		pollID := data["poll_id"][0]

		//TODO return memberships
		memberships, err := db.MembershipTable.SelectAllBy("pollID", pollID)
		if err != nil {
			panic(err)
		}

		jsonBlob, err := json.Marshal(memberships)

		if err != nil {
			panic(err)
		}

		fmt.Printf("Returning memberships from poll %v\n", pollID)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBlob)

	}
}
