package main

import (
	"fmt"
	"net/http"
)

func getMembershipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, map[string]string{
		"success": "false",
		"message": "Not Implemented",
	})
	return
}

func getMembershipsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, map[string]string{
		"success": "false",
		"message": "Not Implemented",
	})
	return
}

func postMembershipHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprint(w, map[string]string{
		"success": "false",
		"message": "Not Implemented",
	})
	return
}
