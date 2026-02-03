package kneeboardview

import (
	"fmt"
	"tuxpit-kneeboard/settingsview"

	"github.com/mappu/miqt/qt"
)

func createBottomToolbar(v *View) *qt.QHBoxLayout {
	root := qt.NewQHBoxLayout2()

	clearBtn := qt.NewQPushButton3(" Clear bookmarks")
	clearBtn.OnClicked(func() {
		v.removeAllBookmarks()
	})
	root.AddWidget(clearBtn.QWidget)

	root.AddStretch()

	settingsButton := qt.NewQPushButton3(" Settings")
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
	return root
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
}
