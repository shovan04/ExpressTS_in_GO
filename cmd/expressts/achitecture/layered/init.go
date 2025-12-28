package layered

import (
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

	// Create subdirectories for layered architecture
	config.CreateDirectories(project.ProjectName, GetLayeredFolders(), "src")

	// create package.json tsconfig.json and .env
	config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:  project.ProjectName,
			FileName: "package.json",
		},
		Content: []byte(shared.GetPackageJsonContent(project.ProjectName, project.ProjectDescription)),
	})
	config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:  project.ProjectName,
			FileName: "tsconfig.json",
		},
		Content: []byte(shared.GetTsConfigJsonContent()),
	})

	config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:  project.ProjectName,
			FileName: ".env",
		},
		Content: []byte(shared.GetENVFileContent()),
	})
}
