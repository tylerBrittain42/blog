package main

import (
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/tylerBrittain42/blog/internal/templates"
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
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	component := templates.Article("speed", "doog")
	if err := component.Render(r.Context(), w); err != nil {
		log.Printf("render error: %v", err)
		http.Error(w, "Render failed", http.StatusInternalServerError)
		return
	}
}

type config struct {
	templateDir string
	name        string
}
