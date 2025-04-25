package blog

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/dlsathvik04/bloger/internal/utils"
)

type Bloger struct {
	blogsDirectory string
	buildDirectory string
	Blogs          []*Blog
}

func NewBlogger(blogsDirectory, buildDirectory string) (*Bloger, error) {
	blogDirs, err := utils.GetSubdirectories("./blogs")
	fmt.Println(blogDirs)
	if err != nil {
		return nil, err
	}
	blogs := make([]*Blog, len(blogDirs))
	fmt.Println(blogs)
	for ind, blogDir := range blogDirs {
		currentBlog, err := NewBlog(blogDir)
		if err != nil {
			return nil, err
		}
		blogs[ind] = currentBlog
	}
	fmt.Println(blogs)
	return &Bloger{
		blogsDirectory, buildDirectory, blogs,
	}, nil
}

func (bloger *Bloger) Build() error {
	utils.CopyDirectory("templates/static", path.Join(bloger.buildDirectory, "static"))
	for _, blog := range bloger.Blogs {
		fmt.Println(blog)
		err := blog.Build(bloger.buildDirectory)
		if err != nil {
			return err
		}
	}
	err := bloger.buildJsonDirectory()
	if err != nil {
		return err
	}

	err = utils.CopyFile("templates/index.html", filepath.Join(bloger.buildDirectory, "index.html"))
	if err != nil {
		return err
	}
	return nil
}

func (bloger *Bloger) buildJsonDirectory() error {
	var blogList []map[string]string
	for _, blog := range bloger.Blogs {
		blogName := path.Base(blog.Path)
		blogEntry := map[string]string{
			"name": blogName,
		}
		blogList = append(blogList, blogEntry)
	}
	jsonData, err := json.Marshal(blogList)
	if err != nil {
		return fmt.Errorf("error marshalling JSON: %w", err)
	}
	outputPath := path.Join(bloger.buildDirectory, "blogs.json") // Changed filename to blogs.json
	err = os.WriteFile(outputPath, jsonData, 0644)
	if err != nil {
		return fmt.Errorf("error writing JSON file: %w", err)
	}
	return nil
}
