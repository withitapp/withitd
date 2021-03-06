package main

import (
	"encoding/json"
	"github.com/withitapp/withitd/dbase"
	"net/http"
)

// Using poll_id paramater, returns members of poll
func NewMembersHandler(db *dbase.Conn) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		pollID := r.FormValue("poll_id")

		members, err := db.SelectAllMembersOfPoll(pollID)
		if err != nil {
			panic(err)
		}

		jsonBlob, err := json.Marshal(members)
		if err != nil {
			panic(err)
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBlob)
	}
}
