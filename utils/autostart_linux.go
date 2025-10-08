//go:build linux

package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

// SetupAutostart creates a .desktop file in ~/.config/autostart.
func SetupAutostart() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	ex, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	autostartDir := filepath.Join(homeDir, ".config", "autostart")
	if err := os.MkdirAll(autostartDir, 0755); err != nil {
		return fmt.Errorf("failed to create autostart directory: %w", err)
	}

	desktopFilePath := filepath.Join(autostartDir, "wallsync.desktop")

	desktopContent := fmt.Sprintf(`[Desktop Entry]
Type=Application
Name=WallSync
Comment=Automatic wallpaper changer
Exec=%s -daemon
Terminal=false
Hidden=false
X-GNOME-Autostart-enabled=true
`, ex)

	if err := os.WriteFile(desktopFilePath, []byte(desktopContent), 0644); err != nil {
		return fmt.Errorf("failed to write desktop file: %w", err)
	}

	return nil
}

// RemoveAutostart removes the .desktop file.
func RemoveAutostart() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	desktopFilePath := filepath.Join(homeDir, ".config", "autostart", "wallsync.desktop")
	os.Remove(desktopFilePath)

	return nil
}
