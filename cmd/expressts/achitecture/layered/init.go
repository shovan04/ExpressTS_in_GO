package layered

import (
	"fmt"

	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/config"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/data/shared"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/types"
)

func InitLayeredArchitecture(project types.ProjectInitStruct) {
	// Implementation for initializing a layered architecture

	crteateDirectory := config.CreateDirectory(project.ProjectName)

	if crteateDirectory != nil {
		panic(crteateDirectory)
	}

	// Create sub-directories for layered architecture
	createLayeredFolders := config.CreateDirectories(project.ProjectName, GetLayeredFolders(), "src")
	if createLayeredFolders != nil {
		panic(createLayeredFolders)
	}

	// create package.json tsconfig.json and .env
	createPackageJson := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:  project.ProjectName,
			FileName: "package.json",
		},
		Content: []byte(shared.GetPackageJsonContent(project.ProjectName, project.ProjectDescription)),
	})

	if createPackageJson != nil {
		fmt.Println("Error: create package.json file")
		panic(createPackageJson)
	}

	createTsConfig := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:  project.ProjectName,
			FileName: "tsconfig.json",
		},
		Content: []byte(shared.GetTsConfigJsonContent()),
	})

	if createTsConfig != nil {
		fmt.Println("Error: create tsconfig.json file")
		panic(createTsConfig)
	}

	createEnvFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:  project.ProjectName,
			FileName: ".env",
		},
		Content: []byte(shared.GetENVFileContent()),
	})

	if createEnvFile != nil {
		fmt.Println("Error: create .env file")
		panic(createEnvFile)
	}
}
