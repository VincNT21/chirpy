package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerAllChirps(w http.ResponseWriter, req *http.Request) {
	// Get ordered list of all chirps
	chirpList, err := cfg.db.GetAllChirps(req.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't retrieve list of all chirps", err)
		return
	}

	// Check if author_id query parameter is provided
	authorID := uuid.Nil
	authorIDString := req.URL.Query().Get("author_id")
	if authorIDString != "" {
		authorID, err = uuid.Parse(authorIDString)
		if err != nil {
			respondWithError(w, 400, "invalid author ID", err)
			return
		}
	}

	// Convert list to json friendly Chirp struct
	chirps := []Chirp{}
	for _, chirp := range chirpList {
		// Sort only chirps from author ID
		if authorID != uuid.Nil && chirp.UserID != authorID {
			continue
		}

		chirps = append(chirps, Chirp{
			ID:        chirp.ID,
			CreatedAt: chirp.CreatedAt,
			UpdatedAt: chirp.UpdatedAt,
			UserID:    chirp.UserID,
			Body:      chirp.Body,
		})
	}

	// Check for optionnal sort query parameter provided and revert sort if it is "desc"
	sortParameter := req.URL.Query().Get("sort")
	if sortParameter == "desc" {
		sort.Slice(chirps, func(i, j int) bool { return chirps[i].CreatedAt.After(chirps[j].CreatedAt) })
	}

	// Respond
	respondWithJSON(w, http.StatusOK, chirps)
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
