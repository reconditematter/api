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
	port := os.Getenv("PORT")
	if port == "" {
		port = "5001"
	}
	//
	R := mux.NewRouter()
	R.Handle("/api/reconditematter", handlers.LoggingHandler(os.Stderr, http.HandlerFunc(usage)))
	//
	svc.RandomNames(R)
	svc.GeoMatrix(R)
	svc.GeoHash(R)
	svc.KFunction(R)
	//
	srv := &http.Server{
		Handler:      R,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	//
	log.Fatal(srv.ListenAndServe())
}

func usage(w http.ResponseWriter, r *http.Request) {
	doc := `
/reconditematter/randomnames -- API to generate random names of both genders.
/reconditematter/geomatrix   -- API to compute a matrix of geographic distances.
/reconditematter/geohash     -- API to compute a geohash of geographic coordinates.
/reconditematter/kfunction   -- API to compute Ripley's K function on the unit sphere.
`
	svc.HS200t(w, []byte(doc))
}
