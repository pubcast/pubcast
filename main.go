package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	log.Println("running")
	http.HandleFunc("/health", helloWorldHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "🎙")
}
