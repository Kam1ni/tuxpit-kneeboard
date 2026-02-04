package kneeboardview

import (
	"fmt"
	"tuxpit-kneeboard/config"
	"tuxpit-kneeboard/inputlogger"
)

func KeybindMatches(kb []config.Keybind, devName string, key int) bool {
	for _, kb := range kb {
		if kb.DeviceName == devName && kb.Key == key {
			return true
		}
	}
	return false
}

func onKeyPress(v *View, deviceName string, key int) {
	fmt.Printf("Got event %s %d\n", deviceName, key)
	if KeybindMatches(v.config.Keybinds.NextPage, deviceName, key) {
		v.NextImage()
	}
	if KeybindMatches(v.config.Keybinds.PreviousPage, deviceName, key) {
		v.PreviousImage()
	}
	if KeybindMatches(v.config.Keybinds.NextCategory, deviceName, key) {
		v.NextCategory()
	}
	if KeybindMatches(v.config.Keybinds.PreviousCategory, deviceName, key) {
		v.PreviousCategory()
	}
	if KeybindMatches(v.config.Keybinds.ToggleBookmark, deviceName, key) {
		v.toggleBookmark(bookmark{category: v.currentCategoryIndex, imagePath: v.getSelectedCategory().currentImage})
	}
	if KeybindMatches(v.config.Keybinds.NextBookmark, deviceName, key) {
		v.NextBookmark()
	}
	if KeybindMatches(v.config.Keybinds.PreviousBookmark, deviceName, key) {
		v.PreviousBookmark()
	}
	if KeybindMatches(v.config.Keybinds.ToggleWindow, deviceName, key) {
		if v.mainWindow.IsHidden() {
			v.mainWindow.Show()
		} else {
			v.mainWindow.Hide()
		}
	}
}

func initInputLoggert(v *View) {
	if v.inputLogger != nil {
		fmt.Println("Input logger already inited")
		return
	}

	v.inputLogger = inputlogger.NewInputLogger(v.config.Keybinds.GetAllDeviceNames(), func(deviceName string, button int) {
		onKeyPress(v, deviceName, button)
	})
}
