package blog

import "github.com/gomarkdown/markdown"

type Markdown struct {
	Content string
	Html    string
}

func NewMarkdown(content string) *Markdown {
	htmlContent := markdown.ToHTML([]byte(content), nil, nil)
	return &Markdown{content, string(htmlContent)}
}
