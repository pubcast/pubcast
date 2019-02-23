package main

import (
	"io"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	negronilogrus "github.com/meatballhat/negroni-logrus"
	"github.com/pubcast/pubcast/handlers/organizations"
	"github.com/pubcast/pubcast/handlers/webfinger"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/negroni"
)

func init() {
	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	log.SetOutput(os.Stdout)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/.well-known/webfinger", webfinger.Get).Methods("GET")
	r.HandleFunc("/api/org/{slug}", organizations.Get).Methods("GET")
	r.HandleFunc("/api/org", organizations.Create).Methods("POST")
	r.HandleFunc("/health", healthHandler)

	n := negroni.New()
	n.Use(negronilogrus.NewMiddleware())
	n.UseHandler(r)

	// Add support for PORT env
	port := "8080"
	if fromEnv := os.Getenv("PORT"); fromEnv != "" {
		port = fromEnv
	}

	log.WithField("port", port).Info("starting pubcast api server")
	log.Fatal(http.ListenAndServe(":"+port, n))
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "🎙")
}
