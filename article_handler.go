package main

import (
	"log"
	"net/http"

	"github.com/tylerBrittain42/blog/pkg/articleTemplate"
	"github.com/tylerBrittain42/blog/pkg/basicArticle"
	"github.com/tylerBrittain42/blog/pkg/validator"
)

func (cfg *config) tableOfContentsHandler(w http.ResponseWriter, r *http.Request) {
	articles, err := articleTemplate.GetArticleList(cfg.templateDir)
	log.Println(articles)
	log.Println(articles[0].Title)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.ServeFile(w, r, "template/toc.html")
}

func (cfg *config) specificArticleHandler(w http.ResponseWriter, r *http.Request) {
	cfg.name = r.PathValue("name")
	isSanitized, err := validator.IsAlphaNumeric(cfg.name)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !isSanitized {
		respondWithError(w, "Invalid characters in name", http.StatusNotAcceptable)
		return
	}

	bArt := basicArticle.BasicArticle{}
	fullPath, err := bArt.GetFilePath(cfg.templateDir, cfg.name)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	canAccess, err := validator.IsAccessible(fullPath)
	if err != nil {
		respondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if !canAccess {
		respondWithError(w, "Unable to access article", http.StatusNotAcceptable)
		return
	}

	articleBytes, _ := articleTemplate.GetTemplate(bArt, cfg.templateDir, cfg.name)

	_, err = w.Write(articleBytes)
	if err != nil {
		log.Printf("Unable to write articleBytes, %v\n", err)
	}
}
