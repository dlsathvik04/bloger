package blog

import "github.com/dlsathvik04/bloger/internal/utils"

type Blog struct {
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

	return &Blog{*frontMatter, *markdown}
}
