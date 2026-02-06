package main

import (
	"fmt"
	"os"
	"tuxpit-kneeboard/config"
	"tuxpit-kneeboard/kneeboardview"

	"github.com/mappu/miqt/qt6"
)

func main() {
	qt6.QGuiApplication_SetDesktopFileName("Tuxpit-Kneeboard")
	_ = qt6.NewQApplication(os.Args)

	mainWindow := qt6.NewQMainWindow2()
	mainWindow.SetWindowTitle("Tuxpit Kneeboard")

	conf, err := config.ReadConfig()
	if err != nil {
		panic("Failed to read config\n" + err.Error())
	}

	view := kneeboardview.CreateKneeboardView(conf, mainWindow)
	if view == nil {
		return
	}
	defer view.Close()
	mainWindow.SetCentralWidget(view.Widget())
	mainWindow.SetMinimumHeight(100)
	mainWindow.SetMinimumWidth(100)
	mainWindow.Resize(700, 1000)
	mainWindow.SetWindowFlag2(qt6.Tool, true)
	mainWindow.SetWindowFlag2(qt6.WindowStaysOnTopHint, true)
	mainWindow.Show()
	mainWindow.OnCloseEvent(func(super func(event *qt6.QCloseEvent), event *qt6.QCloseEvent) {
		qt6.QCoreApplication_Quit()
	})
	qt6.QApplication_Exec()

	fmt.Println("OK")
}
