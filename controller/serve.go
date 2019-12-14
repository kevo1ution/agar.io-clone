package controller

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func SetupControllers() {
	dir := "../client"

	flag.StringVar(&dir, "dir", ".", "the directory to serve files from. Defaults to the current dir")
	flag.Parse()
	r := mux.NewRouter()

	// This will serve files under http://localhost:8000/static/<filename>
	r.PathPrefix("/client").Handler(http.FileServer(http.Dir(dir)))

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8000",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())
}
