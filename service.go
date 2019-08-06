package main

import (
	"GameService/game"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/karlmoad/authentication"
	"log"
	"net/http"
	"os"
)

const auth_aud = "AUTHENTICATION_AUDIENCE"
const auth_iss = "AUTHENTICATION_ISSUER"

func main() {

	audience := os.Getenv(auth_aud)
	issuer := os.Getenv(auth_iss)

	log.SetOutput(os.Stdout)

	// Init authentication provider
	auth, err := authentication.NewAuthenticationProvider(authentication.GOOGLE, issuer, audience)
	if err != nil {
		log.Printf("[ERROR] Issue initializing auth provider: %s", err.Error())
		return
	}

	// Setup the router and routes
	router := mux.NewRouter()
	router.HandleFunc("/", HealthCheckHandler).Methods("GET")
	router.HandleFunc("/draw", auth.AuthenticationHandler(game.GameDrawHandler)).Methods("GET")
	router.HandleFunc("/draw/{count}", auth.AuthenticationHandler(game.GameDrawHandler)).Methods("GET")

	// CORS options
	headersOk := handlers.AllowedHeaders([]string{"Content-Type", "Authorization","Content-Length", "Accept"})
	originsOk := handlers.AllowedOrigins([]string{"*"})  // allow all inbound domains
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS", "DELETE"})

	// Start Service
	log.Println("Starting Service...")

	err =http.ListenAndServe(":30200", handlers.CORS(headersOk, originsOk, methodsOk)(handlers.LoggingHandler(os.Stdout, router)))

	if err != nil {
		log.Printf("[ERROR] On Service Start: %s", err.Error())
	}

	log.Println("___ END OF LINE ___")

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// Per K8s ingress rules only healthy endpoint can be routed
	// reply with http.200 to signal healthy status
	w.WriteHeader(http.StatusOK)
	return
	
}
