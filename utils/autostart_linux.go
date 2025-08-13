//go:build linux

package utils

// SetupAutostart creates a .desktop file in ~/.config/autostart.
func SetupAutostart() error {
	return ErrUnsupportedOS
}
