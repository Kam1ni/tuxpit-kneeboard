package settingsview

import (
	"tuxpit-kneeboard/config"
	"tuxpit-kneeboard/widgets"

	"github.com/mappu/miqt/qt"
)

func createGeneralSettingsTab(conf *config.Config) *qt.QWidget {
	content := qt.NewQVBoxLayout2()
	gameInstallDirInput := widgets.NewFileInput3(conf.DcsInstallPath, "DCS Install directory")
	gameInstallDirInput.OnInput(func(s string) {
		conf.DcsInstallPath = s
	})

	savedGamesDirInput := widgets.NewFileInput3(conf.DcsSavedGamesPath, "DCS Saved games directory")
	savedGamesDirInput.OnInput(func(s string) {
		conf.DcsSavedGamesPath = s
	})

	portNumberInput := qt.NewQSpinBox2()
	portNumberInput.SetMinimum(1)
	portNumberInput.SetMaximum(65535)
	portNumberInput.SetValue(int(conf.ServerPort))
	portNumberInput.OnValueChanged(func(value int) {
		conf.ServerPort = uint16(value)
	})

	content.AddWidget(gameInstallDirInput.QWidget())
	content.AddWidget(savedGamesDirInput.QWidget())
	content.AddWidget(widgets.NewLabeledInput("Server port to communicate with DCS", portNumberInput.QWidget).QWidget())
	content.AddStretch()

	widget := qt.NewQWidget(nil)
	widget.SetLayout(content.QLayout)
	return widget
}
