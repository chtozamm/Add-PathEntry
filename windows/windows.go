package windows

import (
	"fmt"
	"log"
	"strings"
)

// UpdatePath adds a new directory to the current user PATH environment variable.
// If the directory is already present in the PATH, it will not be added again.
func UpdatePath(newPath string) {
	// Retrieve the user PATH
	userPath, err := getUserPath()
	if err != nil {
		log.Fatalf("Error retrieving PATH: %v", err)
	}

	// Split the current PATH into individual directories
	paths := strings.Split(userPath, ";")
	// Check if the newPath is already in the PATH
	for _, path := range paths {
		if path == newPath {
			fmt.Printf("\"%s\" is already in the PATH.", newPath)
			return
		}
	}

	// Prepare the new user PATH by appending the newPath
	newUserPath := userPath
	if userPath != "" {
		newUserPath += ";"
	}
	newUserPath += newPath

	// Set the updated user PATH
	err = setUserPath(newUserPath)
	if err != nil {
		log.Fatalf("Error setting user PATH: %v", err)
	}

	fmt.Printf("Successfully added \"%s\" to user PATH.\n", newPath)
}
