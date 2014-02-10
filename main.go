package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/users/", getUserHandler)
	fmt.Println("Starting http server at port 3000 ...")
	http.ListenAndServe(":3000", nil)
}
