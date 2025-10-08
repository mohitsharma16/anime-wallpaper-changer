package config

import (
	"encoding/json"
	"os"
)

// Config represents the user's preferences.
type Config struct {
	Categories      []string `json:"categories"`
	Purity          []string `json:"purity"`
	Autostart       bool     `json:"autostart"`
	ChangeInterval  int      `json:"change_interval"`  // in minutes
	ShowNotification bool    `json:"show_notification"`
}

// LoadConfig reads the configuration from a file.
func LoadConfig(file string) (*Config, error) {
	f, err := os.Open(file)
	if err != nil {
		if os.IsNotExist(err) {
			return defaultConfig(), nil
		}
		return nil, err
	}
	defer f.Close()

	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		return nil, err
	}
	return &cfg, nil
}

// SaveConfig saves the configuration to a file.
func SaveConfig(file string, cfg *Config) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := json.NewEncoder(f)
	encoder.SetIndent("", "  ")
	return encoder.Encode(cfg)
}

func defaultConfig() *Config {
	return &Config{
		Categories:       []string{"010"}, // Anime by default
		Purity:           []string{"100"}, // SFW
		Autostart:        false,
		ChangeInterval:   60,               // 60 minutes by default
		ShowNotification: true,
	}
}