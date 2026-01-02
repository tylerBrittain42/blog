package basicArticle

import (
	"errors"
	"os"
	"path/filepath"
	"strings"
)

func GetTitle(filePath string) (string, error) {
	fileName := filepath.Base(filePath)
	dotCount := strings.Count(fileName, ".")
	if dotCount != 1 {
		return "", errors.New("filename must contain a single dot")
	}
	return strings.Split(fileName, ".")[0], nil
}

func GetContent(filePath string) (string, error) {
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
