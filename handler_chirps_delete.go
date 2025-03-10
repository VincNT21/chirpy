package main

import (
	"net/http"

	"github.com/VincNT21/chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerDeleteChirp(w http.ResponseWriter, req *http.Request) {
	// Get chirp id from path
	chirpIDString := req.PathValue("chirpID")
	chirpID, err := uuid.Parse(chirpIDString)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid chirp ID", err)
		return
	}

	// Get access token from header
	accessToken, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't find access token", err)
		return
	}
	// Validate the JWT access token and get user ID associated
	userID, err := auth.ValidateJWT(accessToken, cfg.jwtsecret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "couldn't validate access token", err)
		return
	}

	// Get chirp by given id
	chirp, err := cfg.db.GetChirpById(req.Context(), chirpID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "no chirp found with provided id", err)
		return
	}

	// Check if ID from access token == chirp userID
	if userID != chirp.UserID {
		respondWithError(w, http.StatusForbidden, "You can't delete this chirp", nil)
		return
	}

	// If ok, delete the chirp
	err = cfg.db.DeleteChirpByID(req.Context(), chirp.ID)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't delete chirp from db", err)
	}

	// Respond
	w.WriteHeader(http.StatusNoContent)
}
