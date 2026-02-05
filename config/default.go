package config

import (
	"os"
	"path"
)

func GetDefaultConfig() Config {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic("Failed to get user home dir\n" + err.Error())
	}
	return Config{
		Keybinds: Keybinds{
			NextPage: []Keybind{
				{
					DeviceName: "_KEYBOARD_",
					Key:        106,
				},
			},
			PreviousPage: []Keybind{
				{
					DeviceName: "_KEYBOARD_",
					Key:        105,
				},
			},
			NextCategory:     []Keybind{},
			PreviousCategory: []Keybind{},
			ToggleBookmark: []Keybind{
				{
					DeviceName: "_KEYBOARD_",
					Key:        48,
				},
			},
			NextBookmark:     []Keybind{},
			PreviousBookmark: []Keybind{},
			ToggleWindow:     []Keybind{},
		},
		DcsInstallPath:    path.Join(homeDir, ".steam/steam/steamapps/common/DCSWorld"),
		DcsSavedGamesPath: path.Join(homeDir, ".steam/steam/steamapps/compatdata/223750/drive_c/users/steamusers/Saved Games/DCS"),
		ServerPort:        19021,
		DayNightMode:      0,
	}
}
