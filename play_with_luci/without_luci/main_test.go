package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"google.golang.org/appengine/aetest"
	"google.golang.org/appengine/datastore"
)

func Test_HandlePostUser(t *testing.T) {
	ctx, done, err := aetest.NewContext()
	if err != nil {
		t.Fatal(err)
	}
	defer done()

	response := httptest.NewRecorder()
	json := `{"name": "mahiru", "email": "inami@example.com"}`
	request, _ := http.NewRequest("POST", "/user", strings.NewReader(json))

	handlePostUser(ctx, response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("unexpected code: %d", response.Code)
	}

	user := User{}
	datastore.Get(ctx, datastore.NewKey(ctx, userKind, "inami@example.com", 0, nil), &user)
	if user.Email != "inami@example.com" {
		t.Fatalf("invalid email: %s", user.Email)
	}
	if user.Name != "mahiru" {
		t.Fatalf("invalid name: %s", user.Name)
	}
}
