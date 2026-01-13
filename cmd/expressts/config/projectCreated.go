package config

import "fmt"

func ProjectCreated(prjName string) {
	fmt.Println("🚀 Project created successfully!")
	fmt.Println()
	fmt.Println("👉 To get started:")
	fmt.Println("\tcd", prjName)
	fmt.Println("\tpnpm up")
	fmt.Println("\tpnpm dev")
	fmt.Println()
	fmt.Println("Happy hacking 🚀 Go fast 🏎️")
	fmt.Println()
}
