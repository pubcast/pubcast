package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/health", healthHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "🎙")
}
