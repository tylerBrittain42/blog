package basicArticle

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

type BasicArticle struct {
}

func (BasicArticle) GetFilePath(dir string, name string) (string, error) {
	ext := ".md"
	if strings.TrimSpace(name) == "" {
		return "", errors.New("missing file name")
	}
	if !strings.HasSuffix(dir, "/") {
		dir += "/"
	}
	return dir + name + ext, nil

}

func (BasicArticle) GetTitle(filePath string) (string, error) {
	fileName := filepath.Base(filePath)
	dotCount := strings.Count(fileName, ".")
	if dotCount != 1 {
		return "", errors.New("filename must contain a single dot")
	}
	return strings.Split(fileName, ".")[0], nil
}

func (BasicArticle) GetContent(filePath string) (string, error) {
	dat, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	text := strings.TrimSpace(string(dat))
	if len(text) == 0 {
		return "", errors.New("no content found")
	}

	return text, nil
}
