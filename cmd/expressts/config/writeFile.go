package config

import (
	"os"
	"path/filepath"

	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/types"
)

func WriteFile(file types.WriteFileStruct) error {
	var dir string

	// Resolve directory path
	if file.Path.FolderName != nil {
		dir = filepath.Join(
			file.Path.DirPath,
			*file.Path.FolderName,
		)
	} else {
		dir = file.Path.DirPath
	}

	// Ensure directory exists
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	// Full file path
	fullPath := filepath.Join(dir, file.Path.FileName)

	// Write file
	return os.WriteFile(fullPath, []byte(file.Content), 0644)
}
