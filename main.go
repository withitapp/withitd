package main

import (
	"github.com/withitapp/withitd/cntrl"
	"github.com/withitapp/withitd/dbase"
	"net/http"
	"time"
)

func main() {
	dbaseConn, err := dbase.NewConn("root:root@tcp(localhost:3306)/withit")
	if err != nil {
		panic(err)
	}
	defer dbaseConn.Close()

	userController := &cntrl.UserController{dbaseConn}
	pollController := &cntrl.PollController{dbaseConn}

	appHandler := NewAPPHandler()
	appHandler.AddHandler("/users", NewControllerHandler(userController))
	appHandler.AddHandler("/polls", NewControllerHandler(pollController))

	server := &http.Server{
		Addr:           ":3000",
		Handler:        appHandler,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	println("Starting http server at port 3000 ...")
	server.ListenAndServe()
}
