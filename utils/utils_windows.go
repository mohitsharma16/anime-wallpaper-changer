//go:build windows

package utils

import (
	"syscall"
	"unsafe"
)

const (
	spiSetDeskWallpaper = 0x0014
	uiParamUpdateIniFile = 0x01
)

// SetWallpaper sets the desktop wallpaper on Windows.
func SetWallpaper(path string) error {
	pathPtr, err := syscall.UTF16PtrFromString(path)
	if err != nil {
		return err
	}

	_, _, err = syscall.NewLazyDLL("user32.dll").NewProc("SystemParametersInfoW").Call(
		spiSetDeskWallpaper,
		0,
		uintptr(unsafe.Pointer(pathPtr)),
		uiParamUpdateIniFile,
	)

    if err.Error() != "The operation completed successfully." {
        return err
    }

	return nil
}
