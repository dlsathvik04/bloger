package blog

import (
	"os"
	"path"
	"path/filepath"
	"regexp"
	"text/template"

	"github.com/dlsathvik04/bloger/internal/utils"
)

type Blog struct {
	Path               string
	FrontMatterContent FrontMatter
	MarkdownContent    Markdown
	FolderName         string
	CoverImage         string
}

func getFirstImagePath(htmlContent string) string {
	re := regexp.MustCompile(`<img\s+[^>]*?src\s*=\s*["']([^"']+)["'][^>]*?>`)
	match := re.FindStringSubmatch(htmlContent)
	if len(match) > 1 {
		return match[1]
	}
	return ""
}

func NewBlog(path string) (*Blog, error) {
	frontMatterContent, markdownContent, err := utils.ReadAndSplitFile(path)
	if err != nil {
		return &Blog{}, err
	}
	frontMatter := NewFrontMatter(frontMatterContent)
	markdown := NewMarkdown(markdownContent)
	folderName := filepath.Base(path)
	coverImage := getFirstImagePath(markdown.Html)
	blog := Blog{path, *frontMatter, *markdown, folderName, coverImage}
	return &blog, nil
}

func (b *Blog) Build(buildPath string) error {
	blogBuildPath := filepath.Join(buildPath, filepath.Base(b.Path))
	htmlPath := filepath.Join(blogBuildPath, "index.html")
	if err := os.MkdirAll(filepath.Dir(htmlPath), 0755); err != nil {
		return err
	}
	htmlWriter, err := os.OpenFile(htmlPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	templ := template.Must(template.ParseFiles("./templates/blog.html", "./templates/header.html"))
	err = templ.Execute(htmlWriter, b)
	if err != nil {
		return err
	}
	copyNonMdFiles(b.Path, blogBuildPath)
	return nil
}

func copyNonMdFiles(src, dest string) error {
	return filepath.Walk(src, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		destPath := path.Join(dest, filePath[len(src):])
		if !info.IsDir() && info.Name() != "index.md" {
			destDir := path.Dir(destPath)
			if _, err := os.Stat(destDir); os.IsNotExist(err) {
				if err := os.MkdirAll(destDir, 0755); err != nil {
					return err
				}
			}
			if err := utils.CopyFile(filePath, destPath); err != nil {
				return err
			}
		}
		return nil
	})
}
