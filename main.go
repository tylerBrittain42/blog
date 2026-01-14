package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

type articleCreator interface {
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
	mux.HandleFunc("GET /article/", cfg.generalArticleHandler)
	mux.HandleFunc("GET /article/{name}", cfg.specificArticleHandler)

	log.Printf("Serving on port %s\n", port)
	log.Fatal(server.ListenAndServe())

}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "template/index.html")
}

type config struct {
	templateDir string
	name        string
}
