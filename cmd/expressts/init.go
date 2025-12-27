package expressts

import (
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

func Init() {
	var (
		projectName   string
		projectDesc   string
		projectArch   string
		projectConfig string
		confirm       bool = true
	)

	fmt.Println("✨ Welcome to ExpressTs — Fast. Typed. Opinionated.")
	fmt.Println("Let's set up your backend...")

	projectSetup := huh.NewForm(
		// Step 1: Project Name
		huh.NewGroup(
			huh.NewInput().
				Title("Project Name").
				Placeholder("expr").
				Value(&projectName),
		),

		// Step 2: Project Description
		huh.NewGroup(
			huh.NewInput().
				Title("Project Description").
				Placeholder("An Express Typescript project").
				Value(&projectDesc),
		),

		// Step 3: Architecture
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your project architecture").
				Options(
					huh.NewOption("Layered Service Architecture (recommended)", "layered"),
					huh.NewOption("DDD (Domain-Driven Design)", "ddd"),
					huh.NewOption("Minimal (MVC-style)", "minimal"),
					huh.NewOption("Clean Architecture (under development)", "clean"),
				).
				Value(&projectArch),
		),

		// Step 4: Configuration
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Configuration style").
				Options(
					huh.NewOption("Environment variables (.env)", "env"),
					huh.NewOption("YAML", "yaml"),
					huh.NewOption("TOML", "toml"),
				).
				Value(&projectConfig),
		),
	)

	if err := projectSetup.Run(); err != nil {
		fmt.Println("Error during setup:", err)
		os.Exit(1)
	}
	// Display summary and Confirmation
	fmt.Println("📦 Summary")
	fmt.Println("Project Name:", projectName)
	fmt.Println("Project Description:", projectDesc)
	fmt.Println("Project Architecture:", projectArch)
	fmt.Println("Configuration Style:", projectConfig)

	huh.NewConfirm().
		Title("Proceed with project creation?").
		Value(&confirm).
		Run()
	
	if confirm {
		fmt.Println("🚀 Creating your ExpressTS project...\n")
		// Here you would add the logic to scaffold the project based on user input
		fmt.Println("✅ Project created successfully!")
	} else {
		fmt.Println("❌ Project creation cancelled.")
	}

}
