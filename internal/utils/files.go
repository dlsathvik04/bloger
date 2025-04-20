package utils

import (
	"os"
	"path"
)

func CopyFile(src, dst string) error {
	input, err := os.ReadFile(src)
	if err != nil {
		return err
	}
	return os.WriteFile(dst, input, 0644)
}

func ReadAndSplitFile(blogPath string) (string, string, error) {
	markdownFilePath := path.Join(blogPath, "index.md")
	markdownFile, err := os.ReadFile(markdownFilePath)
	if err != nil {
		return "", "", err
	}
	content := string(markdownFile)

	if len(content) > 6 && content[0:3] == "---" {
		end := -1
		for i := 3; i < len(content)-2; i++ {
			if content[i:i+3] == "---" {
				end = i
				break
			}
		}
		if end != -1 {
			return content[3:end], content[end+3:], nil
		}
	}

	return "", content, nil
}
