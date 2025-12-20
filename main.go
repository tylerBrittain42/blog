package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/tylerBrittain42/blog/pkg/helper"
)

func main() {
	port := "8080"
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("GET /article/{name}", articleHandler)

	log.Printf("Serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hello from a byte string"))
}

func articleHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	isSanitized, err := helper.IsAlphaNumeric(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %v\n", err)))
		return
	}
	if isSanitized == false {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Invalid characters in name"))
		return
	}

	canAccess, err := helper.IsAccessable(name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %v\n", err)))
		return
	}
	if canAccess == false {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Unable to access article"))
		return
	}

	http.ServeFile(w, r, "template/index.html")

}
