package wallpaper

import (
	"encoding/json"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"time"

	"wallsync/config"
)

// WallhavenResponse represents the response from the Wallhaven API.
type WallhavenResponse struct {
	Data []Wallpaper `json:"data"`
}

// Wallpaper represents a single wallpaper from Wallhaven.
type Wallpaper struct {
	ID   string `json:"id"`
	Path string `json:"path"`
}

// GetRandomWallpaper fetches a random wallpaper from Wallhaven.
func GetRandomWallpaper(cfg *config.Config) (string, error) {
	// Construct the API URL
	categories := calculateBitmask(cfg.Categories)
	purity := calculateBitmask(cfg.Purity)
	url := fmt.Sprintf("https://wallhaven.cc/api/v1/search?categories=%s&purity=%s&sorting=random&atleast=1920x1080&ratios=16x9", categories, purity)

	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 30 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return "", fmt.Errorf("failed to fetch wallpaper from API: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("API returned status code %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	var wallhavenResp WallhavenResponse
	if err := json.Unmarshal(body, &wallhavenResp); err != nil {
		return "", err
	}

	if len(wallhavenResp.Data) == 0 {
		return "", fmt.Errorf("no wallpapers found")
	}

	// Seed the random number generator
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Select a random wallpaper
	randomWallpaper := wallhavenResp.Data[r.Intn(len(wallhavenResp.Data))]

	return randomWallpaper.Path, nil
}

// DownloadWallpaper downloads a wallpaper to a local file.
func DownloadWallpaper(url string, filepath string) error {
	// Create HTTP client with timeout
	client := &http.Client{
		Timeout: 60 * time.Second, // Longer timeout for downloading images
	}

	resp, err := client.Get(url)
	if err != nil {
		return fmt.Errorf("failed to download wallpaper: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("download failed with status code %d", resp.StatusCode)
	}

	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = io.Copy(out, resp.Body)
	return err
}

func calculateBitmask(values []string) string {
	var mask int
	for _, value := range values {
		i, _ := strconv.ParseInt(value, 2, 0)
		mask |= int(i)
	}
	return fmt.Sprintf("%03b", mask)
}
