package articleTemplate

import (
	"errors"
	"os"
	"path/filepath"
	"reflect"
	"sort"
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

func TestCreateToc(t *testing.T) {
	t.Run("should return error when directory does not exist", func(t *testing.T) {
		_, err := CreateToc("non-existent-dir")
		if err == nil {
			t.Fatal("expected an error but got nil")
		}
	})

	t.Run("should return error when template file does not exist", func(t *testing.T) {
		dir := t.TempDir()
		if err := os.WriteFile(filepath.Join(dir, "test.md"), []byte("content"), 0644); err != nil {
			t.Fatalf("failed to create test file: %v", err)
		}

		_, err := CreateToc(dir)
		if err == nil {
			t.Fatal("expected an error but got nil")
		}
		if !strings.Contains(err.Error(), "template/toc.html") {
			t.Errorf("expected error to contain 'template/toc.html', but it did not. got: %v", err)
		}
	})

	t.Run("should return rendered template on success", func(t *testing.T) {
		if err := os.Mkdir("template", 0755); err != nil && !os.IsExist(err) {
			t.Fatalf("failed to create template dir: %v", err)
		}
		t.Cleanup(func() {
			os.RemoveAll("template")
		})

		tocHTML := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>My Blog - TOC</title>
</head>
<body>
    <h1>Table of Contents</h1>
    <p>List below</p>

    <ul>
    {{range .}}
        <li><a href="article/{{.}}">{{.}}</a></li>
    {{end}}
    </ul>
    
</body>
</html>`
		if err := os.WriteFile("template/toc.html", []byte(tocHTML), 0644); err != nil {
			t.Fatalf("failed to write toc.html: %v", err)
		}

		dir := t.TempDir()
		if err := os.WriteFile(filepath.Join(dir, "article1.md"), []byte("content1"), 0644); err != nil {
			t.Fatalf("failed to create article1: %v", err)
		}
		if err := os.WriteFile(filepath.Join(dir, "article2.md"), []byte("content2"), 0644); err != nil {
			t.Fatalf("failed to create article2: %v", err)
		}

		result, err := CreateToc(dir)
		if err != nil {
			t.Fatalf("CreateToc() returned an unexpected error: %v", err)
		}

		resultStr := string(result)
		if !strings.Contains(resultStr, `<a href="article/article1">article1</a>`) {
			t.Errorf("expected result to contain '<a href=\"article/article1\">article1</a>', got: %s", resultStr)
		}
		if !strings.Contains(resultStr, `<a href="article/article2">article2</a>`) {
			t.Errorf("expected result to contain '<a href=\"article/article2\">article2</a>', got: %s", resultStr)
		}
		if !strings.Contains(resultStr, "<h1>Table of Contents</h1>") {
			t.Errorf("expected result to contain '<h1>Table of Contents</h1>', got: %s", resultStr)
		}
	})
}

func TestGetArticleList(t *testing.T) {
	cases := []struct {
		name          string
		filesToCreate []string
		expected      []string
		expectErr     bool
		useInvalidDir bool
	}{
		{
			name:          "empty directory",
			filesToCreate: []string{},
			expected:      []string{},
			expectErr:     false,
		},
		{
			name:          "one file",
			filesToCreate: []string{"test.md"},
			expected:      []string{"test"},
			expectErr:     false,
		},
		{
			name:          "multiple files",
			filesToCreate: []string{"test1.md", "test2.html"},
			expected:      []string{"test1", "test2"},
			expectErr:     false,
		},
		{
			name:          "non-existent directory",
			useInvalidDir: true,
			expectErr:     true,
		},
		{
			name:          "ignore dot files",
			filesToCreate: []string{".DS_Store", "test1.md"},
			expected:      []string{"test1"},
			expectErr:     false,
		},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			var dir string
			if tc.useInvalidDir {
				dir = "non-existent-dir"
			} else {
				dir = t.TempDir()
				for _, fileName := range tc.filesToCreate {
					if err := os.WriteFile(filepath.Join(dir, fileName), []byte("content"), 0644); err != nil {
						t.Fatalf("failed to create file: %v", err)
					}
				}
			}

			result, err := getArticleList(dir)

			if tc.expectErr {
				if err == nil {
					t.Fatal("expected an error but got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if !sliceIsEqual(result, tc.expected) {
				t.Errorf("expected %+v, got %+v", tc.expected, result)
			}
		})
	}
}

func sliceIsEqual(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	// The order of elements from reading a directory is not guaranteed.
	// To have a consistent comparison, we sort both slices.
	sort.Strings(a)
	sort.Strings(b)

	return reflect.DeepEqual(a, b)
}
