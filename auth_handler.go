package main

import (
	"encoding/json"
	"fmt"
	"github.com/withitapp/facebook"
	"github.com/withitapp/withitd/dbase"
	"github.com/withitapp/withitd/fbook"
	"github.com/withitapp/withitd/model"
	"net/http"
	"strconv"
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

		fmt.Printf("User logged in: %v - %v\n", id, user.FirstName)

		FetchFriends(db, fb, user)

		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonBlob)
	}
}

func FetchFriends(db *dbase.Conn, fb *facebook.App, user *model.User) {
	session := fb.Session(user.FbToken)

	res, err := session.Get("/me/friends?fields=installed", nil)
	err = session.Validate()
	if err != nil {
		panic(err)
	}

	var response struct {
		Data []struct {
			Id        string
			Installed bool
		}
	}

	err = res.Decode(&response)
	if err != nil {
		panic(err)
	}

	friendships, err := db.FriendshipTable.SelectAllBy("alpha_id", strconv.Itoa(user.ID))
	if err != nil {
		panic(err)
	}

	betaExists := func(oldFriendships []*model.Friendship, betaID int) bool {
		for _, oldFriendship := range friendships.([]*model.Friendship) {
			if oldFriendship.BetaID == betaID {
				return true
			}
		}

		return false
	}

	for _, fbUser := range response.Data {
		if !fbUser.Installed {
			continue
		}

		friendUser, err := db.UserTable.SelectBy("fb_id", fbUser.Id)
		if err != nil && user == nil {
			continue
		}

		if betaExists(friendships.([]*model.Friendship), friendUser.(*model.User).ID) {
			continue
		}

		friendship := &model.Friendship{
			CreatedAt: time.Now(),
			AlphaID:   user.ID,
			BetaID:    friendUser.(*model.User).ID,
		}

		_, err = db.FriendshipTable.Insert(friendship)
		if err != nil {
			panic(err)
		}
	}
}

func FetchUser(db *dbase.Conn, fb *facebook.App, id, token string) *model.User {
	user, err := db.UserTable.SelectBy("fb_id", id)
	if err == nil && user != nil {
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

	fmt.Printf("New user created: %v - %v\n", id, newUser.FirstName)
	newUser.ID = newUserId

	return newUser
}
