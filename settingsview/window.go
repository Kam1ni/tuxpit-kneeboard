package settingsview

import (
	"tuxpit-kneeboard/config"
	"tuxpit-kneeboard/widgets"

	"github.com/mappu/miqt/qt"
)

func CreateSettingsWindow(conf *config.Config) {
	root := qt.NewQDialog2()

	tempConf := conf.Clone()

	confirmButtons := widgets.NewFormConfirmButtons()

	rootContainer := qt.NewQVBoxLayout2()

	tabs := qt.NewQTabWidget2()
	tabs.AddTab(createGeneralSettingsTab(&tempConf), "General")
	tabs.AddTab(createKeybindsTabs(&tempConf), "Keybinds")

	rootContainer.AddWidget(tabs.QWidget)
	rootContainer.AddWidget(confirmButtons.QWidget())

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

	root.SetLayout(rootContainer.QLayout)
	root.SetWindowTitle("Tuxpit Kneeboard Settings")
	root.SetFixedSize2(640, 480)
	root.Exec()
}
