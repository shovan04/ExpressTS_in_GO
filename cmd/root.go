package root

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/huh"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/config"
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

	fmt.Println()

	if confirm {
		expressts.Init()
	} else {
		fmt.Println("❌ Project creation cancelled.")
	}

	fmt.Println()
	fmt.Println("✅ Project created successfully!")
	fmt.Print("\n\n")
	fmt.Printf("👉 To get started:\n\tcd %s\n\tpnpm up\n\tpnpm dev", project.ProjectName)
	fmt.Print("\n\n")
	fmt.Println("Happy hacking 🚀 Go fast 🏎️")
	fmt.Println()
}
