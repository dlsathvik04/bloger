package blog

import (
	"os"
	"path"
	"path/filepath"

	"github.com/dlsathvik04/bloger/internal/utils"
)

type Blog struct {
	Path        string
	FrontMatter FrontMatter
	Markdown    Markdown
}

func NewBlog(path string) *Blog {
	frontMatterContent, markdownContent, err := utils.ReadAndSplitFile(path)
	if err != nil {
		return &Blog{}
	}
	frontMatter := NewFrontMatter(frontMatterContent)
	markdown := NewMarkdown(markdownContent)

	return &Blog{path, *frontMatter, *markdown}
}

func (b *Blog) Build(buildPath string) {
	blogBuildPath := filepath.Join(buildPath, filepath.Base(b.Path))
	// htmlPath := filepath.Join(blogBuildPath, "index.html")

	copyNonMdFiles(b.Path, blogBuildPath)
}

func copyNonMdFiles(src, dest string) error {
	return filepath.Walk(src, func(filePath string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		destPath := path.Join(dest, filePath[len(src)-1:])

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
