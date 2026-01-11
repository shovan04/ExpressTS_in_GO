package ddd

import (
	"fmt"

	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/config"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/data/ddd"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/data/shared"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/types"
)

func InitDDDArchitecture(project types.ProjectInitStruct) {
	// Create Project Directory
	createDirectory := config.CreateDirectory(project.ProjectName)
	if createDirectory != nil {
		panic(createDirectory)
	}

	// Create DDD Folder Structure
	createDDDFolders := config.CreateDirectories(project.ProjectName, GetDDDFolders(), "src")
	if createDDDFolders != nil {
		panic(createDDDFolders)
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
		Content: ddd.GetMainAppContent(),
	})
	if createMainAppFile != nil {
		fmt.Println("Error: create main app file")
		panic(createMainAppFile)
	}
	fmt.Println("✔ Core application files")

	// Create Shared Files
	// Http Response Code
	createHttpResponseCodeFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/constants"),
			FileName:   "httpResponseCode.ts",
		},
		Content: []byte(shared.GetHttpResponseCodeContent()),
	})
	if createHttpResponseCodeFile != nil {
		panic(createHttpResponseCodeFile)
	}

	// Create Domain Layer Files
	// Entity
	createEntityFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/domain/entities"),
			FileName:   "User.ts",
		},
		Content: ddd.GetUserEntityContent(),
	})
	if createEntityFile != nil {
		panic(createEntityFile)
	}

	// Repository Interface
	createRepoInterfaceFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/domain/repositories"),
			FileName:   "IUserRepository.ts",
		},
		Content: ddd.GetUserRepositoryInterfaceContent(),
	})
	if createRepoInterfaceFile != nil {
		panic(createRepoInterfaceFile)
	}

	// Create Application Layer Files
	// Use Case
	createUseCaseFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/application/use-cases"),
			FileName:   "CreateUserUseCase.ts",
		},
		Content: ddd.GetCreateUserUseCaseContent(),
	})
	if createUseCaseFile != nil {
		panic(createUseCaseFile)
	}

	// DTO
	createDTOFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/application/dtos"),
			FileName:   "CreateUserDTO.ts",
		},
		Content: ddd.GetCreateUserDTOContent(),
	})
	if createDTOFile != nil {
		panic(createDTOFile)
	}

	// Create Infrastructure Layer Files
	// Persistence
	createRepoImplFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/infrastructure/persistence"),
			FileName:   "InMemoryUserRepository.ts",
		},
		Content: ddd.GetInMemoryUserRepositoryContent(),
	})
	if createRepoImplFile != nil {
		panic(createRepoImplFile)
	}

	// Create Interface Layer Files
	// Controller
	createControllerFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/interfaces/http/controllers"),
			FileName:   "UserController.ts",
		},
		Content: ddd.GetUserControllerContent(),
	})
	if createControllerFile != nil {
		panic(createControllerFile)
	}

	// Routes
	createRoutesFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/interfaces/http/routes"),
			FileName:   "UserRoutes.ts",
		},
		Content: ddd.GetUserRoutesContent(),
	})
	if createRoutesFile != nil {
		panic(createRoutesFile)
	}

	createMainRouterFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/interfaces/http/routes"),
			FileName:   "index.ts",
		},
		Content: ddd.GetMainRouterIndexContent(),
	})
	if createMainRouterFile != nil {
		panic(createMainRouterFile)
	}

	fmt.Println("✔ Domain, Application, Infrastructure & Interface layers")
}
