# Installation Guide for End Users

This guide helps you install and set up WallSync on your computer.

## 📥 Download

Go to the [Releases page](https://github.com/yourusername/wallsync/releases) and download the file for your operating system:

### Windows
- Download `wallsync-windows-amd64.zip`
- Works on Windows 10 and Windows 11

### macOS
- **Intel Mac:** Download `wallsync-darwin-amd64.tar.gz`
- **Apple Silicon (M1/M2/M3):** Download `wallsync-darwin-arm64.tar.gz`
- Works on macOS 10.10 (Yosemite) and later

### Linux
- **Most computers:** Download `wallsync-linux-amd64.tar.gz`
- **ARM devices:** Download `wallsync-linux-arm64.tar.gz`
- Works on GNOME, KDE Plasma, and XFCE desktop environments

---

## 🪟 Windows Installation

### Step 1: Extract Files
1. Right-click on `wallsync-windows-amd64.zip`
2. Select **"Extract All..."**
3. Choose a location (e.g., `C:\Program Files\WallSync`)
4. Click **"Extract"**

### Step 2: Run WallSync
1. Open the extracted folder
2. Double-click `wallsync-windows-amd64.exe`
3. If you see a Windows Security warning:
   - Click **"More info"**
   - Click **"Run anyway"**

### Step 3: Initial Setup
Follow the on-screen prompts:
1. **Categories:** Choose wallpaper types (General, Anime, People)
2. **Purity:** Choose SFW, Sketchy, or Both
3. **Interval:** How often to change wallpaper (15 min to daily)
4. **Auto-start:** Choose "Yes" to start automatically on login
5. **Notifications:** Choose "Yes" to see change notifications

### Step 4: Done!
- Your wallpaper will change immediately
- If you enabled auto-start, WallSync will run automatically when you log in
- The app runs silently in the background

---

## 🍎 macOS Installation

### Step 1: Extract Files
1. Double-click `wallsync-darwin-*.tar.gz` to extract
2. Or use Terminal:
   ```bash
   tar -xzf wallsync-darwin-*.tar.gz
   ```

### Step 2: Move to Applications (Optional)
```bash
# Create a folder for WallSync
mkdir -p ~/Applications/WallSync

# Move files there
mv wallsync-darwin-* ~/Applications/WallSync/wallsync
mv uninstaller-darwin-* ~/Applications/WallSync/uninstaller

# Make executable
chmod +x ~/Applications/WallSync/wallsync
chmod +x ~/Applications/WallSync/uninstaller
```

### Step 3: Run WallSync
1. Open Terminal
2. Navigate to the folder:
   ```bash
   cd ~/Applications/WallSync
   ```
3. Run WallSync:
   ```bash
   ./wallsync
   ```
4. If you see a security warning:
   - Open **System Preferences → Security & Privacy**
   - Click **"Open Anyway"** next to the WallSync message
   - Run `./wallsync` again

### Step 4: Initial Setup
Follow the on-screen prompts (same as Windows above)

### Step 5: Grant Permissions
macOS may ask for permissions:
- **Notifications:** Allow for wallpaper change alerts
- **Files and Folders:** Allow to download and cache wallpapers

### Step 6: Done!
- Your wallpaper will change immediately
- If you enabled auto-start, WallSync will run automatically on login
- Logs are saved to `~/Library/Logs/wallsync.log`

---

## 🐧 Linux Installation

### Step 1: Extract Files
```bash
# Extract archive
tar -xzf wallsync-linux-amd64.tar.gz

# Make files executable
chmod +x wallsync-linux-amd64
chmod +x uninstaller-linux-amd64
```

### Step 2: Move to System Path (Optional)
```bash
# For single user
mkdir -p ~/.local/bin
mv wallsync-linux-amd64 ~/.local/bin/wallsync
mv uninstaller-linux-amd64 ~/.local/bin/wallsync-uninstaller

# Add to PATH if not already (add to ~/.bashrc or ~/.zshrc)
export PATH="$HOME/.local/bin:$PATH"

# For all users (requires sudo)
sudo mv wallsync-linux-amd64 /usr/local/bin/wallsync
sudo mv uninstaller-linux-amd64 /usr/local/bin/wallsync-uninstaller
```

### Step 3: Install Dependencies (if needed)
```bash
# Ubuntu/Debian
sudo apt install libnotify-bin

# Fedora
sudo dnf install libnotify

# Arch Linux
sudo pacman -S libnotify
```

### Step 4: Run WallSync
```bash
./wallsync
# or if installed to system path:
wallsync
```

### Step 5: Initial Setup
Follow the on-screen prompts (same as Windows above)

### Step 6: Done!
- Your wallpaper will change immediately
- If you enabled auto-start, WallSync will start automatically on login
- Autostart file is created at `~/.config/autostart/wallsync.desktop`

---

## 🎯 Using WallSync

### Change Wallpaper Now
```bash
# Windows
wallsync-windows-amd64.exe -once

# macOS
./wallsync -once

# Linux
./wallsync -once
```

### Change Settings
```bash
# Windows
wallsync-windows-amd64.exe -reconfig

# macOS
./wallsync -reconfig

# Linux
./wallsync -reconfig
```

### Check if Running (Windows)
1. Open Task Manager (Ctrl+Shift+Esc)
2. Look for `wallsync-windows-amd64.exe` in processes

### Check if Running (macOS)
```bash
# Check LaunchAgent status
launchctl list | grep wallsync

# View logs
tail -f ~/Library/Logs/wallsync.log
```

### Check if Running (Linux)
```bash
# Check if process is running
ps aux | grep wallsync

# Check autostart file
cat ~/.config/autostart/wallsync.desktop
```

---

## 🗑️ Uninstalling WallSync

### Windows
1. Run `uninstaller-windows-amd64.exe`
2. Delete the WallSync folder

### macOS
```bash
./uninstaller
# Then delete the application folder
rm -rf ~/Applications/WallSync
```

### Linux
```bash
./uninstaller
# Then remove the binary
rm ~/.local/bin/wallsync ~/.local/bin/wallsync-uninstaller
# Or if installed system-wide:
sudo rm /usr/local/bin/wallsync /usr/local/bin/wallsync-uninstaller
```

The uninstaller removes:
- Auto-start configuration
- Cached wallpapers
- Configuration files

---

## ❓ Troubleshooting

### Windows

**Problem: Wallpaper doesn't change**
1. Check if scheduled task exists:
   - Open Task Scheduler
   - Look for "WallSync" task
2. Run manually: `wallsync-windows-amd64.exe -once`
3. Check if you have internet connection

**Problem: "Windows protected your PC" message**
- Click "More info" → "Run anyway"
- This is normal for unsigned applications

### macOS

**Problem: "Cannot be opened because the developer cannot be verified"**
1. Open System Preferences → Security & Privacy
2. Click "Open Anyway"
3. Run the app again

**Problem: Notifications not showing**
- Check System Preferences → Notifications → Script Editor
- Make sure notifications are enabled

**Problem: Wallpaper doesn't change**
```bash
# Check if LaunchAgent is loaded
launchctl list | grep wallsync

# View error logs
cat ~/Library/Logs/wallsync-error.log

# Reload LaunchAgent
launchctl unload ~/Library/LaunchAgents/com.wallsync.daemon.plist
launchctl load ~/Library/LaunchAgents/com.wallsync.daemon.plist
```

### Linux

**Problem: Wallpaper doesn't change**
1. Check your desktop environment:
   ```bash
   echo $XDG_CURRENT_DESKTOP
   ```
2. Verify you're running GNOME, KDE, or XFCE
3. Run manually: `./wallsync -once`

**Problem: Notifications not showing**
```bash
# Test notifications
notify-send "Test" "This is a test"

# Install if missing
sudo apt install libnotify-bin  # Ubuntu/Debian
```

**Problem: Autostart not working**
1. Check desktop file:
   ```bash
   cat ~/.config/autostart/wallsync.desktop
   ```
2. Verify executable path is correct
3. Try running manually first

---

## 📧 Getting Help

- **Issues/Bugs:** [GitHub Issues](https://github.com/yourusername/wallsync/issues)
- **Documentation:** [README.md](README.md) | [BUILD.md](BUILD.md)
- **Source Code:** [GitHub Repository](https://github.com/yourusername/wallsync)

---

## 📜 License

See [LICENSE](LICENSE) file for details.

## 🙏 Credits

Wallpapers provided by [Wallhaven.cc](https://wallhaven.cc)
