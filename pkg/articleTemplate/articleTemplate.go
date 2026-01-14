package articleTemplate

import (
	"bytes"
	"errors"
	"html/template"
	"strings"
)

type ArticleCreator interface {
	GetFilePath(dir string, name string) (string, error)
	GetTitle(fileName string) (string, error)
	GetContent(fileName string) (string, error)
}

type ArticleInfo struct {
	Title string
	Link  string
}

func GetArticleList(dir string) ([]ArticleInfo, error) {
	a1 := ArticleInfo{Title: "this is first", Link: "google.com"}
	te := []ArticleInfo{a1}
	return te, nil
}

func GetTemplate(a ArticleCreator, templateDir, name string) ([]byte, error) {
	type article struct {
		Title   string
		Content string
	}
	fullName, err := a.GetFilePath(templateDir, name)
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
