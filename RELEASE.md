# Release Process Guide

This guide explains how to create and publish releases for WallSync on GitHub.

## Table of Contents
- [Prerequisites](#prerequisites)
- [Building Release Binaries](#building-release-binaries)
- [Creating a GitHub Release](#creating-a-github-release)
- [Automated Release with GitHub Actions](#automated-release-with-github-actions)
- [Release Checklist](#release-checklist)

## Prerequisites

1. **Go 1.16+** installed
2. **Git** installed and configured
3. **GitHub repository** set up
4. **Write access** to the repository
5. **7-Zip or tar/zip** for creating archives (Linux/macOS)

## Building Release Binaries

### Option 1: Using Build Scripts

**On Linux/macOS:**
```bash
chmod +x build-release.sh
./build-release.sh v1.0.0
```

**On Windows:**
```cmd
build-release.bat v1.0.0
```

This will create a `releases/v1.0.0/` directory with:
- `wallsync-windows-amd64.zip`
- `wallsync-darwin-amd64.tar.gz` (Intel Mac)
- `wallsync-darwin-arm64.tar.gz` (Apple Silicon)
- `wallsync-linux-amd64.tar.gz`
- `wallsync-linux-arm64.tar.gz`

### Option 2: Manual Build

```bash
# Create version directory
VERSION=v1.0.0
mkdir -p releases/$VERSION

# Windows
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o releases/$VERSION/wallsync.exe main.go
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o releases/$VERSION/uninstaller.exe uninstaller/uninstaller.go

# macOS Intel
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o releases/$VERSION/wallsync main.go
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o releases/$VERSION/uninstaller uninstaller/uninstaller.go

# macOS Apple Silicon
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o releases/$VERSION/wallsync main.go
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o releases/$VERSION/uninstaller uninstaller/uninstaller.go

# Linux
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o releases/$VERSION/wallsync main.go
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o releases/$VERSION/uninstaller uninstaller/uninstaller.go
```

**Note:** The `-ldflags="-s -w"` flags strip debug info and reduce binary size.

## Creating a GitHub Release

### Step 1: Commit All Changes

```bash
# Check status
git status

# Add and commit
git add .
git commit -m "chore: prepare release v1.0.0"
git push origin main
```

### Step 2: Create and Push Git Tag

```bash
# Create annotated tag
git tag -a v1.0.0 -m "Release v1.0.0 - Initial release with full cross-platform support"

# Push tag to GitHub
git push origin v1.0.0
```

### Step 3: Create GitHub Release

1. **Go to your repository on GitHub**
2. Click on **"Releases"** (right sidebar)
3. Click **"Draft a new release"**
4. **Fill in the release information:**

   **Tag version:** `v1.0.0` (select the tag you just pushed)

   **Release title:** `WallSync v1.0.0 - Initial Release`

   **Description:** Use the template below

5. **Upload the release files:**
   - Drag and drop all files from `releases/v1.0.0/` folder
   - Or click "Attach binaries" and select files

6. **Options:**
   - ☐ Set as pre-release (check if beta/alpha)
   - ☐ Set as latest release (check for stable releases)
   - ☑ Create a discussion for this release (optional)

7. Click **"Publish release"**

### Release Description Template

```markdown
# WallSync v1.0.0 - Initial Release

## 🎉 What's New

- ✨ Automatic wallpaper changing with configurable intervals (15 min to daily)
- 🔄 Background daemon mode for continuous operation
- 🔔 Desktop notifications when wallpapers change
- 🚀 Auto-start on login support
- 🎨 Multiple category selection (General, Anime, People)
- 🌐 Full cross-platform support (Windows, macOS, Linux)
- 💾 Smart caching with automatic cleanup
- ⚡ HTTP timeouts and robust error handling

## 📦 Downloads

### Windows
- **[wallsync-windows-amd64.zip]()**
  - For Windows 10/11 (64-bit)

### macOS
- **[wallsync-darwin-amd64.tar.gz]()** - For Intel Macs
- **[wallsync-darwin-arm64.tar.gz]()** - For Apple Silicon (M1/M2/M3)

### Linux
- **[wallsync-linux-amd64.tar.gz]()** - For 64-bit x86 systems
- **[wallsync-linux-arm64.tar.gz]()** - For ARM64 systems (Raspberry Pi, etc.)

**Supported Linux Desktop Environments:**
- GNOME
- KDE Plasma
- XFCE

## 📖 Installation

### Quick Start

**Windows:**
```cmd
1. Download wallsync-windows-amd64.zip
2. Extract to a folder
3. Run wallsync.exe
4. Follow the setup wizard
```

**macOS:**
```bash
1. Download wallsync-darwin-[amd64/arm64].tar.gz
2. Extract: tar -xzf wallsync-darwin-*.tar.gz
3. Make executable: chmod +x wallsync uninstaller
4. Run: ./wallsync
5. Follow the setup wizard
```

**Linux:**
```bash
1. Download wallsync-linux-amd64.tar.gz
2. Extract: tar -xzf wallsync-linux-amd64.tar.gz
3. Make executable: chmod +x wallsync uninstaller
4. Run: ./wallsync
5. Follow the setup wizard
```

For detailed instructions, see [BUILD.md](BUILD.md)

## 🔧 Usage

```bash
# Change wallpaper once
./wallsync -once

# Run in daemon mode (background service)
./wallsync -daemon

# Reconfigure preferences
./wallsync -reconfig

# Override interval (in minutes)
./wallsync -daemon -interval 30
```

## 📋 Requirements

- **Windows:** Windows 10 or later
- **macOS:** macOS 10.10 (Yosemite) or later
- **Linux:**
  - GNOME, KDE Plasma, or XFCE desktop environment
  - `notify-send` for notifications (usually pre-installed)

## 🐛 Known Issues

- None reported yet

## 📝 Full Changelog

- Initial release with full feature set
- Cross-platform support for Windows, macOS, and Linux
- Automatic wallpaper rotation
- Configurable intervals and categories
- Desktop notifications
- Auto-start on login

## 🙏 Acknowledgments

Wallpapers sourced from [Wallhaven.cc](https://wallhaven.cc)

---

**Full Documentation:** [README.md](README.md) | [BUILD.md](BUILD.md)

**Report Issues:** [GitHub Issues](https://github.com/yourusername/wallsync/issues)
```

## Automated Release with GitHub Actions

Create `.github/workflows/release.yml`:

```yaml
name: Release

on:
  push:
    tags:
      - 'v*'

jobs:
  build:
    name: Build and Release
    runs-on: ubuntu-latest
    steps:
      - name: Checkout code
        uses: actions/checkout@v3

      - name: Set up Go
        uses: actions/setup-go@v4
        with:
          go-version: '1.21'

      - name: Build binaries
        run: |
          chmod +x build-release.sh
          ./build-release.sh ${{ github.ref_name }}

      - name: Create Release
        uses: softprops/action-gh-release@v1
        with:
          files: releases/${{ github.ref_name }}/*
          draft: false
          prerelease: false
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

### Enable GitHub Actions

1. Create `.github/workflows/` directory
2. Add `release.yml` file
3. Commit and push
4. When you push a tag, GitHub Actions will automatically:
   - Build binaries for all platforms
   - Create a GitHub release
   - Upload all binaries

## Release Checklist

Before creating a release, ensure:

- [ ] All tests pass
- [ ] Documentation is updated (README.md, BUILD.md)
- [ ] Version number is updated (if stored in code)
- [ ] CHANGELOG.md is updated (if you have one)
- [ ] All commits are pushed to main branch
- [ ] Build scripts work on all platforms
- [ ] Binaries are tested on target platforms (if possible)
- [ ] .gitignore excludes build artifacts
- [ ] No sensitive information in code or config files

### Post-Release Tasks

- [ ] Verify all download links work
- [ ] Test installation on at least one platform
- [ ] Announce release (social media, forums, etc.)
- [ ] Update documentation site (if applicable)
- [ ] Close resolved issues
- [ ] Create milestone for next release

## Version Numbering

Follow [Semantic Versioning](https://semver.org/):

- **MAJOR** (v2.0.0): Breaking changes
- **MINOR** (v1.1.0): New features, backwards compatible
- **PATCH** (v1.0.1): Bug fixes, backwards compatible

Examples:
- `v1.0.0` - First stable release
- `v1.1.0` - Added new feature
- `v1.0.1` - Fixed bug
- `v2.0.0` - Breaking API changes
- `v1.0.0-beta.1` - Pre-release

## Troubleshooting

### Build fails with "command not found"
- Ensure Go is installed: `go version`
- Check PATH includes Go bin directory

### Git tag already exists
```bash
# Delete local tag
git tag -d v1.0.0

# Delete remote tag
git push origin :refs/tags/v1.0.0

# Create new tag
git tag -a v1.0.0 -m "Release message"
git push origin v1.0.0
```

### Can't upload files to GitHub Release
- Check file size limits (2GB per file for free accounts)
- Ensure you have write access to repository
- Try using GitHub CLI: `gh release create v1.0.0 releases/v1.0.0/*`

### Files are too large
- Ensure `-ldflags="-s -w"` is used in build command
- Consider using UPX to compress binaries further
- Split large packages into multiple archives

## Additional Resources

- [GitHub Releases Documentation](https://docs.github.com/en/repositories/releasing-projects-on-github)
- [Semantic Versioning](https://semver.org/)
- [Keep a Changelog](https://keepachangelog.com/)
- [GitHub Actions Documentation](https://docs.github.com/en/actions)
