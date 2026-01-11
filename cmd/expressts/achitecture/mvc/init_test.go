package mvc

import (
	"os"
	"testing"

	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/types"
)

func TestInitMVCArchitecture(t *testing.T) {
	testProjectName := "test_mvc_project"
	project := types.ProjectInitStruct{
		ProjectName:        testProjectName,
		ProjectDescription: "A test mvc project",
		Options: types.ProjectInitOptions{
			ConfigType: "json",
		},
	}

	// Clean up after test
	defer os.RemoveAll(testProjectName)

	InitMVCArchitecture(project)

	// Verify Folder Structure exists
	expectedDirs := []string{
		testProjectName + "/src/model",
		testProjectName + "/src/controller",
		testProjectName + "/src/routes",
		testProjectName + "/src/middleware",
		testProjectName + "/src/utilities",
		testProjectName + "/src/utilities/dtos",
		testProjectName + "/src/utilities/exceptions",
	}

	for _, dir := range expectedDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("Expected directory %s to exist, but it does not", dir)
		}
	}

	// Verify key files exist
	expectedFiles := []string{
		testProjectName + "/src/bin/app.ts",
		testProjectName + "/src/model/users.ts",
		testProjectName + "/src/controller/users.ts",
		testProjectName + "/src/routes/users.ts",
		testProjectName + "/src/routes/index.ts",
		testProjectName + "/src/middleware/globalErrorHandler.ts",
		testProjectName + "/src/middleware/validateDtos.ts",
		testProjectName + "/src/utilities/httpResponseCode.ts",
		testProjectName + "/src/utilities/dtos/CreateUserDTO.ts",
		testProjectName + "/src/utilities/dtos/ResponseDTO.ts",
		testProjectName + "/src/utilities/dtos/ErrorResponseDTO.ts",
		testProjectName + "/src/utilities/exceptions/ConflictException.ts",
	}

	for _, file := range expectedFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("Expected file %s to exist, but it does not", file)
		}
	}
}
