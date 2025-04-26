package blog

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"text/template"

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
	err = bloger.buildBlogIndexHTML()
	if err != nil {
		fmt.Println(err)
	}
	// filesToCopy := []string{
	// 	"templates/blog_card.html",
	// 	"templates/footer.html",
	// 	"templates/header.html",
	// 	"templates/index.html",
	// }
	// for _, fileToCopy := range filesToCopy {
	// 	err = utils.CopyFile(fileToCopy, filepath.Join(bloger.buildDirectory, filepath.Base(fileToCopy)))
	// 	if err != nil {
	// 		return err
	// 	}
	//
	// }
	return err
}

func (bloger *Bloger) buildJsonDirectory() error {
	var blogList []map[string]any
	for _, blog := range bloger.Blogs {
		blogEntry := map[string]any{
			"FolderName":         blog.FolderName,
			"FrontMatterContent": blog.FrontMatterContent, // Assuming your Blog struct has this field
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

func (bloger *Bloger) buildBlogIndexHTML() error {
	outputPath := path.Join(bloger.buildDirectory, "index.html")
	htmlWriter, err := os.OpenFile(outputPath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("error opening index.html for writing: %w", err)
	}
	defer htmlWriter.Close()
	indexTempl, err := template.ParseFiles("./templates/index.html", "./templates/blog_card.html", "./templates/header.html")
	if err != nil {
		return fmt.Errorf("error parsing index templates: %w", err)
	}
	data := struct {
		Blogs []*Blog
	}{
		Blogs: bloger.Blogs,
	}
	err = indexTempl.Execute(htmlWriter, data)
	if err != nil {
		return fmt.Errorf("error executing index template: %w", err)
	}
	fmt.Println("Generated blog index HTML at:", outputPath)
	return nil
}
