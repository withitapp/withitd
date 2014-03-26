package main

import (
	"fmt"
	"github.com/huandu/facebook"
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
	appHandler.AddHandler("/auth", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		data := r.PostForm
		id := data["fb_id"][0]
		token := data["fb_token"][0]
		FetchUser(id, token)
	})

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

func FetchUser(id string, token string) {
	fbApp := facebook.New("514907081964376", "39cb26381e1cb82ae689ff1d7755f577")
	fbApp.RedirectUri = "http://withitapp.com/"

	session := fbApp.Session(token)

	res, err := session.Get("/me/feed", nil)

	var name string
	res.DecodeField("first_name", &name)
	if err != nil {
		panic(err)
	}

	err = session.Validate()
	if err != nil {
		panic(err)
	}

	fmt.Println(name)
}
