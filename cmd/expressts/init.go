package expressts

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
	"github.com/shovan04/ExpressTS-in-GO/cmd/expressts/config"
)

var (
	projectName   string = "expr"
	projectDesc   string = "An Express Typescript project"
	projectArch   string = "layered"
	projectConfig string = "env"
	confirm       bool = true
)

func Init() {
	fmt.Println("✨ Welcome to ExpressTs — Fast. Typed. Opinionated.")
	fmt.Println("Let's set up your backend...")

	projectSetup := config.ProjectSetupForm(&projectName, &projectDesc, &projectArch, &projectConfig)

	if err := projectSetup.Run(); err != nil {
		fmt.Println("Error during setup:", err)
		os.Exit(1)
	}
	
	config.ProjectSummary(projectName, projectDesc, projectArch, projectConfig)

	huh.NewConfirm().
		Title("Proceed with project creation?").
		Value(&confirm).
		Run()

	if confirm {
		fmt.Println("🚀 Creating your ExpressTS project...")

		// TODO: Implement project scaffolding based on user choices
		// Check the project name
		// Check if the directory already exists

		fmt.Println("✅ Project created successfully!")
	} else {
		fmt.Println("❌ Project creation cancelled.")
	}

}

