package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func changeDirectory(targetDir string, currentDir string) (string, error) {
	targetPath := filepath.Join(currentDir, targetDir)
	if _, err := os.Stat(targetPath); err != nil {
		return currentDir, fmt.Errorf("directory not exist")
	}
	return targetPath, nil
}
