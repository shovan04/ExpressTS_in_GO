package mvc

import (
	"fmt"

	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/config"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/data/mvc"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/data/shared"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/types"
)

func InitMVCArchitecture(project types.ProjectInitStruct) {
	// Create Project Directory
	createDirectory := config.CreateDirectory(project.ProjectName)
	if createDirectory != nil {
		panic(createDirectory)
	}

	// Create MVC Folder Structure
	createMVCFolders := config.CreateDirectories(project.ProjectName, GetMVCFolders(), "src")
	if createMVCFolders != nil {
		panic(createMVCFolders)
	}
	fmt.Println("✔ Project structure")

	// Create Root Files (package.json, tsconfig, etc.)
	createRootFiles := config.CreateProjectRootFiles(project)
	if createRootFiles != nil {
		fmt.Println("Error: create root files")
		panic(createRootFiles)
	}

	// Create Main App File
	createMainAppFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/bin"),
			FileName:   "app.ts",
		},
		Content: mvc.GetMainAppContent(),
	})
	if createMainAppFile != nil {
		fmt.Println("Error: create main app file")
		panic(createMainAppFile)
	}
	fmt.Println("✔ Core application files")

	// Create Model
	createModelFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/model"),
			FileName:   "users.ts",
		},
		Content: mvc.GetUserModelContent(),
	})
	if createModelFile != nil {
		panic(createModelFile)
	}

	// Create Controller
	createControllerFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/controller"),
			FileName:   "users.ts",
		},
		Content: mvc.GetUserControllerContent(),
	})
	if createControllerFile != nil {
		panic(createControllerFile)
	}

	// Create Routes
	createRoutesFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/routes"),
			FileName:   "users.ts",
		},
		Content: mvc.GetUserRoutesContent(),
	})
	if createRoutesFile != nil {
		panic(createRoutesFile)
	}

	createMainRouterFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/routes"),
			FileName:   "index.ts",
		},
		Content: mvc.GetMainRouterIndexContent(),
	})
	if createMainRouterFile != nil {
		panic(createMainRouterFile)
	}

	// Create Middlewares
	createGlobalErrorHandler := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/middleware"),
			FileName:   "globalErrorHandler.ts",
		},
		Content: mvc.GetGlobalErrorHandlerContent(),
	})
	if createGlobalErrorHandler != nil {
		panic(createGlobalErrorHandler)
	}

	createValidateDto := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/middleware"),
			FileName:   "validateDtos.ts",
		},
		Content: mvc.GetValidateDTOContent(),
	})
	if createValidateDto != nil {
		panic(createValidateDto)
	}

	// Create Utilities (Configs/Constants)
	createConstants := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/utilities"),
			FileName:   "httpResponseCode.ts",
		},
		Content: []byte(shared.GetHttpResponseCodeContent()),
	})
	if createConstants != nil {
		panic(createConstants)
	}

	fmt.Println("✔ Model, Controller, Routes & Middleware layers")
}
