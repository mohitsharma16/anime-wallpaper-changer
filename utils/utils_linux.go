//go:build linux

package utils

import (
	"os/exec"
)

// SetWallpaper sets the desktop wallpaper on Linux.
func SetWallpaper(path string) error {
	// Try GNOME first
	cmd := exec.Command("gsettings", "set", "org.gnome.desktop.background", "picture-uri", "file://"+path)
	if err := cmd.Run(); err == nil {
		return nil
	}

	// Try KDE Plasma
	script := `
		qdbus org.kde.plasmashell /PlasmaShell org.kde.PlasmaShell.evaluateScript '
			var allDesktops = desktops();
			for (i=0;i<allDesktops.length;i++) {
				d = allDesktops[i];
				d.wallpaperPlugin = "org.kde.image";
				d.currentConfigGroup = Array("Wallpaper", "org.kde.image", "General");
				d.writeConfig("Image", "file://` + path + `");
			}
		'
	`
	cmd = exec.Command("sh", "-c", script)
	if err := cmd.Run(); err == nil {
		return nil
	}

	// Try XFCE
	cmd = exec.Command("xfconf-query", "-c", "xfce4-desktop", "-p", "/backdrop/screen0/monitor0/workspace0/last-image", "-s", path)
	if err := cmd.Run(); err == nil {
		return nil
	}

	return ErrUnsupportedOS
}

// ShowNotification displays a notification on Linux using notify-send.
func ShowNotification(title, message string) {
	cmd := exec.Command("notify-send", title, message)
	cmd.Run()
}