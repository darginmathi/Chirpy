package main

import (
	"net/http"

	"github.com/darginmathi/Chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsDelete(w http.ResponseWriter, r *http.Request) {
	str := r.PathValue("chirpID")

	id, err := uuid.Parse(str)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Invalid UUID:", err)
		return
	}

	token, err := auth.GetBearerToken(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "token field empty", err)
		return
	}

	user_id, err := auth.ValidateJWT(token, cfg.jwt_secret)
	if err != nil {
		respondWithError(w, http.StatusForbidden, "Forbidden", err)
		return
	}

	chirp, err := cfg.db.GetChirp(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "could not get chirp", err)
		return
	}

	if chirp.UserID != user_id {
		respondWithError(w, http.StatusForbidden, "invalid user", err)
		return
	}

	err = cfg.db.DeleteChirp(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "could not get chirp", err)
		return
	}
	respondWithJSON(w, http.StatusNoContent, nil)
}
