package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// Get the root path of project based on location of go.mod file
func GetProjectRootPath() (string, error) {
	// Get the current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}

	// Traverse up the directory tree until we find a go.mod file
	for {
		modFilePath := filepath.Join(cwd, "go.mod")
		if _, err := os.Stat(modFilePath); err == nil {
			return cwd, nil
		}

		// Move to the parent directory
		parentDir := filepath.Dir(cwd)
		if parentDir == cwd {
			// We reached the root directory without finding a go.mod file
			return "", fmt.Errorf("go.mod file not found in the project directory")
		}
		cwd = parentDir
	}
}
