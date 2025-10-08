//go:build darwin

package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

// SetupAutostart creates a LaunchAgent .plist file in ~/Library/LaunchAgents.
func SetupAutostart() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	ex, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}

	launchAgentsDir := filepath.Join(homeDir, "Library", "LaunchAgents")
	if err := os.MkdirAll(launchAgentsDir, 0755); err != nil {
		return fmt.Errorf("failed to create LaunchAgents directory: %w", err)
	}

	plistPath := filepath.Join(launchAgentsDir, "com.wallsync.daemon.plist")

	plistContent := fmt.Sprintf(`<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
	<key>Label</key>
	<string>com.wallsync.daemon</string>
	<key>ProgramArguments</key>
	<array>
		<string>%s</string>
		<string>-daemon</string>
	</array>
	<key>RunAtLoad</key>
	<true/>
	<key>KeepAlive</key>
	<false/>
	<key>StandardOutPath</key>
	<string>%s/Library/Logs/wallsync.log</string>
	<key>StandardErrorPath</key>
	<string>%s/Library/Logs/wallsync-error.log</string>
</dict>
</plist>`, ex, homeDir, homeDir)

	if err := os.WriteFile(plistPath, []byte(plistContent), 0644); err != nil {
		return fmt.Errorf("failed to write plist file: %w", err)
	}

	// Load the launch agent
	cmd := exec.Command("launchctl", "load", plistPath)
	if err := cmd.Run(); err != nil {
		return fmt.Errorf("failed to load launch agent: %w", err)
	}

	return nil
}

// RemoveAutostart removes the LaunchAgent.
func RemoveAutostart() error {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return fmt.Errorf("failed to get home directory: %w", err)
	}

	plistPath := filepath.Join(homeDir, "Library", "LaunchAgents", "com.wallsync.daemon.plist")

	// Unload the launch agent
	cmd := exec.Command("launchctl", "unload", plistPath)
	cmd.Run() // Ignore error if not loaded

	// Remove the plist file
	os.Remove(plistPath)

	return nil
}
