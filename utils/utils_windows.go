//go:build windows

package utils

import (
	"os/exec"
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

// ShowNotification displays a notification on Windows using PowerShell.
func ShowNotification(title, message string) {
	// Use PowerShell to show a toast notification
	script := `
		[Windows.UI.Notifications.ToastNotificationManager, Windows.UI.Notifications, ContentType = WindowsRuntime] | Out-Null
		[Windows.Data.Xml.Dom.XmlDocument, Windows.Data.Xml.Dom.XmlDocument, ContentType = WindowsRuntime] | Out-Null

		$template = @"
<toast>
	<visual>
		<binding template="ToastText02">
			<text id="1">` + title + `</text>
			<text id="2">` + message + `</text>
		</binding>
	</visual>
</toast>
"@

		$xml = New-Object Windows.Data.Xml.Dom.XmlDocument
		$xml.LoadXml($template)
		$toast = New-Object Windows.UI.Notifications.ToastNotification $xml
		[Windows.UI.Notifications.ToastNotificationManager]::CreateToastNotifier("WallSync").Show($toast)
	`

	cmd := exec.Command("powershell", "-Command", script)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Run()
}
