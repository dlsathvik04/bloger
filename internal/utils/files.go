package utils

import (
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

func CopyFile(srcFile, destFile string) error {
	src, err := os.Open(srcFile)
	if err != nil {
		return err
	}
	defer src.Close()
	dest, err := os.Create(destFile)
	if err != nil {
		return err
	}
	defer dest.Close()
	_, err = io.Copy(dest, src)
	if err != nil {
		return err
	}
	info, err := os.Stat(srcFile)
	if err != nil {
		return err
	}
	return os.Chmod(destFile, info.Mode())
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

func WriteTextToFile(path, content string) error {
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		return err
	}
	if err := os.WriteFile(path, []byte(content), 0644); err != nil {
		return err
	}
	return nil
}

func GetSubdirectories(rootPath string) ([]string, error) {
	var directories []string
	err := filepath.WalkDir(rootPath, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && path != rootPath {
			directories = append(directories, path)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return directories, nil
}

func CopyDirectory(src, dest string) error {
	src, err := filepath.Abs(src)
	if err != nil {
		return err
	}
	dest, err = filepath.Abs(dest)
	if err != nil {
		return err
	}
	if _, err := os.Stat(dest); os.IsNotExist(err) {
		if err := os.MkdirAll(dest, os.ModePerm); err != nil {
			return err
		}
	}
	return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if path == dest || strings.HasPrefix(path, dest+string(os.PathSeparator)) {
			return filepath.SkipDir
		}
		relPath, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		destPath := filepath.Join(dest, relPath)
		if d.IsDir() {
			return os.MkdirAll(destPath, os.ModePerm)
		}
		return CopyFile(path, destPath)
	})
}
