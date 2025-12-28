package config

import (
	"fmt"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/data/shared"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/types"
)

// create package.json tsconfig.json .env and .gitignore files
func CreateProjectRootFiles(project types.ProjectInitStruct) error {
	createPackageJson := WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:  project.ProjectName,
			FileName: "package.json",
		},
		Content: []byte(shared.GetPackageJsonContent(project.ProjectName, project.ProjectDescription)),
	})

	if createPackageJson != nil {
		fmt.Println("Error: create package.json file")
		return createPackageJson
	}

	createTsConfig := WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:  project.ProjectName,
			FileName: "tsconfig.json",
		},
		Content: []byte(shared.GetTsConfigJsonContent()),
	})

	if createTsConfig != nil {
		fmt.Println("Error: create tsconfig.json file")
		return createTsConfig
	}

	createEnvFile := WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:  project.ProjectName,
			FileName: ".env",
		},
		Content: []byte(shared.GetENVFileContent()),
	})

	if createEnvFile != nil {
		fmt.Println("Error: create .env file")
		return createEnvFile
	}

	createGitIgnoreFile := WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:  project.ProjectName,
			FileName: ".gitignore",
		},
		Content: []byte(shared.GetGitIgnoreContent()),
	})

	if createGitIgnoreFile != nil {
		fmt.Println("Error: create .gitignore file")
		return createGitIgnoreFile
	}
	return nil
}
