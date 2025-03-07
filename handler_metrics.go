package main

import (
	"fmt"
	"net/http"
)

// Handler method to return the fileserver Hits count
func (cfg *apiConfig) handlerHits(w http.ResponseWriter, req *http.Request) {
	// Set the content type header
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// Set the status code
	w.WriteHeader(http.StatusOK)

	// Write the response body
	count := fmt.Sprintf(`
	<html>
	
	<body>
		<h1>Welcome, Chirpy Admin</h1>
		<p>Chirpy has been visited %d times!</p>
	</body>
	</html>
		`, cfg.fileserverHits.Load())

	w.Write([]byte(count))

}

// Middleware that increments the fileserver Hits count
func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		// Increment counter before
		cfg.fileserverHits.Add(1)

		// Call the next handler
		next.ServeHTTP(w, req)
	})
}
