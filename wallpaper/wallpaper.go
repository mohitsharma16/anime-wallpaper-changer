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

	"anime-wallpaper-changer/config"
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

	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

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
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

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
