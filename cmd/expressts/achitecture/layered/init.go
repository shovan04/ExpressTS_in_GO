package layered

import (
	"fmt"

	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/config"
	dataLayered "github.com/shovan04/ExpressTS-in-GO/cmd/expressts/data/layered"
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

	// Create main application file
	createMainAppFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath:    project.ProjectName,
			FolderName: config.StrPointer("/src/bin"),
			FileName:   "app.ts",
		},
		Content: []byte(dataLayered.GetMainAppContent()),
	})

	if createMainAppFile != nil {
		fmt.Println("Error: create main application file")
		panic(createMainAppFile)
	}

	// Create Global Error Handler file
	createGlobalErrorHandlerFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath: project.ProjectName,
			FolderName: config.StrPointer("/src/middlewares"),
			FileName: "globalErrorHandler.ts",
		},
		Content: []byte(dataLayered.GetGlobalErrorHandlerContent()),
	})

	if createGlobalErrorHandlerFile != nil {
		fmt.Println("Error: create global error handler file")
		panic(createGlobalErrorHandlerFile)
	}

	// Create Validate DTOs Middleware file
	createValidateDTOsMiddlewareFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath: project.ProjectName,
			FolderName: config.StrPointer("/src/middlewares"),
			FileName: "validateDTOs.ts",
		},
		Content: []byte(dataLayered.GetValidateDTOClassContent()),
	})

	if createValidateDTOsMiddlewareFile != nil {
		fmt.Println("Error: create validate DTOs middleware file")
		panic(createValidateDTOsMiddlewareFile)
	}

	// Create exeption files
	createConflictExceptionFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath: project.ProjectName,
			FolderName: config.StrPointer("/src/exceptions"),
			FileName: "conflictExceptions.ts",
		},
		Content: []byte(dataLayered.GetConflictExceptionContent()),
	})

	if createConflictExceptionFile != nil {
		fmt.Println("Error: create conflict exception file")
		panic(createConflictExceptionFile)
	}
	createValidationExceptionFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath: project.ProjectName,
			FolderName: config.StrPointer("/src/exceptions"),
			FileName: "validationErrors.ts",
		},
		Content: []byte(dataLayered.GetValidationExceptionContent()),
	})

	if createValidationExceptionFile != nil {
		fmt.Println("Error: create validation exception file")
		panic(createValidationExceptionFile)
	}

	// Create DTO Class files
	createResponseDTOFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath: project.ProjectName,
			FolderName: config.StrPointer("/src/DTOClass"),
			FileName: "response.DTO.ts",
		},
		Content: []byte(dataLayered.GetResponseDTOContent()),
	})

	if createResponseDTOFile != nil {
		fmt.Println("Error: create response DTO file")
		panic(createResponseDTOFile)
	}

	createErrorResponseDTOFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath: project.ProjectName,
			FolderName: config.StrPointer("/src/DTOClass"),
			FileName: "errorResponse.DTO.ts",
		},
		Content: []byte(dataLayered.GetErrorResponseDTOContent()),
	})

	if createErrorResponseDTOFile != nil {
		fmt.Println("Error: create error response DTO file")
		panic(createErrorResponseDTOFile)
	}

	// Create Constant files
	createHttpResponseCodeFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath: project.ProjectName,
			FolderName: config.StrPointer("/src/constants"),
			FileName: "httpResponseCode.ts",
		},
		Content: []byte(shared.GetHttpResponseCodeContent()),
	})

	if createHttpResponseCodeFile != nil {
		fmt.Println("Error: create HTTP response code file")
		panic(createHttpResponseCodeFile)
	}

	// Create Config files - routes
	createConfigRouteFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath: project.ProjectName,
			FolderName: config.StrPointer("/src/configs/routes"),
			FileName: "wellcome.ts",
		},
		Content: []byte(dataLayered.GetRoutesConfigContent()),
	})

	if createConfigRouteFile != nil {
		fmt.Println("Error: create config route file")
		panic(createConfigRouteFile)
	}

	// Create routes files - wellcome route and main route file
	createMainRouteFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath: project.ProjectName,
			FolderName: config.StrPointer("/src/routes/"),
			FileName: "index.ts",
		},
		Content: []byte(dataLayered.GetMainRouterIndexContent()),
	})

	if createMainRouteFile != nil {
		fmt.Println("Error: create main route file")
		panic(createMainRouteFile)
	}

	createHelloRouteFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath: project.ProjectName,
			FolderName: config.StrPointer("/src/routes/hello"),
			FileName: "wellcome.ts",
		},
		Content: []byte(dataLayered.GetWellcomeRouterIndexContent()),
	})

	if createHelloRouteFile != nil {
		fmt.Println("Error: create hello route file")
		panic(createHelloRouteFile)
	}

	// Create Controller files - wellcome controller
	createWellcomeControllerFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath: project.ProjectName,
			FolderName: config.StrPointer("/src/controllers/hello"),
			FileName: "wellcome.controller.ts",
		},
		Content: []byte(dataLayered.GetWellcomeControllerContent()),
	})

	if createWellcomeControllerFile != nil {
		fmt.Println("Error: create wellcome controller file")
		panic(createWellcomeControllerFile)
	}

	// Create Srevice files - wellcome service
	createWellcomeServiceFile := config.WriteFile(types.WriteFileStruct{
		Path: types.FilePath{
			DirPath: project.ProjectName,
			FolderName: config.StrPointer("/src/services/hello"),
			FileName: "wellcome.service.ts",
		},
		Content: []byte(dataLayered.GetHelloServiceContent()),
	})
	
	if createWellcomeServiceFile != nil {
		fmt.Println("Error: create wellcome service file")
		panic(createWellcomeServiceFile)
	}
}