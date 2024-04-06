package main

import (
    "fmt"
    "github.com/charmbracelet/huh"
    "net/url"
    "regexp"
)

var (
    DefaultHost            = "https://example.com"
    FrontendVirtualDirectory string
    DatabaseConnectionString string
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

    // Now you can use the exported variables:
    // DefaultHost, FrontendVirtualDirectory, and DatabaseConnectionString
    fmt.Printf("Default Host: %s\n", DefaultHost)
    fmt.Printf("Frontend Virtual Directory: %s\n", FrontendVirtualDirectory)
    fmt.Printf("Database Connection String: %s\n", DatabaseConnectionString)
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

