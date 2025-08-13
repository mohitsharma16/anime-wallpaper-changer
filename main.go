package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/manifoldco/promptui"
	"anime-wallpaper-changer/config"
	"anime-wallpaper-changer/utils"
	"anime-wallpaper-changer/wallpaper"
)

func main() {
	// Get the config directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Error getting user config dir: %v", err)
	}
	appConfigDir := filepath.Join(configDir, "anime-wallpaper-changer")
	if err := os.MkdirAll(appConfigDir, 0755); err != nil {
		log.Fatalf("Error creating app config dir: %v", err)
	}

	// Load the configuration
	configFile := filepath.Join(appConfigDir, "config.json")
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// If the config file doesn't exist, run the interactive setup
	if _, err := os.Stat(configFile); os.IsNotExist(err) {
		fmt.Println("First time setup...")

		categoriesPrompt := promptui.Select{
			Label: "Choose your favorite categories",
			Items: []string{"General", "Anime", "People"},
		}
		_, categories, err := categoriesPrompt.Run()
		if err != nil {
			log.Fatalf("Error with interactive setup: %v", err)
		}

		purityPrompt := promptui.Select{
			Label: "Choose the purity level",
			Items: []string{"SFW", "Sketchy"},
		}
		_, purity, err := purityPrompt.Run()
		if err != nil {
			log.Fatalf("Error with interactive setup: %v", err)
		}

		// Convert answers to config format
		switch categories {
		case "General":
			cfg.Categories = []string{"100"}
		case "Anime":
			cfg.Categories = []string{"010"}
		case "People":
			cfg.Categories = []string{"001"}
		}

		switch purity {
		case "SFW":
			cfg.Purity = []string{"100"}
		case "Sketchy":
			cfg.Purity = []string{"010"}
		}

		autostartPrompt := promptui.Select{
			Label: "Run on login?",
			Items: []string{"Yes", "No"},
		}
		_, autostart, err := autostartPrompt.Run()
		if err != nil {
			log.Fatalf("Error with interactive setup: %v", err)
		}

		if autostart == "Yes" {
			if err := utils.SetupAutostart(); err != nil {
				log.Printf("Error setting up autostart: %v", err)
			}
		}

		if err := config.SaveConfig(configFile, cfg); err != nil {
			log.Fatalf("Error saving config: %v", err)
		}
	}

	// Get a random wallpaper
	wallpaperURL, err := wallpaper.GetRandomWallpaper(cfg)
	if err != nil {
		log.Fatalf("Error getting wallpaper: %v", err)
	}

	// Download the wallpaper
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatalf("Error getting user cache dir: %v", err)
	}
	appCacheDir := filepath.Join(cacheDir, "anime-wallpaper-changer")
	if err := os.MkdirAll(appCacheDir, 0755); err != nil {
		log.Fatalf("Error creating app cache dir: %v", err)
	}
	wpaperFile := filepath.Join(appCacheDir, filepath.Base(wallpaperURL))
	if err := wallpaper.DownloadWallpaper(wallpaperURL, wpaperFile); err != nil {
		log.Fatalf("Error downloading wallpaper: %v", err)
	}

	// Set the wallpaper
	if err := utils.SetWallpaper(wpaperFile); err != nil {
		log.Fatalf("Error setting wallpaper: %v", err)
	}

	fmt.Println("Wallpaper set successfully!")
}