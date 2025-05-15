package main

import (
	"errors"
	"net/http"

	"github.com/darginmathi/Chirpy/internal/auth"
)

func (cfg *apiConfig) handlerRefresh(w http.ResponseWriter, r *http.Request) {
	type response struct {
		Token string `json:"token"`
	}

	str, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "token field empty", err)
		return
	}

	refreshToken, err := cfg.db.GetRefreshToken(r.Context(), str)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "refresh token not found", err)
		return
	}

	if refreshToken.RevokedAt.Valid {
		respondWithError(w, http.StatusUnauthorized, "invalid user", errors.New("token revoked"))
		return
	}

	token, err := auth.MakeJWT(refreshToken.UserID, cfg.jwt_secret)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "error making JWT", err)
		return
	}
	respondWithJSON(w, http.StatusOK, response{
		Token: token,
	})
}

func (cfg *apiConfig) handlerRevoke(w http.ResponseWriter, r *http.Request) {
	str, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "token field empty", err)
		return
	}

	err = cfg.db.RevokeRefreshToken(r.Context(), str)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't revoke session", err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
