package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tylerBrittain42/blog/pkg/validator"
)

type article interface {
	GetFilePath(dir string, name string) (string, error)
	GetTitle(fileName string) (string, error)
	GetContent(fileName string) (string, error)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	port := os.Getenv("PORT")
	mux := http.NewServeMux()
	server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	cfg := config{templateDir: os.Getenv("DIR")}

	mux.HandleFunc("/", indexHandler)
	mux.HandleFunc("GET /article/{name}", cfg.articleHandler)

	log.Printf("Serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/index.html")
}

func (cfg *config) articleHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	isSanitized, err := validator.IsAlphaNumeric(name)
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

	nameExt := name + ".md"
	canAccess, err := validator.IsAccessible(cfg.templateDir, nameExt)
	// TODO: improve me
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(fmt.Sprintf("error: %v\n", err)))
		return
	}
	if !canAccess {
		w.WriteHeader(http.StatusNotAcceptable)
		w.Write([]byte("Unable to access article"))
		return
	}

	http.ServeFile(w, r, "template/index.html")

	// doing template stuff here
	// remove line above and also this comment
	// use name from above

	// CONSIDER: how to handle spaces? swap name with underscore, but we already do not allow symbols so maybe table this as future issue
	// could also be diff in the metadata version so don't trip
	type article interface {
		GetTitle(fileName string) string
		GetContent(fileName string) string
	}

	// procedure
	// 1. get title - DONE
	// 2. get content
	// 2.1 if empty, then return a custom error message
	// 3. render template
	// 4. return template

}

type config struct {
	templateDir string
}
