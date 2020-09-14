package main

import (
	"log"
	"net/http"
)

func main() {

	r := mux.NewRouter()
	r.HandlerFunc("/", handle)

	log.Print("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}

func handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, World"))
}
