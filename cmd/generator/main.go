package main

import (
	"fmt"

	"github.com/dlsathvik04/bloger/internal/blog"
)

func main() {
	blogPath := "./blogs/blog1"

	blog := blog.NewBlog(blogPath)
	fmt.Println(blog.FrontMatter.Author)
	fmt.Println(blog.Markdown.Html)

	blog.Build("./build")
}
