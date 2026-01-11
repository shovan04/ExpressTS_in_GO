package ddd

import (
	"os"
	"testing"

	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/types"
)

func TestInitDDDArchitecture(t *testing.T) {
	testProjectName := "test_ddd_project"

	// Cleanup before validation
	os.RemoveAll(testProjectName)
	defer os.RemoveAll(testProjectName)

	project := types.ProjectInitStruct{
		ProjectName:        testProjectName,
		ProjectDescription: "Test DDD Project",
		Options: types.ProjectInitOptions{
			ConfigType: "env",
		},
	}

	// Run the initialization
	InitDDDArchitecture(project)

	// Verify directories exist
	expectedDirs := []string{
		testProjectName + "/src/domain/entities",
		testProjectName + "/src/domain/repositories",
		testProjectName + "/src/application/use-cases",
		testProjectName + "/src/infrastructure/persistence",
		testProjectName + "/src/interfaces/http/controllers",
		testProjectName + "/src/constants",
	}

	for _, dir := range expectedDirs {
		if _, err := os.Stat(dir); os.IsNotExist(err) {
			t.Errorf("Expected directory %s to exist, but it does not", dir)
		}
	}

	// Verify key files exist
	expectedFiles := []string{
		testProjectName + "/src/domain/entities/User.ts",
		testProjectName + "/src/application/use-cases/CreateUserUseCase.ts",
		testProjectName + "/src/infrastructure/persistence/InMemoryUserRepository.ts",
		testProjectName + "/src/bin/app.ts",
		testProjectName + "/src/constants/httpResponseCode.ts",
	}

	for _, file := range expectedFiles {
		if _, err := os.Stat(file); os.IsNotExist(err) {
			t.Errorf("Expected file %s to exist, but it does not", file)
		}
	}

	// if t.Failed() {
	// 	// Keep files for debugging if failed
	// 	// return
	// }
}
