package blog

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

type FrontMatter struct {
	Title       string `yaml:"title"`
	Date        string `yaml:"date"`
	Author      string `yaml:"author"`
	Description string `yaml:"description"`
}

func NewFrontMatter(content string) *FrontMatter {
	var frontMatter FrontMatter
	err := yaml.Unmarshal([]byte(content), &frontMatter)
	if err != nil {
		fmt.Println(err)
		return &FrontMatter{}
	}
	return &frontMatter
}
