package main

import (
	"github.com/huandu/facebook"
	"github.com/withitapp/withitd/dbase"
	"net/http"
	"time"
)

const (
	dbURL = "root:root@tcp(localhost:3306)/withit"
	fbKey = "514907081964376"
	fbSec = "39cb26381e1cb82ae689ff1d7755f577"
)

func main() {
	db, err := dbase.NewConn(dbURL)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fb := facebook.New(fbKey, fbSec)
	fb.RedirectUri = "http://withitapp.com/"

	router := NewRouter()
	router.AddHandler("/users", NewControllerHandler(&UserController{db}))
	router.AddHandler("/polls", NewControllerHandler(&PollController{db}))
	router.AddHandler("/memberships", NewControllerHandler(&MembershipController{db}))
	router.AddHandler("/auth", NewAuthHandler(db, fb))
	router.AddHandler("/members", NewMembersHandler(db))

	server := &http.Server{
		Addr:           ":3000",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	println("Starting http server at port 3000 ...")
	server.ListenAndServe()
}
