#!/bin/bash
# Build script for creating release binaries for all platforms

VERSION=${1:-"v1.0.0"}
OUTPUT_DIR="releases/$VERSION"

echo "Building WallSync $VERSION for all platforms..."

# Create output directory
mkdir -p "$OUTPUT_DIR"

# Build for Windows (amd64)
echo "Building for Windows (amd64)..."
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o "$OUTPUT_DIR/wallsync-windows-amd64.exe" main.go
GOOS=windows GOARCH=amd64 go build -ldflags="-s -w" -o "$OUTPUT_DIR/uninstaller-windows-amd64.exe" uninstaller/uninstaller.go

# Build for macOS (amd64 - Intel)
echo "Building for macOS (amd64 - Intel)..."
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o "$OUTPUT_DIR/wallsync-darwin-amd64" main.go
GOOS=darwin GOARCH=amd64 go build -ldflags="-s -w" -o "$OUTPUT_DIR/uninstaller-darwin-amd64" uninstaller/uninstaller.go

# Build for macOS (arm64 - Apple Silicon)
echo "Building for macOS (arm64 - Apple Silicon)..."
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o "$OUTPUT_DIR/wallsync-darwin-arm64" main.go
GOOS=darwin GOARCH=arm64 go build -ldflags="-s -w" -o "$OUTPUT_DIR/uninstaller-darwin-arm64" uninstaller/uninstaller.go

# Build for Linux (amd64)
echo "Building for Linux (amd64)..."
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o "$OUTPUT_DIR/wallsync-linux-amd64" main.go
GOOS=linux GOARCH=amd64 go build -ldflags="-s -w" -o "$OUTPUT_DIR/uninstaller-linux-amd64" uninstaller/uninstaller.go

# Build for Linux (arm64 - for ARM servers/Raspberry Pi)
echo "Building for Linux (arm64)..."
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o "$OUTPUT_DIR/wallsync-linux-arm64" main.go
GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o "$OUTPUT_DIR/uninstaller-linux-arm64" uninstaller/uninstaller.go

# Create archives
echo "Creating archives..."
cd "$OUTPUT_DIR"

# Windows
zip wallsync-windows-amd64.zip wallsync-windows-amd64.exe uninstaller-windows-amd64.exe
rm wallsync-windows-amd64.exe uninstaller-windows-amd64.exe

# macOS Intel
tar -czf wallsync-darwin-amd64.tar.gz wallsync-darwin-amd64 uninstaller-darwin-amd64
rm wallsync-darwin-amd64 uninstaller-darwin-amd64

# macOS Apple Silicon
tar -czf wallsync-darwin-arm64.tar.gz wallsync-darwin-arm64 uninstaller-darwin-arm64
rm wallsync-darwin-arm64 uninstaller-darwin-arm64

# Linux amd64
tar -czf wallsync-linux-amd64.tar.gz wallsync-linux-amd64 uninstaller-linux-amd64
rm wallsync-linux-amd64 uninstaller-linux-amd64

# Linux arm64
tar -czf wallsync-linux-arm64.tar.gz wallsync-linux-arm64 uninstaller-linux-arm64
rm wallsync-linux-arm64 uninstaller-linux-arm64

cd ../..

echo ""
echo "Build complete! Release files created in $OUTPUT_DIR/"
echo ""
echo "Files created:"
ls -lh "$OUTPUT_DIR/"
echo ""
echo "To create a GitHub release:"
echo "1. git tag $VERSION"
echo "2. git push origin $VERSION"
echo "3. Go to GitHub > Releases > Draft a new release"
echo "4. Upload the files from $OUTPUT_DIR/"
