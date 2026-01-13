package root

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/config"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/types"
)

var (
	projectName   string
	projectDesc   string
	projectArch   string
	projectConfig string
	confirm       bool = true
)

func CmdInit() {
	fmt.Println("✨ Welcome to ExpressTs — Fast. Typed. Opinionated.")
	fmt.Println("Let's set up your backend...")
	fmt.Println()

	projectSetup := config.ProjectSetupForm(&projectName, &projectDesc, &projectArch, &projectConfig)

	if err := projectSetup.Run(); err != nil {
		fmt.Println("Error during setup:", err)
		os.Exit(1)
	}

	// Check the project name
	// Check if the directory already exists or if the name is '.'
	if prjName := strings.ToLower(strings.TrimSpace(projectName)); prjName == "" || prjName == "." {
		projectName = config.GetCurrentWorkingDirectoryName()
	}
	if projectDesc == "" {
		projectDesc = "An Express Typescript project"
	}

	config.ProjectSummary(projectName, projectDesc, projectArch, projectConfig)
	fmt.Println()
	// Confirm to proceed
	huh.NewConfirm().
		Title("Proceed with project creation?").
		Value(&confirm).
		Run()

	if confirm {
		expressts.Init(types.ProjectInitStruct{
			ProjectName:        projectName,
			ProjectDescription: projectDesc,
			Options: types.ProjectInitOptions{
				ConfigType:  projectConfig,
				ProjectArch: &projectArch,
			},
		})
		config.ProjectCreated(projectName)
	} else {
		fmt.Println("❌ Project creation cancelled.")
	}

}
