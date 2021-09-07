package main

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gorilla/mux"
    "github.com/stretchr/testify/assert"
)

Testuser := models.JWTUser {
	UserID: 01,
	Username: "testname",
	Role: "testrole",
}

func Router() *mux.Router {
    router := mux.NewRouter()
	router.HandleFunc("/getuser", GetUser).Methods("GET")
    router.HandleFunc("/setuser", SetUser).Methods("POST")
	//todo - JWTMiddleware function tests, add more conditions to bellow tests, subdivide to ensure they can be ran in isolation to eachother?
    return router
}

func TestGetUserNil(t *testing.T) {
	request, _ := http.NewRequest("GET", "/getuser", nil)
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
	assert.Equal(t, nil, response.Username, "Username should be Nil")
}

func TestSetUser(t *testing.T) {
	request, _ := http.NewRequest("POST", "/setuser", Testuser)
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
	assert.Equal(t, "testuser", response.Username, "Username should be testuser")
}
