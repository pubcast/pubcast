package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pubcast/pubcast/handlers/organizations"
	"github.com/pubcast/pubcast/handlers/webfinger"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/.well-known/webfinger", webfinger.Get).Methods("GET")
	r.HandleFunc("/api/org/{slug}", organizations.Get).Methods("GET")
	r.HandleFunc("/health", healthHandler)

	fmt.Println("Serving on port :8080")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "🎙")
}
