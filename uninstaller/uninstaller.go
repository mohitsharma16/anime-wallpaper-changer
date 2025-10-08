package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"wallsync/utils"
)

func main() {
	fmt.Println("WallSync Uninstaller")
	fmt.Println("====================")
	fmt.Println()
	fmt.Println("This will remove:")
	fmt.Println("- Autostart configuration")
	fmt.Println("- Configuration files")
	fmt.Println("- Cached wallpapers")
	fmt.Println()
	fmt.Print("Continue? (y/n): ")

	var response string
	fmt.Scanln(&response)

	if response != "y" && response != "Y" && response != "yes" && response != "Yes" {
		fmt.Println("Uninstallation cancelled.")
		return
	}

	fmt.Println()
	fmt.Println("Starting uninstallation...")
	fmt.Println()

	// Remove autostart
	if err := utils.RemoveAutostart(); err != nil {
		log.Printf("Warning: Could not remove autostart: %v", err)
	} else {
		fmt.Println("✓ Autostart entry removed")
	}

	// Remove configuration directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Printf("Warning: Error getting config directory: %v", err)
	} else {
		appConfigDir := filepath.Join(configDir, "wallsync")
		if err := os.RemoveAll(appConfigDir); err != nil {
			log.Printf("Warning: Could not remove config directory: %v", err)
		} else {
			fmt.Printf("✓ Configuration removed (%s)\n", appConfigDir)
		}
	}

	// Remove cache directory
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Printf("Warning: Error getting cache directory: %v", err)
	} else {
		appCacheDir := filepath.Join(cacheDir, "wallsync")
		if err := os.RemoveAll(appCacheDir); err != nil {
			log.Printf("Warning: Could not remove cache directory: %v", err)
		} else {
			fmt.Printf("✓ Cache removed (%s)\n", appCacheDir)
		}
	}

	fmt.Println()
	fmt.Println("Uninstallation complete!")
	fmt.Println()
	fmt.Println("Note: You can now manually delete the application folder/executables.")
	fmt.Println()
	fmt.Print("Press Enter to exit...")
	fmt.Scanln()
}
