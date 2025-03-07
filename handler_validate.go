package main

import (
	"encoding/json"
	"net/http"
	"strings"
)

func handlerValidity(w http.ResponseWriter, req *http.Request) {
	type returnVals struct {
		CleanedBody string `json:"cleaned_body"`
	}

	type parameters struct {
		Body string `json:"body"`
	}

	// Get the body from request
	decoder := json.NewDecoder(req.Body)
	params := parameters{}

	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}

	// Test if body is too long and respond with error (code=400) if yes
	const maxChirpLength = 140
	if len(params.Body) > maxChirpLength {
		respondWithError(w, http.StatusBadRequest, "Chirp is too long", nil)
		return
	}

	// Clean the body sentence
	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}
	cleaned := getCleanedBody(params.Body, badWords)

	// Respond with valid if ok
	respondWithJSON(w, http.StatusOK, returnVals{
		CleanedBody: cleaned,
	})
}

func getCleanedBody(body string, badWords map[string]struct{}) string {
	words := strings.Split(body, " ")
	for i, word := range words {
		loweredWord := strings.ToLower(word)
		if _, ok := badWords[loweredWord]; ok {
			words[i] = "****"
		}
	}
	return strings.Join(words, " ")
}
