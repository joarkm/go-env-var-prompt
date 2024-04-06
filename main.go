package main

import (
	"fmt"
	"os"

	"strconv"

	"github.com/charmbracelet/huh"
	"github.com/go-playground/validator/v10"
)

var (
	DefaultHost                  string
	FrontendVirtualDirectory     string
	DatabaseConnectionString     string
	ExportEnv                    bool
	AccessTokenLifetimeInMinutes string
	SecurityTokenServiceUrl      string
)

// use a single instance of Validate, it caches struct info
var validate validator.Validate

func main() {

	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Default Host").
				Placeholder("example.com").
				Value(&DefaultHost).
				Validate(validateHost),
			huh.NewInput().
				Title("Frontend Virtual Directory").
				Value(&FrontendVirtualDirectory).
				Validate((validateVirtualDirectory)),
			huh.NewInput().
				Title("Database Connection String").
				Value(&DatabaseConnectionString),
			huh.NewInput().
				Title("Url for security token service").
				Placeholder("https://example.com").
				Value(&SecurityTokenServiceUrl).
				Validate((validateUrl)),
			huh.NewInput().
				Title("Access token lifetime in minutes").
				Value(&AccessTokenLifetimeInMinutes).
				Suggestions([]string{"3600"}).
				Validate(validateAccessTokenLifetime),
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
				Value(&ExportEnv),
		),
	)
	fmt.Printf("Default Host: %s\n", DefaultHost)
	fmt.Printf("Frontend Virtual Directory: %s\n", FrontendVirtualDirectory)
	fmt.Printf("Database Connection String: %s\n", DatabaseConnectionString)
	fmt.Printf("Access token lifetime in minutes: %s\n", AccessTokenLifetimeInMinutes)
	fmt.Printf("Url for security token service: %s\n", SecurityTokenServiceUrl)

	writeEnvForm.Run()
	if ExportEnv {
		os.Setenv("DEFAULT_HOST", DefaultHost)
		os.Setenv("FRONTEND_VIRTUAL_DIR", FrontendVirtualDirectory)
		os.Setenv("DB_CONNECTION_STRING", DatabaseConnectionString)
		os.Setenv("ACCESS_TOKEN_LIFETIME_IN_MINUTES", AccessTokenLifetimeInMinutes)
		os.Setenv("SECURITY_TOKEN_SERVICE_URL", SecurityTokenServiceUrl)
		fmt.Println("Environment variables exported successfully.")
	}
}

func validateAccessTokenLifetime(input string) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	errs := formatPotentialValidationErrors(validate.Var(input, "required,number"))
	if errs != nil {
		return errs
	}
	// Input should now be safe to parse as a number
	lifetime, errs := strconv.ParseInt(input, 10, 64)
	if errs != nil {
		return errs
	}
	return formatPotentialValidationErrors(validate.Var(lifetime, "gte=1"))
}

func validateUrl(input string) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return formatPotentialValidationErrors(validate.Var(input, "required,http_url"))
}

func validateVirtualDirectory(input string) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	// No suitable validator found for virtual directories so went with the closest I could find
	// TODO: Define custom validator
	return formatPotentialValidationErrors(validate.Var(input, "required,alphanum"))
}

func validateHost(input string) error {
	validate := validator.New(validator.WithRequiredStructEnabled())
	return formatPotentialValidationErrors(validate.Var(input, "required,fqdn"))
}

func formatPotentialValidationErrors(errs error) error {
	if errs != nil {
		return fmt.Errorf(errs.Error()) // output: Key: "" Error:Field validation for "" failed on the "email" tag
	}

	return nil
}
