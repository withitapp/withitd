package main

import (
	"github.com/huandu/facebook"
	"github.com/withitapp/withitd/dbase"
	"net/http"
	"time"
)

func main() {
	db, err := dbase.NewConn("root:root@tcp(localhost:3306)/withit")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fb := facebook.New("514907081964376", "39cb26381e1cb82ae689ff1d7755f577")
	fb.RedirectUri = "http://withitapp.com/"

	userController := &UserController{db}
	pollController := &PollController{db}

	router := NewRouter()
	router.AddHandler("/users", NewControllerHandler(userController))
	router.AddHandler("/polls", NewControllerHandler(pollController))
	router.AddHandler("/auth", NewAuthHandler(db, fb))

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
