//go:build linux

package utils

import (
	"os/exec"
)

// SetWallpaper sets the desktop wallpaper on Linux.
func SetWallpaper(path string) error {
	// GNOME
	cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", "file://"+path)
	if err := cmd.Run(); err == nil {
		return nil
	}

	// KDE
	// TODO: Implement KDE support

	// XFCE
	// TODO: Implement XFCE support

	return ErrUnsupportedOS
}