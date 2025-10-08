# Build and Setup Instructions

This guide provides detailed instructions for building and setting up WallSync on Windows, macOS, and Linux.

## Prerequisites

- [Go](https://golang.org/dl/) 1.16 or higher
- Git (optional, for cloning the repository)

## Building from Source

### Windows

1. **Open Command Prompt or PowerShell**

2. **Navigate to the project directory:**
   ```cmd
   cd D:\Projects\wallsync
   ```

3. **Build the application:**
   ```cmd
   go build -o wallsync.exe main.go
   go build -o uninstall.exe uninstaller\uninstaller.go
   ```

4. **The executables will be created in the current directory:**
   - `wallsync.exe` - Main application
   - `uninstall.exe` - Uninstaller

### macOS

1. **Open Terminal**

2. **Navigate to the project directory:**
   ```bash
   cd /path/to/wallsync
   ```

3. **Build the application:**

   **For Intel Macs:**
   ```bash
   GOOS=darwin GOARCH=amd64 go build -o wallsync main.go
   GOOS=darwin GOARCH=amd64 go build -o uninstaller uninstaller/uninstaller.go
   ```

   **For Apple Silicon (M1/M2/M3):**
   ```bash
   GOOS=darwin GOARCH=arm64 go build -o wallsync main.go
   GOOS=darwin GOARCH=arm64 go build -o uninstaller uninstaller/uninstaller.go
   ```

   **Or build for your current architecture:**
   ```bash
   go build -o wallsync main.go
   go build -o uninstaller uninstaller/uninstaller.go
   ```

4. **Make the binaries executable:**
   ```bash
   chmod +x wallsync uninstaller
   ```

5. **Optionally, move to a system path:**
   ```bash
   sudo mv wallsync /usr/local/bin/
   sudo mv uninstaller /usr/local/bin/wallsync-uninstaller
   ```

### Linux

1. **Open Terminal**

2. **Navigate to the project directory:**
   ```bash
   cd /path/to/wallsync
   ```

3. **Build the application:**
   ```bash
   go build -o wallsync main.go
   go build -o uninstaller uninstaller/uninstaller.go
   ```

4. **Make the binaries executable:**
   ```bash
   chmod +x wallsync uninstaller
   ```

5. **Optionally, move to a system path:**
   ```bash
   sudo mv wallsync /usr/local/bin/
   sudo mv uninstaller /usr/local/bin/wallsync-uninstaller
   ```

## Cross-Compilation

You can build for different platforms from any OS:

### From Windows

**Build for macOS (Intel):**
```cmd
set GOOS=darwin
set GOARCH=amd64
go build -o wallsync main.go
go build -o uninstaller uninstaller/uninstaller.go
```

**Build for macOS (Apple Silicon):**
```cmd
set GOOS=darwin
set GOARCH=arm64
go build -o wallsync main.go
go build -o uninstaller uninstaller/uninstaller.go
```

**Build for Linux:**
```cmd
set GOOS=linux
set GOARCH=amd64
go build -o wallsync main.go
go build -o uninstaller uninstaller/uninstaller.go
```

### From macOS/Linux

**Build for Windows:**
```bash
GOOS=windows GOARCH=amd64 go build -o wallsync.exe main.go
GOOS=windows GOARCH=amd64 go build -o uninstall.exe uninstaller/uninstaller.go
```

**Build for macOS (Intel):**
```bash
GOOS=darwin GOARCH=amd64 go build -o wallsync main.go
GOOS=darwin GOARCH=amd64 go build -o uninstaller uninstaller/uninstaller.go
```

**Build for macOS (Apple Silicon):**
```bash
GOOS=darwin GOARCH=arm64 go build -o wallsync main.go
GOOS=darwin GOARCH=arm64 go build -o uninstaller uninstaller/uninstaller.go
```

**Build for Linux:**
```bash
GOOS=linux GOARCH=amd64 go build -o wallsync main.go
GOOS=linux GOARCH=amd64 go build -o uninstaller uninstaller/uninstaller.go
```

## First-Time Setup

After building, run the application for the first time to configure your preferences:

### Windows
```cmd
wallsync.exe
```

### macOS/Linux
```bash
./wallsync
```

You'll be prompted to configure:
- **Wallpaper categories** - Choose from General, Anime, and People (can select multiple)
- **Purity level** - SFW, Sketchy, or Both
- **Change interval** - 15 minutes, 30 minutes, 1 hour, 2 hours, 4 hours, or Daily
- **Auto-start on login** - Yes or No
- **Desktop notifications** - Yes or No

## Platform-Specific Setup

### Windows

**Autostart Configuration:**
- When enabled, WallSync creates a scheduled task named "WallSync"
- The task runs at logon using a VBScript wrapper for silent execution
- VBScript location: `%TEMP%\wallsync_autostart.vbs`

**Configuration Location:**
```
%APPDATA%\wallsync\config.json
```

**Cache Location:**
```
%LOCALAPPDATA%\wallsync\
```

**To manually check scheduled task:**
```cmd
schtasks /query /tn WallSync
```

### macOS

**Autostart Configuration:**
- Creates a LaunchAgent plist file
- Location: `~/Library/LaunchAgents/com.wallsync.daemon.plist`
- Logs: `~/Library/Logs/wallsync.log` and `~/Library/Logs/wallsync-error.log`

**Configuration Location:**
```
~/Library/Application Support/wallsync/config.json
```

**Cache Location:**
```
~/Library/Caches/wallsync/
```

**To manually manage the LaunchAgent:**
```bash
# Check status
launchctl list | grep wallsync

# Start manually
launchctl load ~/Library/LaunchAgents/com.wallsync.daemon.plist

# Stop
launchctl unload ~/Library/LaunchAgents/com.wallsync.daemon.plist
```

### Linux

**Autostart Configuration:**
- Creates a .desktop file
- Location: `~/.config/autostart/wallsync.desktop`

**Configuration Location:**
```
~/.config/wallsync/config.json
```

**Cache Location:**
```
~/.cache/wallsync/
```

**Desktop Environment Support:**
- **GNOME** - Uses `gsettings` (usually pre-installed)
- **KDE Plasma** - Uses `qdbus` (usually pre-installed)
- **XFCE** - Uses `xfconf-query` (usually pre-installed)

**Notification Requirements:**
- Requires `notify-send` (usually pre-installed with most desktop environments)
- If missing, install via:
  ```bash
  # Ubuntu/Debian
  sudo apt install libnotify-bin

  # Fedora
  sudo dnf install libnotify

  # Arch
  sudo pacman -S libnotify
  ```

## Usage Examples

### Run Once (Change Wallpaper Immediately)
```bash
# Windows
wallsync.exe -once

# macOS/Linux
./wallsync -once
```

### Run in Daemon Mode (Background Service)
```bash
# Windows
wallsync.exe -daemon

# macOS/Linux
./wallsync -daemon
```

### Reconfigure Preferences
```bash
# Windows
wallsync.exe -reconfig

# macOS/Linux
./wallsync -reconfig
```

### Override Change Interval
```bash
# Change wallpaper every 30 minutes
# Windows
wallsync.exe -daemon -interval 30

# macOS/Linux
./wallsync -daemon -interval 30
```

## Uninstallation

### Windows
```cmd
uninstall.exe
```

### macOS/Linux
```bash
./uninstaller
```

The uninstaller will:
- Remove autostart configuration
- Remove configuration files
- Remove cached wallpapers
- Clean up all WallSync-related files

## Troubleshooting

### Windows

**Issue: Scheduled task not running**
```cmd
# Delete and recreate the task
schtasks /delete /tn WallSync /f
wallsync.exe -reconfig
```

**Issue: Wallpaper not changing**
- Check if the scheduled task is enabled
- Run manually to test: `wallsync.exe -once`
- Check Windows Event Viewer for errors

### macOS

**Issue: LaunchAgent not running**
```bash
# Check LaunchAgent status
launchctl list | grep wallsync

# Reload LaunchAgent
launchctl unload ~/Library/LaunchAgents/com.wallsync.daemon.plist
launchctl load ~/Library/LaunchAgents/com.wallsync.daemon.plist

# Check logs
tail -f ~/Library/Logs/wallsync.log
```

**Issue: Notifications not showing**
- Check System Preferences → Notifications → Script Editor (notifications are sent via osascript)

### Linux

**Issue: Wallpaper not changing on KDE**
```bash
# Manually test KDE wallpaper command
qdbus org.kde.plasmashell /PlasmaShell org.kde.PlasmaShell.evaluateScript 'print("test")'
```

**Issue: Notifications not showing**
```bash
# Test notify-send
notify-send "Test" "This is a test notification"

# If not installed, install libnotify
sudo apt install libnotify-bin  # Ubuntu/Debian
```

**Issue: Autostart not working**
```bash
# Check if desktop file exists
cat ~/.config/autostart/wallsync.desktop

# Verify executable path is correct
which wallsync
```

## Development

### Running without Building
```bash
go run main.go
```

### Running Tests (if available)
```bash
go test ./...
```

### Installing Dependencies
```bash
go mod download
```

### Updating Dependencies
```bash
go get -u ./...
go mod tidy
```

## Additional Notes

- **Configuration File Format:** JSON
- **Supported Image Formats:** JPG, PNG (from Wallhaven API)
- **Minimum Resolution:** 1920x1080 (configurable in code)
- **Default Ratio:** 16:9 (configurable in code)
- **Cache Management:** Automatically keeps only the last 5 wallpapers
- **Network Timeout:** 30s for API calls, 60s for image downloads

## Contributing

See the main [README.md](README.md) for contribution guidelines.
