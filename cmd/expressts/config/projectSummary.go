package config

import (
	"fmt"
)

func ProjectSummary(projectName string, projectDesc string, projectArch string, projectConfig string) {
	// Display summary and Confirmation
	fmt.Println("📦 Summary")
	fmt.Println("Project Name:", projectName)
	fmt.Println("Project Description:", projectDesc)
	fmt.Println("Project Architecture:", projectArch)
	fmt.Println("Configuration Style:", projectConfig)
}