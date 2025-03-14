package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"sync/atomic"

	"github.com/VincNT21/chirpy/internal/database"
	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

type apiConfig struct {
	fileserverHits atomic.Int32
	db             *database.Queries
	platform       string
	jwtsecret      string
	polkaKey       string
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	// Get the secret info from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	dbURL := os.Getenv("DB_URL")
	if dbURL == "" {
		log.Fatal("DB_URL must be set")
	}
	platform := os.Getenv("PLATFORM")
	if platform == "" {
		log.Fatal("PLATFORM must be set")
	}
	jwtsecret := os.Getenv("SECRET")
	if jwtsecret == "" {
		log.Fatal("SECRET must be set")
	}
	polkaKey := os.Getenv("POLKA_KEY")
	if polkaKey == "" {
		log.Fatal("POLKA_KEY must be set")
	}

	// Open a connection to database
	dbConnection, err := sql.Open("postgres", dbURL)
	if err != nil {
		log.Fatal("Error connecting to db")
	}

	// Create a *database.Queries to store in config struct
	dbQueries := database.New(dbConnection)

	// Init the api config struct
	apiCfg := apiConfig{
		fileserverHits: atomic.Int32{},
		db:             dbQueries,
		platform:       platform,
		jwtsecret:      jwtsecret,
		polkaKey:       polkaKey,
	}

	// Create the request multiplexer (router) that will matches incoming request to registered handlers/
	mux := http.NewServeMux()

	// Wrap the handler into a metrics incrementer
	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot))))

	// Register the file server thant will serve files from the root directory when called the /app/ path
	mux.Handle("/app/", fsHandler)

	// Register other handlers that can be called with path provided
	mux.HandleFunc("GET /api/healthz", handlerReadiness)

	mux.HandleFunc("GET /admin/metrics", apiCfg.handlerHits)
	mux.HandleFunc("POST /admin/reset", apiCfg.handlerReset)

	mux.HandleFunc("POST /api/users", apiCfg.handlerUsersCreate)
	mux.HandleFunc("PUT /api/users", apiCfg.handlerUserUpdate)
	mux.HandleFunc("POST /api/login", apiCfg.handlerLogin)

	mux.HandleFunc("POST /api/refresh", apiCfg.handlerRefresh)
	mux.HandleFunc("POST /api/revoke", apiCfg.handlerRevoke)

	mux.HandleFunc("POST /api/chirps", apiCfg.handlerCreateChirp)
	mux.HandleFunc("GET /api/chirps", apiCfg.HandlerAllChirps)
	mux.HandleFunc("GET /api/chirps/{chirpID}", apiCfg.HandlerSingletonChirp)
	mux.HandleFunc("DELETE /api/chirps/{chirpID}", apiCfg.HandlerDeleteChirp)

	mux.HandleFunc("POST /api/polka/webhooks", apiCfg.HandlerChangeToRed)

	// Create a new http server that will listens on port specified and uses the multiplexer for handling requests.
	srv := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)

	// Start the server and log a fatal error if it fails to start
	log.Fatal(srv.ListenAndServe())
}
