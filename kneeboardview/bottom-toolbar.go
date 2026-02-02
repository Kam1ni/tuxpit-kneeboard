package kneeboardview

import (
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
		settingsview.CreateSettingsWindow(&v.config)
	})

	if !v.config.ComesFromFile {
		settingsview.CreateSettingsWindow(&v.config)
		if !v.config.ComesFromFile {
			return nil
		}
	}

	root.AddWidget(settingsButton.QWidget)
	return root
}
