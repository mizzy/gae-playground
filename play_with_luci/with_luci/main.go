package main

import (
	"encoding/json"
	"net/http"

	"github.com/luci/gae/impl/prod"
	"github.com/luci/gae/service/datastore"
	"golang.org/x/net/context"
	"google.golang.org/appengine"
)

const userKind = "User"

// User ...
type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

// UserEntity ...
type UserEntity struct {
	User
	Key *datastore.Key `gae:"$key"`
}

func init() {
	http.HandleFunc("/user", userHandler)
}

func userHandler(w http.ResponseWriter, r *http.Request) {
	ctx := prod.Use(appengine.NewContext(r), r)
	switch r.Method {
	case "POST":
		handlePostUser(ctx, w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handlePostUser(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	key := datastore.NewKey(ctx, userKind, user.Email, 0, nil)
	err = datastore.Put(ctx, &UserEntity{user, key})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
