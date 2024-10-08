package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	fileServer := http.FileServer(http.Dir("./ui/static/"))
	mux.Handle("GET /static/", http.StripPrefix("/static", fileServer))
	mux.HandleFunc("GET /", home)

	log.Println("Starting server on port 5000")
	err := http.ListenAndServe(":5000", mux)
	log.Fatal(err)

}
