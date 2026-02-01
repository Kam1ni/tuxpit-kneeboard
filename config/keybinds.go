package config

import "slices"

type Keybind struct {
	DeviceName string `json:"deviceName"`
	Key        int    `json:"key"`
}

type Keybinds struct {
	NextPage         []Keybind `json:"nextPage"`
	PreviousPage     []Keybind `json:"previousPage"`
	NextCategory     []Keybind `json:"nextCategory"`
	PreviousCategory []Keybind `json:"previousCategory"`
	ToggleBookmark   []Keybind `json:"toggleBookmark"`
	NextBookmark     []Keybind `json:"nextBookmark"`
	PreviousBookmark []Keybind `json:"previousBookmark"`
	ToggleWindow     []Keybind `json:"toggleWindow"`
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

	return allNames
}
