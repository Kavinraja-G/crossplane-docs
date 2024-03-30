package utils

import (
	"os"
	"path/filepath"
	"strings"
)

// FindYAMLFiles lists all YAML files in the specified directory
func FindYAMLFiles(directory string) ([]string, error) {
	var yamlFiles []string
	files, err := os.ReadDir(directory)
	if err != nil {
		return nil, err
	}

	for _, file := range files {
		if !file.IsDir() {
			ext := strings.ToLower(filepath.Ext(file.Name()))
			if ext == ".yaml" || ext == ".yml" {
				fullPath := filepath.Join(directory, file.Name())
				yamlFiles = append(yamlFiles, fullPath)
			}
		}
	}

	return yamlFiles, nil
}
