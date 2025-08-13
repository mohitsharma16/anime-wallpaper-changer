//go:build darwin

package utils

// SetupAutostart creates a LaunchAgent .plist file in ~/Library/LaunchAgents.
func SetupAutostart() error {
	return ErrUnsupportedOS
}
