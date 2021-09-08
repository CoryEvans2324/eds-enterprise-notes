package main

import (
	"context"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
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
	database.Mgr.AutoMigrate()
}

func main() {
	cfg := config.Get()
	r := mux.NewRouter().StrictSlash(true)

	r.Use(middleware.JWTMiddleware)

	r.HandleFunc("/", routes.Index)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(cfg.Server.StaticFolder))))

	userRouter := r.PathPrefix("/user").Subrouter()
	userRouter.HandleFunc("/@{username}", routes.UserView)
	userRouter.HandleFunc("/signin", routes.UserSignIn)
	userRouter.HandleFunc("/signout", routes.UserSignOut)
	userRouter.HandleFunc("/create", routes.UserSignUp)
	userRouter.HandleFunc("/search", routes.UserSearch)

	noteRouter := r.PathPrefix("/note").Subrouter()
	noteRouter.HandleFunc("/create", routes.CreateNote)
	noteRouter.HandleFunc("/{id:[0-9]+}", routes.GetNote)
	noteRouter.HandleFunc("/{id:[0-9]+}/edit", routes.EditNote)

	debugRouter := r.PathPrefix("/debug").Subrouter()
	debugRouter.HandleFunc("/reset", routes.DebugResetDB)
	debugRouter.HandleFunc("/createusers", routes.DebugCreateDummyUsers)

	srv := &http.Server{
		Handler:      r,
		Addr:         cfg.Server.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Graceful shutdown from https://github.com/gorilla/mux#graceful-shutdown

	// Run our server in a goroutine so that it doesn't block.
	go func() {
		if err := srv.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	log.Println("shutting down")
	srv.Shutdown(ctx)

	os.Exit(0)
}
