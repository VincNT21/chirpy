package main

import (
	"encoding/json"
	"net/http"

	"github.com/VincNT21/chirpy/internal/auth"
	"github.com/VincNT21/chirpy/internal/database"
)

func (cfg *apiConfig) handlerUserUpdate(w http.ResponseWriter, req *http.Request) {
	type response struct {
		User
	}

	type parameters struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	// Get the access token from header
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

	// Get the body from request
	decoder := json.NewDecoder(req.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode parameters", err)
		return
	}

	// Hash user password
	hash, err := auth.HashPassword(params.Password)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't hash password", err)
		return
	}

	// Update user info
	user, err := cfg.db.UpdateUser(req.Context(), database.UpdateUserParams{
		Email:          params.Email,
		HashedPassword: hash,
		ID:             userID,
	})
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't update user info", err)
		return
	}

	// Respond
	respondWithJSON(w, http.StatusOK, response{
		User: User{
			ID:          user.ID,
			CreatedAt:   user.CreatedAt,
			UpdatedAt:   user.UpdatedAt,
			Email:       user.Email,
			IsChirpyRed: user.IsChirpyRed,
		}})

}
