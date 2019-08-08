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
	router.HandleFunc("/", HealthCheckHandlerRawRoot).Methods("GET")
	router.HandleFunc("/games/", HealthCheckHandler).Methods("GET")
	router.HandleFunc("/games/draw", auth.AuthenticationHandler(game.GameDrawHandler)).Methods("GET")
	router.HandleFunc("/games/draw/{count}", auth.AuthenticationHandler(game.GameDrawHandler)).Methods("GET")

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

// Per K8s ingress rules only healthy endpoint can be routed
// reply with http.200 to signal healthy status,
// TODO: stand in for now add better Health Check functionality when needed

func HealthCheckHandlerRawRoot(w http.ResponseWriter, r *http.Request) {
	// identify which route is called for root health check
	// given GKE Ingress does not handle rewrite target
	log.Printf("Health Check to RAW ROOT path /")
	w.WriteHeader(http.StatusOK)
	return
	
}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	// identify which route is called for root health check
	// given GKE Ingress does not handle rewrite target
	log.Printf("Health Check to root /games/")
	w.WriteHeader(http.StatusOK)
	return

}