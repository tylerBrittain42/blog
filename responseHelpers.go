package main

import (
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	_, err := w.Write([]byte(msg))
	if err != nil {
		log.Printf("error: %v\n", err)
	}

}
