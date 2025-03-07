package main

import (
	"log"
	"net/http"
)

// Handler method to reser the fileserver Hits count
func (cfg *apiConfig) handlerReset(w http.ResponseWriter, req *http.Request) {
	// Check if platform is dev
	if cfg.platform != "dev" {
		w.WriteHeader(http.StatusForbidden)
		w.Write([]byte("Reset is only allowed in dev environment."))
		return
	}

	// Set the content type header
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Set the status code
	w.WriteHeader(http.StatusOK)

	// Reset the fileserver Hits count
	cfg.fileserverHits.Store(0)

	// Delete all users from users table in db
	err := cfg.db.DeleteUsers(req.Context())
	if err != nil {
		log.Printf("error while deleting users from db: %s", err)
	}

	// Write the response body
	w.Write([]byte("Hits reset to 0 and database Users reset to initial state"))
}
