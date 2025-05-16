package main

import (
	"net/http"
	"sort"

	"github.com/google/uuid"
)

func (cfg *apiConfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.db.GetChirps(r.Context())
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps", err)
		return
	}

	authorID := uuid.Nil
	authorIDString := r.URL.Query().Get("author_id")
	if authorIDString != "" {
		authorID, err = uuid.Parse(authorIDString)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Invalid author ID", err)
			return
		}
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		if authorID != uuid.Nil && dbChirp.UserID != authorID {
			continue
		}

		chirps = append(chirps, Chirp{
			ID:        dbChirp.ID,
			CreatedAt: dbChirp.CreatedAt,
			UpdatedAt: dbChirp.UpdatedAt,
			UserID:    dbChirp.UserID,
			Body:      dbChirp.Body,
		})
	}

	sortMtd := r.URL.Query().Get("sort")
	if sortMtd != "" {
		if sortMtd == "asc" {
			sort.Slice(chirps, func(i int, j int) bool {
				return chirps[i].CreatedAt.Before(chirps[j].CreatedAt)
			})
		} else if sortMtd == "desc" {
			sort.Slice(chirps, func(i int, j int) bool {
				return chirps[i].CreatedAt.After(chirps[j].CreatedAt)
			})
		}
	}

	respondWithJSON(w, http.StatusOK, chirps)
}

func (cfg *apiConfig) handlerChirpsGet(w http.ResponseWriter, r *http.Request) {
	str := r.PathValue("chirpID")

	id, err := uuid.Parse(str)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Invalid UUID:", err)
		return
	}

	chirp, err := cfg.db.GetChirp(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "could not get chirp", err)
		return
	}

	respondWithJSON(w, http.StatusOK, Chirp{
		ID:        chirp.ID,
		CreatedAt: chirp.CreatedAt,
		UpdatedAt: chirp.UpdatedAt,
		Body:      chirp.Body,
		UserID:    chirp.UserID,
	})
}
