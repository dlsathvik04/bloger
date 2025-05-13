package config

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

type Config struct {
	BlogsDirectory   string
	BuildDirectory   string
	PreBuildCommand  string
	PostBuildCommand string
}

func NewConfig(path string) (*Config, error) {
	var configPath string
	fileInfo, err := os.Stat(path)
	if err != nil {
		return nil, fmt.Errorf("failed to access path '%s': %w", path, err)
	}
	if !fileInfo.IsDir() {
		configPath = path
	} else {
		configPath = filepath.Join(path, "bloger.json")
	}
	configFile, err := os.Open(configPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open config file '%s': %w", configPath, err)
	}
	defer configFile.Close()
	var config Config
	decoder := json.NewDecoder(configFile)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, fmt.Errorf("failed to decode config file '%s': %w", configPath, err)
	}
	return &config, nil
}

func runCommand(command string) error {
	if command == "" {
		return nil // No command to run
	}
	fmt.Println("Running Command:", command)
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/C", command)
	} else {
		cmd = exec.Command("sh", "-c", command)
	}
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr // Capture stderr as well
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("failed to execute command '%s': %w", command, err)
	}
	return nil
}

func (c *Config) RunPreBuildCommand() error {
	return runCommand(c.PreBuildCommand)
}

func (c *Config) RunPostBuildCommand() error {
	return runCommand(c.PostBuildCommand)
}
