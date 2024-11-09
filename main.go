package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/chtozamm/Add-PathEntry/windows"
)

func main() {
	var targetDir string

	// Check for command-line arguments
	if len(os.Args) > 2 {
		printHelp()
		return
	}

	// If a single argument is provided, check if it's a help command
	if len(os.Args) == 2 {
		if os.Args[1] == "/?" || os.Args[1] == "--help" || os.Args[1] == "-h" {
			printHelp()
			return
		}
		// Set targetDir to the provided argument
		targetDir = os.Args[1]
	} else {
		// If no argument is provided, get the current working directory
		var err error
		targetDir, err = os.Getwd()
		if err != nil {
			log.Fatalf("Error getting current working directory: %v", err)
		}
	}

	// Normalize the target directory to an absolute path
	absoluteTargetDir, err := filepath.Abs(targetDir)
	if err != nil {
		log.Fatalf("Error normalizing target directory: %v", err)
	}

	// Update the user PATH based on the operating system
	switch runtime.GOOS {
	case "windows":
		windows.UpdatePath(absoluteTargetDir)
	default:
		log.Fatalf("Unsupported operating system: %s", runtime.GOOS)
	}
}

func printHelp() {
	fmt.Println("Usage: Add-PathEntry [path]")
	fmt.Println("If no path is provided, the current working directory will be added to the PATH.")
}
