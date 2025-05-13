package main

import (
	"fmt"

	"github.com/dlsathvik04/bloger/internal/blog"
	"github.com/dlsathvik04/bloger/internal/config"
)

func buildFromConfig(configPath string) {
	blogerConfig, err := config.NewConfig(configPath)
	if err != nil {
		fmt.Println(err)
	}
	blogerConfig.RunPreBuildCommand()
	bloger, err := blog.NewBlogger(blogerConfig.BlogsDirectory, blogerConfig.BuildDirectory)
	if err != nil {
		fmt.Println(err)
	}
	bloger.Build()
	fmt.Println(blogerConfig)
	blogerConfig.RunPostBuildCommand()
}

func main() {
	buildFromConfig("./")
}
