//go:build darwin

package utils

import (
	"os/exec"
)

// SetWallpaper sets the desktop wallpaper on macOS.
func SetWallpaper(path string) error {
	cmd := exec.Command("osascript", "-e", `tell application "Finder" to set desktop picture to POSIX file "`+path+`"`)
	return cmd.Run()
}

// ShowNotification displays a notification on macOS using osascript.
func ShowNotification(title, message string) {
	script := `display notification "` + message + `" with title "` + title + `"`
	cmd := exec.Command("osascript", "-e", script)
	cmd.Run()
}