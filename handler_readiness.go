package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, req *http.Request) {
	// Set the content type header.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	// Set the status code
	w.WriteHeader(http.StatusOK)

	// Write the response body
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
