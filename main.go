package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"wallsync/config"
	"wallsync/utils"
	"wallsync/wallpaper"
	"github.com/manifoldco/promptui"
)

var (
	daemonMode     = flag.Bool("daemon", false, "Run in daemon mode (background service)")
	onceMode       = flag.Bool("once", false, "Change wallpaper once and exit")
	reconfigMode   = flag.Bool("reconfig", false, "Reconfigure preferences")
	intervalFlag   = flag.Int("interval", 0, "Change interval in minutes (0 = use config value)")
)

func main() {
	flag.Parse()

	// Get the config directory
	configDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Error getting user config dir: %v", err)
	}
	appConfigDir := filepath.Join(configDir, "wallsync")

	if err := os.MkdirAll(appConfigDir, 0755); err != nil {
		log.Fatalf("Error creating app config dir: %v", err)
	}

	// Load the configuration
	configFile := filepath.Join(appConfigDir, "config.json")
	cfg, err := config.LoadConfig(configFile)
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// If the config file doesn't exist or reconfigure mode, run the interactive setup
	if _, err := os.Stat(configFile); os.IsNotExist(err) || *reconfigMode {
		if *reconfigMode {
			fmt.Println("Reconfiguring WallSync...")
		} else {
			fmt.Println("First time setup...")
		}

		cfg = runInteractiveSetup(cfg)

		if err := config.SaveConfig(configFile, cfg); err != nil {
			log.Fatalf("Error saving config: %v", err)
		}
		fmt.Println("Configuration saved!")
	}

	// Set default interval if not configured (for old config files)
	if cfg.ChangeInterval <= 0 {
		cfg.ChangeInterval = 60 // Default to 60 minutes
	}

	// Override interval if flag is provided
	if *intervalFlag > 0 {
		cfg.ChangeInterval = *intervalFlag
	}

	// Setup autostart based on config
	if cfg.Autostart {
		if err := utils.SetupAutostart(); err != nil {
			log.Printf("Warning: Error setting up autostart: %v", err)
			fmt.Println("\nNote: On Windows, creating autostart requires administrator privileges.")
			fmt.Println("To enable autostart: Right-click wallsync.exe → 'Run as administrator' → Run setup again")
			fmt.Println("Or run manually with: wallsync.exe -daemon")
			fmt.Println()
		} else {
			fmt.Println("✓ Autostart configured successfully")
		}
	} else {
		// If autostart is not enabled, ensure the task is removed
		if err := utils.RemoveAutostart(); err != nil {
			// Ignore error if task doesn't exist
		}
	}

	// Determine run mode
	if *onceMode {
		// Run once and exit
		changeWallpaper(cfg)
	} else if *daemonMode || cfg.Autostart {
		// Run in daemon mode
		runDaemon(cfg)
	} else {
		// Default: change wallpaper once
		changeWallpaper(cfg)
	}
}

func runInteractiveSetup(cfg *config.Config) *config.Config {
	// Reset categories for fresh setup
	cfg.Categories = []string{}

	// Multiple category selection
	categoryOptions := []string{"General", "Anime", "People"}
	fmt.Println("\nSelect categories (space to select, enter to confirm):")

	selectedCategories := make(map[string]bool)
	for i, cat := range categoryOptions {
		prompt := promptui.Select{
			Label: fmt.Sprintf("Include %s?", cat),
			Items: []string{"Yes", "No"},
		}
		_, result, err := prompt.Run()
		if err != nil {
			log.Fatalf("Error with interactive setup: %v", err)
		}
		if result == "Yes" {
			selectedCategories[cat] = true

			// Convert to config format
			switch cat {
			case "General":
				cfg.Categories = append(cfg.Categories, "100")
			case "Anime":
				cfg.Categories = append(cfg.Categories, "010")
			case "People":
				cfg.Categories = append(cfg.Categories, "001")
			}
		}
		_ = i
	}

	// If no categories selected, default to Anime
	if len(cfg.Categories) == 0 {
		cfg.Categories = []string{"010"}
	}

	purityPrompt := promptui.Select{
		Label: "Choose the purity level",
		Items: []string{"SFW", "Sketchy", "Both"},
	}
	_, purity, err := purityPrompt.Run()
	if err != nil {
		log.Fatalf("Error with interactive setup: %v", err)
	}

	switch purity {
	case "SFW":
		cfg.Purity = []string{"100"}
	case "Sketchy":
		cfg.Purity = []string{"010"}
	case "Both":
		cfg.Purity = []string{"100", "010"}
	}

	// Ask for change interval
	intervalPrompt := promptui.Select{
		Label: "How often should the wallpaper change?",
		Items: []string{"15 minutes", "30 minutes", "1 hour", "2 hours", "4 hours", "Daily"},
	}
	_, interval, err := intervalPrompt.Run()
	if err != nil {
		log.Fatalf("Error with interactive setup: %v", err)
	}

	switch interval {
	case "15 minutes":
		cfg.ChangeInterval = 15
	case "30 minutes":
		cfg.ChangeInterval = 30
	case "1 hour":
		cfg.ChangeInterval = 60
	case "2 hours":
		cfg.ChangeInterval = 120
	case "4 hours":
		cfg.ChangeInterval = 240
	case "Daily":
		cfg.ChangeInterval = 1440
	}

	autostartPrompt := promptui.Select{
		Label: "Run on login and change wallpapers automatically?",
		Items: []string{"Yes", "No"},
	}
	_, autostart, err := autostartPrompt.Run()
	if err != nil {
		log.Fatalf("Error with interactive setup: %v", err)
	}

	cfg.Autostart = (autostart == "Yes")

	notificationPrompt := promptui.Select{
		Label: "Show notifications when wallpaper changes?",
		Items: []string{"Yes", "No"},
	}
	_, notification, err := notificationPrompt.Run()
	if err != nil {
		log.Fatalf("Error with interactive setup: %v", err)
	}

	cfg.ShowNotification = (notification == "Yes")

	return cfg
}

func runDaemon(cfg *config.Config) {
	fmt.Println("WallSync daemon started")
	fmt.Printf("Wallpaper will change every %d minutes\n", cfg.ChangeInterval)
	fmt.Println("Press Ctrl+C to stop")

	// Setup signal handling for graceful shutdown
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	// Change wallpaper immediately on start
	changeWallpaper(cfg)

	// Create ticker for periodic wallpaper changes
	ticker := time.NewTicker(time.Duration(cfg.ChangeInterval) * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			changeWallpaper(cfg)
		case <-sigChan:
			fmt.Println("\nShutting down WallSync daemon...")
			return
		}
	}
}

func changeWallpaper(cfg *config.Config) {
	fmt.Println("Changing wallpaper...")

	// Get a random wallpaper
	wallpaperURL, err := wallpaper.GetRandomWallpaper(cfg)
	if err != nil {
		log.Printf("Error getting wallpaper: %v", err)
		utils.ShowNotification("WallSync Error", "Failed to fetch wallpaper. Will retry later.")
		return
	}

	// Download the wallpaper
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Printf("Error getting user cache dir: %v", err)
		return
	}
	appCacheDir := filepath.Join(cacheDir, "wallsync")
	if err := os.MkdirAll(appCacheDir, 0755); err != nil {
		log.Printf("Error creating app cache dir: %v", err)
		return
	}

	// Generate unique filename with timestamp
	filename := fmt.Sprintf("wallpaper_%d%s", time.Now().Unix(), filepath.Ext(wallpaperURL))
	wpaperFile := filepath.Join(appCacheDir, filename)

	if err := wallpaper.DownloadWallpaper(wallpaperURL, wpaperFile); err != nil {
		log.Printf("Error downloading wallpaper: %v", err)
		utils.ShowNotification("WallSync Error", "Failed to download wallpaper. Will retry later.")
		return
	}

	// Set the wallpaper
	if err := utils.SetWallpaper(wpaperFile); err != nil {
		log.Printf("Error setting wallpaper: %v", err)
		utils.ShowNotification("WallSync Error", "Failed to set wallpaper.")
		return
	}

	fmt.Printf("Wallpaper changed successfully! (%s)\n", time.Now().Format("15:04:05"))

	if cfg.ShowNotification {
		utils.ShowNotification("WallSync", "Wallpaper changed successfully!")
	}

	// Clean up old wallpapers (keep only last 5)
	cleanupOldWallpapers(appCacheDir, 5)
}

func cleanupOldWallpapers(cacheDir string, keepCount int) {
	entries, err := os.ReadDir(cacheDir)
	if err != nil {
		return
	}

	// Filter wallpaper files
	var wallpapers []os.DirEntry
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasPrefix(entry.Name(), "wallpaper_") {
			wallpapers = append(wallpapers, entry)
		}
	}

	// If we have more than keepCount wallpapers, delete the oldest ones
	if len(wallpapers) > keepCount {
		// Sort by modification time (oldest first)
		for i := 0; i < len(wallpapers)-keepCount; i++ {
			os.Remove(filepath.Join(cacheDir, wallpapers[i].Name()))
		}
	}
}
