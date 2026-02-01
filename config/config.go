package config

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"
)

type Config struct {
	Keybinds          Keybinds `json:"keybinds"`
	DcsInstallPath    string   `json:"dcsInstallPath"`
	DcsSavedGamesPath string   `json:"dcsSavedGamesPath"`
}

func GetDefaultConfigPath() string {
	configDir, err := os.UserConfigDir()
	if err != nil {
		panic("Failed to get user config dir\n" + err.Error())
	}
	return path.Join(configDir, "tuxpit-kneeboard/config.json")
}

func ensureConfigFileExists() error {
	configPath := GetDefaultConfigPath()
	stats, err := os.Stat(configPath)
	if err == nil {
		if stats.IsDir() {
			return fmt.Errorf("%s is a directory. Json file was expected", configPath)
		}
		return nil
	}

	parts := strings.Split(path.Dir(configPath), "/")
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

	err = WriteConfig(GetDefaultConfig())
	if err != nil {
		return fmt.Errorf("Failed to create default\n%s", err.Error())
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
	err := ensureConfigFileExists()
	if err != nil {
		return Config{}, err
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
	return config, nil
}
