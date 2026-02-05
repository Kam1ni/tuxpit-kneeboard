package kneeboardview

import (
	"fmt"
	"tuxpit-kneeboard/settingsview"

	"github.com/mappu/miqt/qt6"
)

func createBottomToolbar(v *View) *qt6.QWidget {
	root := qt6.NewQHBoxLayout2()
	root.SetContentsMargins(0, 0, 0, 0)
	root.SetSpacing(0)

	clearBtn := qt6.NewQPushButton3(" Clear bookmarks")
	clearBtn.OnClicked(func() {
		v.removeAllBookmarks()
	})
	root.AddWidget(clearBtn.QWidget)

	previousPageButton := qt6.NewQPushButton3("")
	previousPageButton.OnClicked(func() {
		v.PreviousPage()
	})
	previousPageButton.SetSizePolicy(*qt6.NewQSizePolicy2(qt6.QSizePolicy__Expanding, qt6.QSizePolicy__Ignored))

	nextPageButton := qt6.NewQPushButton3("")
	nextPageButton.OnClicked(func() {
		v.NextPage()
	})
	nextPageButton.SetSizePolicy(*qt6.NewQSizePolicy2(qt6.QSizePolicy__Expanding, qt6.QSizePolicy__Ignored))

	root.AddWidget(previousPageButton.QWidget)
	root.AddWidget(nextPageButton.QWidget)

	//	root.AddStretch()

	settingsButton := qt6.NewQPushButton3(" Settings")
	settingsButton.OnClicked(func() {
		showSettings(v)
	})

	if !v.config.ComesFromFile {
		showSettings(v)
		if !v.config.ComesFromFile {
			return nil
		}
	}

	root.AddWidget(settingsButton.QWidget)

	widget := qt6.NewQWidget(nil)
	widget.SetLayout(root.QLayout)
	widget.SetSizePolicy2(qt6.QSizePolicy__Expanding, qt6.QSizePolicy__Minimum)
	return widget
}

func showSettings(v *View) {
	settingsview.CreateSettingsWindow(&v.config)
	if v.server != nil {
		err := v.server.Close()
		if err != nil {
			fmt.Println("Failed to close UDP server", err.Error())
		}
	}
	v.server = createServer(v)
	if v.inputLogger != nil {
		v.inputLogger.Close()
		v.inputLogger = nil
	}
	initInputLoggert(v)
}
