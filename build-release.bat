@echo off
REM Build script for creating release binaries for all platforms (Windows version)

set VERSION=%1
if "%VERSION%"=="" set VERSION=v1.0.0

set OUTPUT_DIR=releases\%VERSION%

echo Building WallSync %VERSION% for all platforms...
echo.

REM Create output directory
if not exist "%OUTPUT_DIR%" mkdir "%OUTPUT_DIR%"

REM Build for Windows (amd64)
echo Building for Windows (amd64)...
set GOOS=windows
set GOARCH=amd64
go build -ldflags="-s -w" -o "%OUTPUT_DIR%\wallsync-windows-amd64.exe" main.go
go build -ldflags="-s -w" -o "%OUTPUT_DIR%\uninstaller-windows-amd64.exe" uninstaller\uninstaller.go

REM Build for macOS (amd64 - Intel)
echo Building for macOS (amd64 - Intel)...
set GOOS=darwin
set GOARCH=amd64
go build -ldflags="-s -w" -o "%OUTPUT_DIR%\wallsync-darwin-amd64" main.go
go build -ldflags="-s -w" -o "%OUTPUT_DIR%\uninstaller-darwin-amd64" uninstaller\uninstaller.go

REM Build for macOS (arm64 - Apple Silicon)
echo Building for macOS (arm64 - Apple Silicon)...
set GOOS=darwin
set GOARCH=arm64
go build -ldflags="-s -w" -o "%OUTPUT_DIR%\wallsync-darwin-arm64" main.go
go build -ldflags="-s -w" -o "%OUTPUT_DIR%\uninstaller-darwin-arm64" uninstaller\uninstaller.go

REM Build for Linux (amd64)
echo Building for Linux (amd64)...
set GOOS=linux
set GOARCH=amd64
go build -ldflags="-s -w" -o "%OUTPUT_DIR%\wallsync-linux-amd64" main.go
go build -ldflags="-s -w" -o "%OUTPUT_DIR%\uninstaller-linux-amd64" uninstaller\uninstaller.go

REM Build for Linux (arm64)
echo Building for Linux (arm64)...
set GOOS=linux
set GOARCH=arm64
go build -ldflags="-s -w" -o "%OUTPUT_DIR%\wallsync-linux-arm64" main.go
go build -ldflags="-s -w" -o "%OUTPUT_DIR%\uninstaller-linux-arm64" uninstaller\uninstaller.go

REM Reset environment variables
set GOOS=
set GOARCH=

echo.
echo Build complete! Release files created in %OUTPUT_DIR%\
echo.
echo Note: Archive creation requires 7-Zip or similar tool
echo You can manually create archives or use the following commands:
echo.
echo For Windows (PowerShell):
echo   Compress-Archive -Path "%OUTPUT_DIR%\wallsync-windows-amd64.exe","%OUTPUT_DIR%\uninstaller-windows-amd64.exe" -DestinationPath "%OUTPUT_DIR%\wallsync-windows-amd64.zip"
echo.
echo To create a GitHub release:
echo 1. git tag %VERSION%
echo 2. git push origin %VERSION%
echo 3. Go to GitHub ^> Releases ^> Draft a new release
echo 4. Upload the files from %OUTPUT_DIR%\
echo.

pause
