package main

import (
	"github.com/huandu/facebook"
	"github.com/withitapp/withitd/dbase"
	"github.com/withitapp/withitd/fbook"
	"github.com/withitapp/withitd/model"
	"net/http"
	"time"
)

var (
	db  *dbase.Conn
	fb  *facebook.App
	err error
)

func main() {
	db, err = dbase.NewConn("root:root@tcp(localhost:3306)/withit")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fb = facebook.New("514907081964376", "39cb26381e1cb82ae689ff1d7755f577")
	fb.RedirectUri = "http://withitapp.com/"

	userController := &UserController{db}
	pollController := &PollController{db}

	router := NewRouter()
	router.AddHandler("/users", NewControllerHandler(userController))
	router.AddHandler("/polls", NewControllerHandler(pollController))
	router.AddHandler("/auth", NewAuthHandler())

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

func FetchUser(id string, token string) *model.User {
	user, err := db.UserTable.SelectBy("fb_id", id)
	if err == nil && user != nil {
		println("existing user")
		return user.(*model.User)
	}

	session := fb.Session(token)

	res, err := session.Get("/me", nil)
	err = session.Validate()
	if err != nil {
		panic(err)
	}

	fbUser := fbook.User{}
	err = res.Decode(&fbUser)
	if err != nil {
		panic(err)
	}

	newUser := &model.User{
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Username:   fbUser.Username,
		Email:      fbUser.Email,
		FirstName:  fbUser.FirstName,
		LastName:   fbUser.LastName,
		FbID:       id,
		FbToken:    token,
		FbSyncedAt: time.Now(),
	}

	newUserId, err := db.UserTable.Insert(newUser)
	if err != nil {
		panic(err)
	}

	newUser.ID = newUserId

	return newUser
}
