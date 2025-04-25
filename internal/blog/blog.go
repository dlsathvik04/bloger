package blog

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"text/template"

	"github.com/dlsathvik04/bloger/internal/utils"
)

type Blog struct {
	Path               string
	FrontMatterContent FrontMatter
	MarkdownContent    Markdown
}

func NewBlog(path string) (*Blog, error) {
	frontMatterContent, markdownContent, err := utils.ReadAndSplitFile(path)
	if err != nil {
		return &Blog{}, err
	}
	frontMatter := NewFrontMatter(frontMatterContent)
	markdown := NewMarkdown(markdownContent)

	blog := Blog{path, *frontMatter, *markdown}
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
	templ := template.Must(template.ParseFiles("./templates/blog.html"))
	fmt.Println("Read the template")
	err = templ.Execute(htmlWriter, b)
	if err != nil {
		return err
	}
	copyNonMdFiles(b.Path, blogBuildPath)
	return nil
}

func copyNonMdFiles(src, dest string) error {
	fmt.Println(src)
	fmt.Println(dest)
	return filepath.Walk(src, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		destPath := path.Join(dest, filePath[len(src):])
		fmt.Println(destPath)
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
