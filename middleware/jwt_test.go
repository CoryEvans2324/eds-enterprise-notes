package middleware

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

var Set bool

func Router() *mux.Router {
    router := mux.NewRouter()
	router.HandleFunc("/getuser", GetUser).Methods("GET")
    router.HandleFunc("/setuser", SetUser).Methods("POST")
    return router
}

func TestCookie(t *testing.T) {
	assert.Equal(t, "enterprisenotesauth", JWT_TOKEN_COOKIE_NAME, "Token name should be enterprisenotesauth")
}

func TestSigningMethod(t *testing.T) {
	assert.Equal(t, "jwt.SigningMethodHS256", string(jwt.SigningMethodHS256), "Signing Method should be jwt.SigningMethodHS256")
}

func TestGetUser(t *testing.T) {
	request, _ := http.NewRequest("GET", "/getuser", nil)
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
	if Set == false {
		assert.Equal(t, nil, response.Username, "Username should be Nil")
	} else {
		assert.Equal(t, "testuser", response.Username, "Username should be testuser")
		assert.Equal(t, 01, response.UserID, "UserID should be 01")
		assert.Equal(t, "testrole", response.Role, "Role should be testrole")
	}
}

func TestSetUser(t *testing.T) {
	request, _ := http.NewRequest("POST", "/setuser", Testuser)
    response := httptest.NewRecorder()
    Router().ServeHTTP(response, request)
	assert.Equal(t, "testuser", response.Username, "Username should be testuser")
	assert.Equal(t, 01, response.UserID, "UserID should be 01")
	assert.Equal(t, "testrole", response.Role, "Role should be testrole")
	Set := true
}
