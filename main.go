package main

import (
	"fmt"

	"github.com/charmbracelet/huh"
)

var (
	DefaultHost              string
	FrontendVirtualDirectory string
	DatabaseConnectionString string
)

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Default Host").
				Value(&DefaultHost),
			huh.NewInput().
				Title("Frontend Virtual Directory").
				Value(&FrontendVirtualDirectory),
			huh.NewInput().
				Title("Database Connection String").
				Value(&DatabaseConnectionString),
		),
	)

	if err := form.Run(); err != nil {
		// Handle any errors
		panic(err)
	}

	// Now you can use the exported variables:
	// DefaultHost, FrontendVirtualDirectory, and DatabaseConnectionString
	fmt.Printf("Default Host: %s\n", DefaultHost)
	fmt.Printf("Frontend Virtual Directory: %s\n", FrontendVirtualDirectory)
	fmt.Printf("Database Connection String: %s\n", DatabaseConnectionString)
}

