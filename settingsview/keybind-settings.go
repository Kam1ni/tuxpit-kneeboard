package settingsview

import (
	"tuxpit-kneeboard/config"
	"tuxpit-kneeboard/widgets"

	"github.com/mappu/miqt/qt"
)

func createKeybindsTabs(conf *config.Config) *qt.QWidget {
	root := qt.NewQScrollArea2()
	contentBox := qt.NewQVBoxLayout2()

	contentBox.AddWidget(widgets.NewKeybindSetting2(conf.Keybinds.NextPage, "Next page", func(binds []config.Keybind) { conf.Keybinds.NextPage = binds }).QWidget())
	contentBox.AddWidget(widgets.NewKeybindSetting2(conf.Keybinds.PreviousPage, "Previous page", func(binds []config.Keybind) { conf.Keybinds.PreviousPage = binds }).QWidget())
	contentBox.AddWidget(widgets.NewKeybindSetting2(conf.Keybinds.NextCategory, "Next category", func(binds []config.Keybind) { conf.Keybinds.NextCategory = binds }).QWidget())
	contentBox.AddWidget(widgets.NewKeybindSetting2(conf.Keybinds.PreviousCategory, "Previous category", func(binds []config.Keybind) { conf.Keybinds.PreviousCategory = binds }).QWidget())
	contentBox.AddWidget(widgets.NewKeybindSetting2(conf.Keybinds.ToggleBookmark, "Toggle bookmark", func(binds []config.Keybind) { conf.Keybinds.ToggleBookmark = binds }).QWidget())
	contentBox.AddWidget(widgets.NewKeybindSetting2(conf.Keybinds.NextBookmark, "Next bookmark", func(binds []config.Keybind) { conf.Keybinds.NextBookmark = binds }).QWidget())
	contentBox.AddWidget(widgets.NewKeybindSetting2(conf.Keybinds.PreviousBookmark, "Previous bookmark", func(binds []config.Keybind) { conf.Keybinds.PreviousBookmark = binds }).QWidget())
	contentBox.AddWidget(widgets.NewKeybindSetting2(conf.Keybinds.ToggleWindow, "Toggle window", func(binds []config.Keybind) { conf.Keybinds.ToggleWindow = binds }).QWidget())
	contentBox.AddStretch()

	root.SetVerticalScrollBarPolicy(qt.ScrollBarAsNeeded)
	root.SetWidgetResizable(true)

	widget := qt.NewQWidget2()
	widget.SetLayout(contentBox.QLayout)
	root.SetWidget(widget)
	//widget := qt.NewQWidget2()
	//widget.SetLayout(contentBox.QLayout)
	return root.QWidget
}
