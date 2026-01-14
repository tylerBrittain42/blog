package main

import (
	"bytes"
	"errors"
	"html/template"
	"log"
	"net/http"
	"strings"

	"github.com/tylerBrittain42/blog/pkg/basicArticle"
	"github.com/tylerBrittain42/blog/pkg/validator"
)

func (cfg *config) generalArticleHandler(w http.ResponseWriter, r *http.Request) {
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

	// doing template stuff here
	// remove line above and also this comment
	// use name from above

	// CONSIDER: how to handle spaces? swap name with underscore, but we already do not allow symbols so maybe table this as future issue
	// could also be diff in the metadata version so don't trip

	articleBytes, _ := cfg.getTemplate(bArt)
	_, err = w.Write(articleBytes)
	if err != nil {
		log.Printf("Unable to write articleBytes, %v\n", err)
	}
}

func (cfg *config) getTemplate(a articleCreator) ([]byte, error) {
	type article struct {
		Title   string
		Content string
	}
	fullName, err := a.GetFilePath(cfg.templateDir, cfg.name)
	if err != nil {
		return nil, err
	}

	title, err := a.GetTitle(fullName)
	if err != nil {
		return nil, err
	}

	content, err := a.GetContent(fullName)
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(content) == "" {
		return nil, errors.New("article contents were empty")
	}

	t, err := template.ParseFiles("template/base.html")
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	data := article{
		Title:   title,
		Content: content,
	}
	if err := t.Execute(&buf, data); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
