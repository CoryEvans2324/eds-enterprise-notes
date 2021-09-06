package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestJWTMiddleware(t *testing.T) {
	router := mux.NewRouter()

	// The middleware we are testing
	router.Use(JWTMiddleware)

	// create a simple route that doesn't do much
	router.HandleFunc("/testing", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK"))
	})

	request, _ := http.NewRequest("GET", "/testing", nil)
	response := httptest.NewRecorder()

	// Simple get request with nothing set
	router.ServeHTTP(response, request)

	// TODO: Test all other functions/code

}
