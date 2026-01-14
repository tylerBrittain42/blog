package articleTemplate

import (
	"bytes"
	"errors"
	"html/template"
	"strings"
)

func (cfg *config) getTemplate(a articleCreator) ([]byte, error) {
type ArticleCreator interface {
	GetFilePath(dir string, name string) (string, error)
	GetTitle(fileName string) (string, error)
	GetContent(fileName string) (string, error)
}

func GetTemplate(a ArticleCreator, templateDir, name string) ([]byte, error) {
	type article struct {
		Title   string
		Content string
	}
	fullName, err := a.GetFilePath(cfg.templateDir, cfg.name)
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
