package main

import (
	"fmt"
	"os"
	"tuxpit-kneeboard/config"
	"tuxpit-kneeboard/kneeboardview"

	"github.com/mappu/miqt/qt"
)

func main() {
	qt.QGuiApplication_SetDesktopFileName("Tuxpit-Kneeboard")
	qt.NewQApplication(os.Args)

	mainWindow := qt.NewQMainWindow2()
	mainWindow.SetWindowTitle("Tuxpit Kneeboard")

	conf, err := config.ReadConfig()
	if err != nil {
		panic("Failed to read config\n" + err.Error())
	}

	view := kneeboardview.CreateKneeboardView(conf)
	defer view.Close()
	mainWindow.SetCentralWidget(view.Widget())
	mainWindow.SetMinimumHeight(100)
	mainWindow.SetMinimumWidth(100)
	mainWindow.Show()
	qt.QApplication_Exec()

	fmt.Println("OK")
}
