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