package main

import (
	"fmt"
	"net/http"
)

func getFriendshipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, map[string]string{
		"success": "false",
		"message": "Not Implemented",
	})
	return
}

func postFriendshipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, map[string]string{
		"success": "false",
		"message": "Not Implemented",
	})
	return
}
