package main

import (
	"database/sql"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/VincNT21/chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) HandlerChangeToRed(w http.ResponseWriter, req *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserID uuid.UUID `json:"user_id"`
		}
	}

	// Get API key from header and compare it to store api
	apiKey, err := auth.GetAPIKey(req.Header)
	if err != nil {
		respondWithError(w, 401, "couldn't get API Key from header-Authorization", err)
		return
	}
	if apiKey != cfg.polkaKey {
		respondWithError(w, 401, "API key is invalid", err)
		return
	}

	// Get request body
	var params parameters
	decoder := json.NewDecoder(req.Body)
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't decode request's body", err)
		return
	}

	// Check if event is user.upgraded
	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// If ok, upgrade user
	err = cfg.db.ChangeToRed(req.Context(), params.Data.UserID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondWithError(w, 404, "user not found", err)
			return
		}
		respondWithError(w, 500, "couldn't update user to red", err)
		return
	}

	// If upgrade ok, respond
	w.WriteHeader(204)
}
