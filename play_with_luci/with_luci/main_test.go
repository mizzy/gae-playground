package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/luci/gae/impl/memory"
	"github.com/luci/gae/service/datastore"
	"golang.org/x/net/context"
)

func Test_HandlePostUser(t *testing.T) {
	ctx := memory.Use(context.Background())

	response := httptest.NewRecorder()
	json := `{"name": "mahiru", "email": "inami@example.com"}`
	request, _ := http.NewRequest("POST", "/user", strings.NewReader(json))

	handlePostUser(ctx, response, request)

	if response.Code != http.StatusOK {
		t.Fatalf("unexpected code: %d", response.Code)
	}

	user := UserEntity{Key: datastore.NewKey(ctx, userKind, "inami@example.com", 0, nil)}
	err := datastore.Get(ctx, &user)
	if err != nil {
		t.Fatal(err)
	}
	if user.Email != "inami@example.com" {
		t.Fatalf("invalid email: %s", user.Email)
	}
	if user.Name != "mahiru" {
		t.Fatalf("invalid name: %s", user.Name)
	}
}
