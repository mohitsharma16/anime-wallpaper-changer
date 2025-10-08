//go:build windows
// +build windows

package utils

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
)

const taskName = "WallSync"

// SetupAutostart creates a new task in the Task Scheduler to run the application on logon.
func SetupAutostart() error {
	ex, err := os.Executable()
	if err != nil {
		return fmt.Errorf("failed to get executable path: %w", err)
	}
	appPath := filepath.Dir(ex)
	appName := filepath.Base(ex)

	// Create a VBScript to run the application silently in daemon mode
	fullAppPath := filepath.Join(appPath, appName)
	vbsContent := fmt.Sprintf(`Set WshShell = CreateObject("WScript.Shell")
WshShell.Run "%s -daemon", 0, false`, fullAppPath)

	vbsPath := filepath.Join(os.TempDir(), "wallsync_autostart.vbs")
	err = os.WriteFile(vbsPath, []byte(vbsContent), 0644)
	if err != nil {
		return fmt.Errorf("failed to write VBScript: %w", err)
	}

	// Register the VBScript with Task Scheduler
	cmd := exec.Command("schtasks", "/create", "/tn", taskName, "/tr", fmt.Sprintf("wscript.exe \"%s\"", vbsPath), "/sc", "onlogon", "/f")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create scheduled task: %w\nOutput: %s", err, output)
	}

	return nil
}

// RemoveAutostart removes the scheduled task.
func RemoveAutostart() error {
	cmd := exec.Command("schtasks", "/delete", "/tn", taskName, "/f")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to delete scheduled task: %w\nOutput: %s", err, output)
	}

	// Clean up the VBScript file
	vbsPath := filepath.Join(os.TempDir(), "wallsync_autostart.vbs")
	os.Remove(vbsPath)

	return nil
}
