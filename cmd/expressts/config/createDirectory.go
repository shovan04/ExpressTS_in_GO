package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/types"
)

func CreateDirectory(dirName string) error {
	if err := os.MkdirAll(dirName, 0755); err != nil {
		return fmt.Errorf("failed to create directory %q: %w", dirName, err)
	}
	return nil
}

func CreateDirectories(projectName string, dirs []types.LayeredFolderStruct, prefixSlug ...string) error {
	prefix := ""
	if len(prefixSlug) > 0 {
		prefix = prefixSlug[0]
	}

	for _, dir := range dirs {
		path := filepath.Join(projectName, prefix, dir.FolderName)
		if err := os.MkdirAll(path, 0755); err != nil {
			return err
		}
	}
	return nil
}
