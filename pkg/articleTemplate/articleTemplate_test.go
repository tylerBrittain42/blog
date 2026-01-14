package articleTemplate

import (
	"errors"
	"os"
	"strings"
	"testing"
)

type mockArticleCreator struct {
	filePath    string
	filePathErr error
	title       string
	titleErr    error
	content     string
	contentErr  error
}

func (m *mockArticleCreator) GetFilePath(dir string, name string) (string, error) {
	return m.filePath, m.filePathErr
}

func (m *mockArticleCreator) GetTitle(fileName string) (string, error) {
	return m.title, m.titleErr
}

func (m *mockArticleCreator) GetContent(fileName string) (string, error) {
	return m.content, m.contentErr
}

func TestGetTemplate(t *testing.T) {
	t.Run("should return error when GetFilePath fails", func(t *testing.T) {
		mock := &mockArticleCreator{
			filePathErr: errors.New("GetFilePath error"),
		}
		_, err := GetTemplate(mock, "someDir", "someName")
		if err == nil {
			t.Fatal("expected an error but got nil")
		}
		if err.Error() != "GetFilePath error" {
			t.Errorf("expected error 'GetFilePath error', got '%v'", err)
		}
	})

	t.Run("should return error when GetTitle fails", func(t *testing.T) {
		mock := &mockArticleCreator{
			titleErr: errors.New("GetTitle error"),
		}
		_, err := GetTemplate(mock, "someDir", "someName")
		if err == nil {
			t.Fatal("expected an error but got nil")
		}
		if err.Error() != "GetTitle error" {
			t.Errorf("expected error 'GetTitle error', got '%v'", err)
		}
	})

	t.Run("should return error when GetContent fails", func(t *testing.T) {
		mock := &mockArticleCreator{
			contentErr: errors.New("GetContent error"),
		}
		_, err := GetTemplate(mock, "someDir", "someName")
		if err == nil {
			t.Fatal("expected an error but got nil")
		}
		if err.Error() != "GetContent error" {
			t.Errorf("expected error 'GetContent error', got '%v'", err)
		}
	})

	t.Run("should return error for empty content", func(t *testing.T) {
		mock := &mockArticleCreator{
			content: "   \n\t",
		}
		_, err := GetTemplate(mock, "someDir", "someName")
		if err == nil {
			t.Fatal("expected an error but got nil")
		}
		if err.Error() != "article contents were empty" {
			t.Errorf("expected error 'article contents were empty', got '%v'", err)
		}
	})

	t.Run("should return error when template file does not exist", func(t *testing.T) {
		mock := &mockArticleCreator{
			filePath: "a/b.md",
			title:    "My Title",
			content:  "My Content",
		}
		_, err := GetTemplate(mock, "a", "b")
		if err == nil {
			t.Fatal("expected an error but got nil")
		}
		if !strings.Contains(err.Error(), "template/base.html") {
			t.Errorf("expected error to contain 'template/base.html', but it did not. got: %v", err)
		}
	})

	t.Run("should return rendered template on success", func(t *testing.T) {
		if err := os.Mkdir("template", 0755); err != nil && !os.IsExist(err) {
			t.Fatalf("failed to create template dir: %v", err)
		}
		t.Cleanup(func() {
			os.RemoveAll("template")
		})

		baseHTML := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>{{ .Title }}</h1>
    <p>{{ .Content }}</p>
    
</body>
</html>`
		if err := os.WriteFile("template/base.html", []byte(baseHTML), 0644); err != nil {
			t.Fatalf("failed to write base.html: %v", err)
		}

		mock := &mockArticleCreator{
			filePath: "a/b.md",
			title:    "My Title",
			content:  "My Content",
		}

		result, err := GetTemplate(mock, "a", "b")
		if err != nil {
			t.Fatalf("GetTemplate() returned an unexpected error: %v", err)
		}

		expected := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Document</title>
</head>
<body>
    <h1>My Title</h1>
    <p>My Content</p>
    
</body>
</html>`
		if string(result) != expected {
			t.Errorf("expected \n%s\n, got \n%s\n", expected, string(result))
		}
	})
}
