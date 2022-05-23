package utils

import (
	"os"
	"path/filepath"
)

// GetCurrentPath is a function for get MXfederal path
func GetCurrentPath() (string, error) {
	ex, err := os.Executable()
	if err != nil {
		return "", err
	}

	dir := filepath.Dir(ex) + "/.."
	path := filepath.Clean(dir)
	return path, nil
}
