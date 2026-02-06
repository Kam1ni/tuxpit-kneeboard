package settingsview

import (
	"fmt"
	"tuxpit-kneeboard/config"
	"tuxpit-kneeboard/widgets"

	"github.com/mappu/miqt/qt6"
)

func createGeneralSettingsTab(conf *config.Config) *qt6.QWidget {
	content := qt6.NewQVBoxLayout2()
	content.SetContentsMargins(18, 18, 18, 18)
	gameInstallDirInput := widgets.NewFileInput3(conf.DcsInstallPath, "DCS Install directory")
	gameInstallDirInput.OnInput(func(s string) {
		conf.DcsInstallPath = s
	})

	savedGamesDirInput := widgets.NewFileInput3(conf.DcsSavedGamesPath, "DCS Saved games directory")
	savedGamesDirInput.OnInput(func(s string) {
		conf.DcsSavedGamesPath = s
	})

	portNumberInput := qt6.NewQSpinBox2()
	portNumberInput.SetMinimum(1)
	portNumberInput.SetMaximum(65535)
	portNumberInput.SetValue(int(conf.ServerPort))
	portNumberInput.OnValueChanged(func(value int) {
		conf.ServerPort = uint16(value)
	})

	dayNightModeInput := qt6.NewQCheckBox3("Toggle day/night mode button enabled")
	dayNightModeInput.SetToolTip(`Adds the option to sort for only kneeboard pages that do not contain "_Day_" or "_Night_" in the filename`)
	dayNightModeOrignalValue := conf.DayNightMode
	fmt.Println(dayNightModeOrignalValue, dayNightModeOrignalValue != config.DAY_NIGHT_MODE_DISABLED)
	dayNightModeInput.SetChecked(dayNightModeOrignalValue != config.DAY_NIGHT_MODE_DISABLED)
	dayNightModeInput.OnClickedWithChecked(func(checked bool) {
		if !checked {
			conf.DayNightMode = config.DAY_NIGHT_MODE_DISABLED
			return
		}
		if dayNightModeOrignalValue == config.DAY_NIGHT_MODE_DISABLED {
			conf.DayNightMode = config.DAY_NIGHT_MODE_DAY
			return
		}
		conf.DayNightMode = dayNightModeOrignalValue
	})

	content.AddWidget(gameInstallDirInput.QWidget())
	content.AddWidget(savedGamesDirInput.QWidget())
	content.AddWidget(widgets.NewLabeledInput("Server port to communicate with DCS", portNumberInput.QWidget).QWidget())
	content.AddWidget(dayNightModeInput.QWidget)
	content.AddStretch()

	widget := qt6.NewQWidget(nil)
	widget.SetLayout(content.QLayout)
	return widget
}
