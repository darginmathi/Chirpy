package main

import (
	"encoding/json"
	"net/http"

	"github.com/darginmathi/Chirpy/internal/auth"
	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerUsersUpgrade(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Event string `json:"event"`
		Data  struct {
			UserId string `json:"user_id"`
		} `json:"data"`
	}

	apiKey, err := auth.GetApiKey(r.Header)
	if err != nil {
		respondWithError(w, http.StatusUnauthorized, "error retriving api key", err)
		return
	}

	if apiKey != cfg.polka_key {
		respondWithError(w, http.StatusUnauthorized, "wrong api key", err)
		return
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't decode parameters", err)
		return
	}

	if params.Event != "user.upgraded" {
		w.WriteHeader(http.StatusNoContent)
		return
	}

	id, err := uuid.Parse(params.Data.UserId)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Invalid UUID:", err)
		return
	}

	_, err = cfg.db.UpgradeUser(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "User not Found", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
