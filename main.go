package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"
	mux := http.NewServeMux()

	fileServer := http.FileServer(http.Dir("."))

	mux.Handle("/", fileServer)

	serv := &http.Server{
		Handler: mux,
		Addr:    ":" + port,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(serv.ListenAndServe())
}
