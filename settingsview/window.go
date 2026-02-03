package settingsview

import (
	"tuxpit-kneeboard/config"
	"tuxpit-kneeboard/widgets"

	"github.com/mappu/miqt/qt"
)

func CreateSettingsWindow(conf *config.Config) {
	root := qt.NewQDialog2()

	tempConf := conf.Clone()

	content := qt.NewQFormLayout2()
	gameInstallDirInput := widgets.NewFileInput3(tempConf.DcsInstallPath, "DCS Install directory")
	gameInstallDirInput.OnInput(func(s string) {
		tempConf.DcsInstallPath = s
	})

	savedGamesDirInput := widgets.NewFileInput3(tempConf.DcsSavedGamesPath, "DCS Saved games directory")
	savedGamesDirInput.OnInput(func(s string) {
		tempConf.DcsSavedGamesPath = s
	})

	portNumberInput := qt.NewQSpinBox2()
	portNumberInput.SetValue(int(tempConf.ServerPort))
	portNumberInput.SetMinimum(1)
	portNumberInput.SetMaximum(65535)
	portNumberInput.OnValueChanged(func(value int) {
		tempConf.ServerPort = uint16(value)
	})

	content.AddWidget(gameInstallDirInput.QWidget())
	content.AddWidget(savedGamesDirInput.QWidget())
	content.AddWidget(widgets.NewLabeledInput("Server port to communicate with DCS", portNumberInput.QWidget).QWidget())

	confirmButtons := widgets.CreateFormConfirmButtons()
	content.AddWidget(confirmButtons.QWidget())

	confirmButtons.OnConfirm(func() {
		err := ensureDcsPluginIsInstalled(tempConf)
		if err != nil {
			qt.NewQMessageBox3(qt.QMessageBox__Critical, "Error", err.Error()).Exec()
			return
		}
		*conf = tempConf
		conf.ComesFromFile = true
		config.WriteConfig(*conf)
		root.Close()
	})

	confirmButtons.OnCancel(func() {
		root.Close()
	})

	if !conf.ComesFromFile {
		confirmButtons.SetCancelDisabled(true)
	}

	root.SetLayout(content.QLayout)

	root.SetWindowTitle("Tuxpit Kneeboard Settings")

	root.Exec()
}
