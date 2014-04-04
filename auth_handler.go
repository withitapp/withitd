package main

import (
	"encoding/json"
	"fmt"
	"github.com/huandu/facebook"
	"github.com/withitapp/withitd/dbase"
	"github.com/withitapp/withitd/fbook"
	"github.com/withitapp/withitd/model"
	"net/http"
	"time"
)

func NewAuthHandler(db *dbase.Conn, fb *facebook.App) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		data := r.PostForm

		id := data["fb_id"][0]
		token := data["fb_token"][0]

		user := FetchUser(db, fb, id, token)
		jsonBlob, err := json.Marshal(user)
		if err != nil {
			panic(err)
		}

		fmt.Printf("User logged in: id: %s, %s", user.ID, user.FirstName)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBlob)
	}
}

func FetchUser(db *dbase.Conn, fb *facebook.App, id, token string) *model.User {
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
