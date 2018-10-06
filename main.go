package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/metapods/metapods/handlers/webfinger"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/.well-known/webfinger", webfinger.Get).Methods("GET")

	http.HandleFunc("/health", healthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "🎙")
}
