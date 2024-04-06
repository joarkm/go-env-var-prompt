package main

import (
	"fmt"
	"net/url"
	"os"
	"regexp"

	"github.com/charmbracelet/huh"
)

var (
	DefaultHost              = "https://example.com"
	FrontendVirtualDirectory string
	DatabaseConnectionString string
	exportEnv                bool
)

func main() {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Default Host").
				Value(&DefaultHost).
				Validate(validateHost),
			huh.NewInput().
				Title("Frontend Virtual Directory").
				Value(&FrontendVirtualDirectory),
			huh.NewInput().
				Title("Database Connection String").
				Value(&DatabaseConnectionString),
		),
	)

	if err := form.Run(); err != nil {
		panic(err)
	}

	// Prompt user to export as environment variables
	writeEnvForm := huh.NewForm(
		huh.NewGroup(
			huh.NewConfirm().
				Title("Export as environment variables?").
				Value(&exportEnv),
		),
	)
	fmt.Printf("Default Host: %s\n", DefaultHost)
	fmt.Printf("Frontend Virtual Directory: %s\n", FrontendVirtualDirectory)
	fmt.Printf("Database Connection String: %s\n", DatabaseConnectionString)

	writeEnvForm.Run()
	if exportEnv {
		os.Setenv("DEFAULT_HOST", DefaultHost)
		os.Setenv("FRONTEND_VIRTUAL_DIR", FrontendVirtualDirectory)
		os.Setenv("DB_CONNECTION_STRING", DatabaseConnectionString)
		fmt.Println("Environment variables exported successfully.")
	}
}

func validateHost(input string) error {
	// Validate if the input is a valid URL with an optional port
	u, err := url.Parse(input)
	if err != nil {
		return fmt.Errorf("invalid URL format")
	}

	// Check if the host is not empty and has a valid format
	if u.Host == "" {
		return fmt.Errorf("host cannot be empty")
	}

	// Optional: Check if the port is valid (if provided)
	portPattern := regexp.MustCompile(`^\d+$`)
	if u.Port() != "" && !portPattern.MatchString(u.Port()) {
		return fmt.Errorf("invalid port format")
	}

	return nil
}

