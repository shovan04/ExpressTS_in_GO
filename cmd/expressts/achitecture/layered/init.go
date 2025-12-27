package layered

import (
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/config"
)

func InitLayeredArchitecture(projectName string) {
	// Implementation for initializing a layered architecture

	crteateDirectory := config.CreateDirectory(projectName)

	if crteateDirectory != nil {
		panic(crteateDirectory)
	}

	// Create subdirectories for layered architecture
	config.CreateDirectories(projectName, GetLayeredFolders(), "src")
}
