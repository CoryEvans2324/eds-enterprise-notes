package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/CoryEvans2324/eds-enterprise-notes/config"
	"github.com/CoryEvans2324/eds-enterprise-notes/routes"
	"github.com/gorilla/mux"
)

func init() {
	cfgData, err := ioutil.ReadFile("config.yml")
	if err != nil {
		log.Fatalln(err.Error())
	}
	config.LoadConfig(cfgData)
}

func main() {
	cfg := config.Get()
	r := mux.NewRouter()
	r.HandleFunc("/", routes.Index)

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir(cfg.Server.StaticFolder))))

	srv := &http.Server{
		Handler:      r,
		Addr:         cfg.Server.Address,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
