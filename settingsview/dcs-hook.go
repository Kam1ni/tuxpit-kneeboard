package settingsview

import (
	_ "embed"
	"fmt"
	"os"
	"path"
	"strings"
	"tuxpit-kneeboard/config"
)

//go:embed tuxpit-kneeboard.lua
var dcsHookScript string

func ensureDcsPluginIsInstalled(conf config.Config) error {
	entries, err := os.ReadDir(conf.DcsSavedGamesPath)
	if err != nil {
		return fmt.Errorf("Failed to read %s\n%s", conf.DcsSavedGamesPath, err.Error())
	}

	found := false
	pth := conf.DcsSavedGamesPath

	for _, entry := range entries {
		if strings.ToLower(entry.Name()) == "scripts" {
			found = true
			pth = path.Join(pth, entry.Name())
			break
		}
	}

	if !found {
		pth = path.Join(conf.DcsSavedGamesPath, "Scripts")
		err = os.Mkdir(pth, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Failed to create %s\n%s", pth, err.Error())
		}
	}

	found = false
	entries, err = os.ReadDir(pth)
	if err != nil {
		return fmt.Errorf("Failed to read %s\n%s", pth, err.Error())
	}
	for _, entry := range entries {
		if strings.ToLower(entry.Name()) == "hooks" {
			found = true
			pth = path.Join(pth, entry.Name())
			break
		}
	}

	if !found {
		pth = path.Join(pth, "Hooks")
		err = os.Mkdir(pth, os.ModePerm)
		if err != nil {
			return fmt.Errorf("Failed to create %s\n%s", pth, err.Error())
		}
	}

	// add the server port to the script
	scriptSrc := fmt.Sprintf("local serverPort = %d\n%s", conf.ServerPort, dcsHookScript)

	pth = path.Join(pth, "tuxpit-kneeboard.lua")
	return os.WriteFile(pth, []byte(scriptSrc), os.ModePerm)
}
