package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

type Config struct {
	ComesFromFile     bool     `json:"-"`
	Keybinds          Keybinds `json:"keybinds"`
	DcsInstallPath    string   `json:"dcsInstallPath"`
	DcsSavedGamesPath string   `json:"dcsSavedGamesPath"`
	ServerPort        uint16   `json:"serverPort"`
}

func (c Config) Clone() Config {
	return Config{
		Keybinds:          c.Keybinds.Clone(),
		DcsInstallPath:    c.DcsInstallPath,
		DcsSavedGamesPath: c.DcsSavedGamesPath,
		ServerPort:        c.ServerPort,
	}
}

func GetDefaultConfigPath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic("Failed to get user config dir\n" + err.Error())
	}
	return path.Join(configDir, "tuxpit-kneeboard/config.json")
}

func ensureConfigDirExists() error {
	configDirPath := path.Dir(GetDefaultConfigPath())
	stats, err := os.Stat(configDirPath)
	if err == nil {
		if !stats.IsDir() {
			return fmt.Errorf("%s is not a directory", configDirPath)
		}
		return nil
	}

	parts := strings.Split(configDirPath, "/")
	currentPath := "/"
	for _, part := range parts {
		if part == "" {
			continue
		}
		currentPath = path.Join(currentPath, part)
		if currentPath == "" {
			continue
		}
		fileInfo, err := os.Stat(currentPath)
		if err != nil {
			err = os.Mkdir(currentPath, os.ModePerm)
			if err != nil {
				return fmt.Errorf("Failed to create dir %s\n%s", currentPath, err.Error())
			}
			continue
		}
		if !fileInfo.IsDir() {
			return fmt.Errorf("%s is not a directory", currentPath)
		}
	}

	return nil
}

func WriteConfig(config Config) error {
	configContent, err := json.MarshalIndent(config, "", "	")
	if err != nil {
		return fmt.Errorf("Failed to marshal config json\n%s", err.Error())
	}
	configPath := GetDefaultConfigPath()
	err = os.WriteFile(configPath, configContent, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Failed to create config json\n%s", err.Error())
	}
	return nil
}

// Reads the config from  ~/.config/tuxpit-kneeboard/config.json
func ReadConfig() (Config, error) {
	err := ensureConfigDirExists()
	if err != nil {
		return Config{}, err
	}

	_, err = os.Stat(GetDefaultConfigPath())
	if err != nil {
		return GetDefaultConfig(), nil
	}

	configPath := GetDefaultConfigPath()
	confContent, err := os.ReadFile(configPath)
	if err != nil {
		return Config{}, fmt.Errorf("Failed to read config file\n%s", err.Error())
	}

	config := Config{}
	err = json.Unmarshal(confContent, &config)
	if err != nil {
		return Config{}, fmt.Errorf("Config json at %s has errors\n%s", configPath, err.Error())
	}
	config.ComesFromFile = true
	return config, nil
}
