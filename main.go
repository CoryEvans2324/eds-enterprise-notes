package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/CoryEvans2324/eds-enterprise-notes/config"
	"github.com/CoryEvans2324/eds-enterprise-notes/database"
	"github.com/CoryEvans2324/eds-enterprise-notes/middleware"
	"github.com/CoryEvans2324/eds-enterprise-notes/routes"
	"github.com/gorilla/mux"
)

func init() {
	cfgData, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	config.LoadConfig(cfgData)
	database.CreateDatabaseManager(config.Get().Database.DataSourceName())
}

func main() {
	cfg := config.Get()
	r := mux.NewRouter().StrictSlash(true)

	r.Use(middleware.JWTMiddleware)

	r.HandleFunc("/", routes.Index)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(cfg.Server.StaticFolder))))

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/signin", routes.UserSignIn)
	userRouter.HandleFunc("/create", routes.UserSignUp)

	srv := &http.Server{
		Handler:      r,
		Addr:         cfg.Server.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
