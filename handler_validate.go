package main

import (
	"fmt"
	"strings"
)

func handlerValidateChirp(chirp string) (string, error) {

	const maxChirpLength = 140

	if len(chirp) > maxChirpLength {
		return chirp, fmt.Errorf("chirp too long")
	}

	badWords := map[string]struct{}{
		"kerfuffle": {},
		"sharbert":  {},
		"fornax":    {},
	}

	return getCleanBody(chirp, badWords), nil
}

func getCleanBody(body string, badWords map[string]struct{}) string {
	words := strings.Split(body, " ")
	for i, word := range words {
		lowerWord := strings.ToLower(word)
		if _, ok := badWords[lowerWord]; ok {
			words[i] = "****"
		}
	}
	return strings.Join(words, " ")
}
