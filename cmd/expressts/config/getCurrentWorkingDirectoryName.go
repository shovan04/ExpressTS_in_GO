package config

import (
	"fmt"
	"os"
	"path/filepath"
)

func GetCurrentWorkingDirectoryName() string {
	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("Error: geting current working directory name")
		return ""
	}
	return filepath.Base(wd)
}