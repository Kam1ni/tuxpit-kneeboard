package config

import (
	"fmt"
	"slices"

	"github.com/MarinX/keylogger"
)

type Keybind struct {
	DeviceName string `json:"deviceName"`
	Key        int    `json:"key"`
}

type Keybinds struct {
	NextPage           []Keybind `json:"nextPage"`
	PreviousPage       []Keybind `json:"previousPage"`
	NextCategory       []Keybind `json:"nextCategory"`
	PreviousCategory   []Keybind `json:"previousCategory"`
	ToggleBookmark     []Keybind `json:"toggleBookmark"`
	NextBookmark       []Keybind `json:"nextBookmark"`
	PreviousBookmark   []Keybind `json:"previousBookmark"`
	ToggleWindow       []Keybind `json:"toggleWindow"`
	ToggleDayNightMode []Keybind `json:"toggleDayNightMode"`
}

func (k Keybinds) GetAllDeviceNames() []string {
	allNames := []string{}

	addDeviceNames := func(binds []Keybind) {
		for _, bind := range binds {
			if slices.Contains(allNames, bind.DeviceName) {
				continue
			}
			allNames = append(allNames, bind.DeviceName)
		}
	}

	addDeviceNames(k.NextPage)
	addDeviceNames(k.PreviousPage)
	addDeviceNames(k.NextCategory)
	addDeviceNames(k.PreviousCategory)
	addDeviceNames(k.ToggleBookmark)
	addDeviceNames(k.NextBookmark)
	addDeviceNames(k.PreviousBookmark)
	addDeviceNames(k.ToggleWindow)
	addDeviceNames(k.ToggleDayNightMode)

	return allNames
}

func (k Keybinds) Clone() Keybinds {
	cloneBinds := func(binds []Keybind) []Keybind {
		result := make([]Keybind, len(binds))
		copy(result, binds)
		return result
	}

	return Keybinds{
		NextPage:           cloneBinds(k.NextPage),
		PreviousPage:       cloneBinds(k.PreviousPage),
		NextCategory:       cloneBinds(k.NextCategory),
		PreviousCategory:   cloneBinds(k.PreviousCategory),
		ToggleBookmark:     cloneBinds(k.ToggleBookmark),
		NextBookmark:       cloneBinds(k.NextBookmark),
		PreviousBookmark:   cloneBinds(k.PreviousBookmark),
		ToggleWindow:       cloneBinds(k.ToggleWindow),
		ToggleDayNightMode: cloneBinds(k.ToggleDayNightMode),
	}
}

func (k Keybind) ToString() string {
	if k.DeviceName == "_KEYBOARD_" {
		return fmt.Sprintf("Keyboard: Key %s", (&keylogger.InputEvent{Code: uint16(k.Key), Type: keylogger.EvKey}).KeyString())
	}
	return fmt.Sprintf("%s: Key %d", k.DeviceName, k.Key)
}
