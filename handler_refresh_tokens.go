package main

import (
	"net/http"
	"time"

	"github.com/VincNT21/chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, req *http.Request) {
	type response struct {
		Token string `json:"token"`
	}

	// Get the token bearer from header
	refreshTokenString, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get refresh token from header", err)
		return
	}

	// Check if refresh token exists in db
	refreshToken, err := cfg.db.GetRefreshToken(req.Context(), refreshTokenString)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "refresh token doesn't exist in db", err)
		return
	}
	// Check if refresh token is revoked
	if refreshToken.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "refresh token has been been revoked", err)
		return
	}

	// Check refresh token expiration
	if time.Now().UTC().Compare(refreshToken.ExpiresAt) >= 0 {
		respondWithError(w, http.StatusUnauthorized, "refresh token is expired", err)
		return
	}

	// Create a new JWT access token (valid for 1 hour)
	accessToken, err := auth.MakeJWT(refreshToken.UserID, cfg.jwtsecret, time.Hour)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't make new JWT", err)
		return
	}

	// If ok, respond
	respondWithJSON(w, http.StatusOK, response{
		Token: accessToken,
	})
}

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, req *http.Request) {
	type response struct{}
	// Get the token bearer from header
	refreshTokenString, err := auth.GetBearerToken(req.Header)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't get refresh token from header", err)
		return
	}

	// Revoke Refresh token in DB (by setting its revoked_at field to NOW())
	err = cfg.db.RevokeRefreshToken(req.Context(), refreshTokenString)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "couldn't revoke refresh token in db", err)
		return
	}

	// If ok, respond
	w.WriteHeader(http.StatusNoContent)
}
