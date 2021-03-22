package main

import (
	"log"
	"net/http"

	"davlist/internal/app"
)

func main() {
	const addr = "localhost:8016"

	s := &http.Server{
		Handler: app.DavlistMux,
		Addr:    addr,
	}

	log.Printf("Listening on http://%s ...", addr)
	log.Fatal(s.ListenAndServe())
}
