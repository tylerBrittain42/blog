package articleTemplate

import (
	"bytes"
	"errors"
	"html/template"
	"os"
	"strings"
)

type ArticleCreator interface {
	GetFilePath(dir string, name string) (string, error)
	GetTitle(fileName string) (string, error)
	GetContent(fileName string) (string, error)
}

func CreateToc(dir string) ([]byte, error) {
	list, err := getArticleList(dir)
	if err != nil {
		return []byte{}, err
	}

	t, err := template.ParseFiles("template/toc.html")
	if err != nil {
		return []byte{}, err
	}

	var buf bytes.Buffer
	if err := t.Execute(&buf, list); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}

func getArticleList(dir string) ([]string, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return []string{}, err
	}
	if len(files) == 0 {
		return []string{}, nil
	}

	articleList := []string{}
	for _, v := range files {
		fName := v.Name()
		if fName[0] != '.' {
			articleList = append(articleList, strings.Split(fName, ".")[0])
		}
	}
	return articleList, nil
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
