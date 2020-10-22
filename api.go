// Copyright (c) 2019-2020 Leonid Kneller. All rights reserved.
// Licensed under the MIT license.
// See the LICENSE file for full license information.

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
		port = "80"
	}
	//
	R := mux.NewRouter()
	R.Handle("/api", handlers.LoggingHandler(os.Stderr, http.HandlerFunc(usage))).Methods("GET")
	//
	svc.RandomNames(R)
	svc.GeoMatrix(R)
	svc.GeoHash(R)
	svc.HashGeo(R)
	svc.KFunction(R)
	svc.Pop2010(R)
	svc.GreatEll(R)
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
/api/randomnames -- API to generate random names of both genders.
/api/geomatrix   -- API to compute a matrix of distances on the WGS1984 spheroid.
/api/geohash     -- API to convert geographic coordinates to a geohash.
/api/hashgeo     -- API to convert a geohash to geographic coordinates.
/api/kfunction   -- API to compute Ripley's K function on the unit sphere.
/api/pop2010     -- API to summarize US Census 2010 population.
/api/greatell    -- API for great ellipse navigation on the WGS1984 spheroid.
`
	svc.HS200t(w, []byte(doc))
}
