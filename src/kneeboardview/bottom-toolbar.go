package kneeboardview

import (
	"fmt"
	"tuxpit-kneeboard/config"
	"tuxpit-kneeboard/settingsview"

	"github.com/mappu/miqt/qt6"
)

type bottomToolbar struct {
	root                     *qt6.QHBoxLayout
	toggleDayNightModeButton *qt6.QPushButton
}

func (b *bottomToolbar) Widget() *qt6.QWidget {
	widget := qt6.NewQWidget(nil)
	widget.SetLayout(b.root.QLayout)
	widget.SetSizePolicy2(qt6.QSizePolicy__Expanding, qt6.QSizePolicy__Minimum)
	return widget
}

func createBottomToolbar(v *View) *bottomToolbar {
	result := bottomToolbar{
		root:                     qt6.NewQHBoxLayout2(),
		toggleDayNightModeButton: qt6.NewQPushButton3("Day"),
	}
	result.root.SetContentsMargins(0, 0, 0, 0)
	result.root.SetSpacing(0)

	clearBtn := qt6.NewQPushButton3(" Clear bookmarks")
	clearBtn.OnClicked(func() {
		v.removeAllBookmarks()
	})
	clearBtn.SetFocusPolicy(qt6.ClickFocus)
	result.root.AddWidget(clearBtn.QWidget)

	result.toggleDayNightModeButton.SetSizePolicy(*qt6.NewQSizePolicy2(qt6.QSizePolicy__Expanding, qt6.QSizePolicy__Ignored))
	result.toggleDayNightModeButton.SetFocusPolicy(qt6.ClickFocus)
	result.toggleDayNightModeButton.OnClicked(func() {
		if v.config.DayNightMode == config.DAY_NIGHT_MODE_DAY {
			v.config.DayNightMode = config.DAY_NIGHT_MODE_NIGHT
			result.toggleDayNightModeButton.SetText("Night")
		} else {
			v.config.DayNightMode = config.DAY_NIGHT_MODE_DAY
			result.toggleDayNightModeButton.SetText("Day")
		}
		err := config.WriteConfig(v.config)
		if err != nil {
			fmt.Println(err.Error())
		}
		v.ReloadImages()
	})
	switch v.config.DayNightMode {
	case config.DAY_NIGHT_MODE_DISABLED:
		result.toggleDayNightModeButton.SetVisible(false)
	case config.DAY_NIGHT_MODE_NIGHT:
		result.toggleDayNightModeButton.SetText("Night")
	}
	result.root.AddWidget(result.toggleDayNightModeButton.QWidget)

	previousPageButton := qt6.NewQPushButton3("")
	previousPageButton.SetFocusPolicy(qt6.ClickFocus)
	previousPageButton.OnClicked(func() {
		v.PreviousPage()
	})
	previousPageButton.SetSizePolicy(*qt6.NewQSizePolicy2(qt6.QSizePolicy__Expanding, qt6.QSizePolicy__Ignored))

	nextPageButton := qt6.NewQPushButton3("")
	nextPageButton.SetFocusPolicy(qt6.ClickFocus)
	nextPageButton.OnClicked(func() {
		v.NextPage()
	})
	nextPageButton.SetSizePolicy(*qt6.NewQSizePolicy2(qt6.QSizePolicy__Expanding, qt6.QSizePolicy__Ignored))

	result.root.AddWidget(previousPageButton.QWidget)
	result.root.AddWidget(nextPageButton.QWidget)

	//	result.root.AddStretch()

	settingsButton := qt6.NewQPushButton3(" Settings")
	settingsButton.SetFocusPolicy(qt6.ClickFocus)
	settingsButton.OnClicked(func() {
		showSettings(v)
	})

	if !v.config.ComesFromFile {
		showSettings(v)
		if !v.config.ComesFromFile {
			return nil
		}
	}

	result.root.AddWidget(settingsButton.QWidget)

	return &result
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
	v.SetDayNightMode(v.config.DayNightMode)
	initInputLoggert(v)
}
