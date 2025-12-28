package config

import "github.com/charmbracelet/huh"



func ProjectSetupForm(projectName *string, projectDesc *string, projectArch *string, projectConfig *string) *huh.Form {
	return  huh.NewForm(
		// Step 1: Project Name
		huh.NewGroup(
			huh.NewInput().
				Title("Project Name").
				Placeholder("expr").
				Value(projectName),
		),

		// Step 2: Project Description
		huh.NewGroup(
			huh.NewInput().
				Title("Project Description").
				Placeholder("An Express Typescript project").
				Value(projectDesc),
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
				Value(projectArch),
		),

		// Step 4: Configuration
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Configuration style").
				Options(
					huh.NewOption("Environment variables (.env)", "env"),
				).
				Value(projectConfig),
		),
	)
}
