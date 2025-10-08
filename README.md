# WallSync

A lightweight background application that automatically changes your desktop wallpaper to random high-quality wallpapers at configurable intervals.

## Features

-   **Automatic wallpaper changes** - Set custom intervals (15 min to daily)
-   **Background daemon mode** - Runs silently in the background
-   **Multiple category selection** - Choose from General, Anime, and People wallpapers
-   **Desktop notifications** - Get notified when wallpaper changes
-   **Auto-start on login** - Optionally launch on system startup
-   **Smart caching** - Automatic cleanup of old wallpapers
-   **Timeout handling** - Robust network error recovery
-   **Cross-platform** - Windows, macOS, and Linux support
-   Fetches high-quality wallpapers from [Wallhaven](https://wallhaven.cc)

## Installation

**Quick Build:**

**Windows:**
```bash
go build -o wallsync.exe main.go
go build -o uninstall.exe uninstaller/uninstaller.go
```

**macOS/Linux:**
```bash
go build -o wallsync main.go
go build -o uninstaller uninstaller/uninstaller.go
```

**For detailed build instructions, cross-compilation, and platform-specific setup, see [BUILD.md](BUILD.md)**

## Usage

### First-Time Setup

Run the application for the first time to configure your preferences:

```bash
# Windows
./wallsync.exe

# macOS/Linux
./wallsync
```

You'll be prompted to:
-   Select wallpaper categories (General, Anime, People)
-   Choose purity level (SFW, Sketchy, or Both)
-   Set change interval (15 min, 30 min, 1 hour, 2 hours, 4 hours, or Daily)
-   Enable/disable auto-start on login
-   Enable/disable desktop notifications

### Command-Line Options

```bash
# Run in daemon mode (background service)
./wallsync -daemon

# Change wallpaper once and exit
./wallsync -once

# Reconfigure preferences
./wallsync -reconfig

# Override change interval (in minutes)
./wallsync -interval 30
```

### Uninstall

Run the uninstaller to remove WallSync:

```bash
# Windows
./uninstall.exe

# macOS/Linux
./uninstaller
```

## Supported Platforms

### ✅ Fully Supported

-   **Windows** - All features working
-   **macOS** - All features working (notifications via osascript, autostart via LaunchAgent)
-   **Linux** - All features working with GNOME, KDE Plasma, and XFCE
    - Notifications via `notify-send`
    - Autostart via `.desktop` file
    - Wallpaper changes on GNOME (gsettings), KDE (qdbus), and XFCE (xfconf-query)

### Platform-Specific Notes

**macOS:**
- Notifications use native macOS notification center
- Autostart creates LaunchAgent at `~/Library/LaunchAgents/com.wallsync.daemon.plist`
- Logs available at `~/Library/Logs/wallsync.log`

**Linux:**
- Requires `notify-send` for notifications (usually pre-installed)
- Desktop file created at `~/.config/autostart/wallsync.desktop`
- Automatically detects and supports GNOME, KDE Plasma, and XFCE desktop environments

## Contributing

Contributions are welcome! If you want to add support for other desktop environments on Linux or improve the application in any other way, feel free to open a pull request.
