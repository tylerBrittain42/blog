package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

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
	mux.HandleFunc("/toc", cfg.tableOfContentsHandler)
	mux.HandleFunc("GET /article/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/toc", http.StatusTemporaryRedirect)
	})
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
