package main

import (
	"fmt"

	"github.com/dlsathvik04/bloger/internal/blog"
)

func main() {
	blogsPath := "./blogs"
	buildPath := "./build"

	bloger, err := blog.NewBlogger(blogsPath, buildPath)
	if err != nil {
		fmt.Println(err)
	}
	bloger.Build()
}
