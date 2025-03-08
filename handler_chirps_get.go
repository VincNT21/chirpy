package main

import (
	"net/http"

	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerAllChirps(w http.ResponseWriter, req *http.Request) {
	// Get ordered list of all chirps
	dbChirpsList, err := cfg.db.GetAllChirps(req.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't retrieve list of all chirps", err)
		return
	}

	// Convert list to json friendly Chirp struct
	jsonChirpsList := []Chirp{}
	for _, chirp := range dbChirpsList {
		jsonChirpsList = append(jsonChirpsList, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			Body:      chirp.Body,
			UserID:    chirp.UserID,
		})
	}

	// Respond
	respondWithJSON(w, http.StatusOK, jsonChirpsList)
}

func (cfg *apiConfig) HandlerSingletonChirp(w http.ResponseWriter, req *http.Request) {
	// Get ID from path
	idString := req.PathValue("chirpID")
	id, err := uuid.Parse(idString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	// Get chirp from db
	dbChirp, err := cfg.db.GetChirpById(req.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "no chirp found with provided id", err)
		return
	}

	// Respond
	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        dbChirp.ID,
		CreatedAt: dbChirp.CreatedAt,
		UpdatedAt: dbChirp.UpdatedAt,
		Body:      dbChirp.Body,
		UserID:    dbChirp.UserID,
	})
}
