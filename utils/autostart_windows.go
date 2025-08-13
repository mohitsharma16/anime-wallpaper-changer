//go:build windows

package utils

import (
	"os"
	"os/exec"
)

// SetupAutostart creates a new task in the Task Scheduler to run the application on logon.
func SetupAutostart() error {
	exePath, err := os.Executable()
	if err != nil {
		return err
	}

	cmd := exec.Command("schtasks", "/create", "/tn", "Anime Wallpaper Changer", "/tr", exePath, "/sc", "onlogon", "/rl", "highest", "/f")
	return cmd.Run()
}
