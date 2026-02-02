package main

import (
	"html/template"
	"log"
	"net/http"
)

func respondWithError(w http.ResponseWriter, msg string, status int) {
	w.WriteHeader(status)
	t, err := template.ParseFiles("template/error.html")
	if err != nil {
		log.Printf("template parse error: %v\n", err)
		http.Error(w, msg, status)
		return
	}

	data := map[string]interface{}{
		"Status":       status,
		"ErrorMessage": msg,
	}

	if err := t.Execute(w, data); err != nil {
		log.Printf("template execute error: %v\n", err)
	}
}
