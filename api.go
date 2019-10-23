package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/reconditematter/svc"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	const ADDR = ":5001"
	//
	R := mux.NewRouter()
	R.Handle("/api/reconditematter", handlers.LoggingHandler(os.Stderr, http.HandlerFunc(usage)))
	//
	svc.RandomNames(R)
	//
	srv := &http.Server{
		Handler:      R,
		Addr:         ADDR,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	//
	log.Fatal(srv.ListenAndServe())
}

func usage(w http.ResponseWriter, r *http.Request) {
	doc := `
/reconditematter/randomnames -- API to generate random names of both genders.
`
	svc.HS200t(w, []byte(doc))
}
