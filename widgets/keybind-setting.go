package widgets

import (
	"tuxpit-kneeboard/config"
	"tuxpit-kneeboard/inputlogger"

	"github.com/mappu/miqt/qt"
	"github.com/mappu/miqt/qt/mainthread"
)

type KeybindSetting struct {
	rootElement    *qt.QGroupBox
	content        *qt.QVBoxLayout
	onConfigUpdate func([]config.Keybind)
	binds          []config.Keybind
}

func (k *KeybindSetting) QWidget() *qt.QWidget {
	return k.rootElement.QWidget
}

func NewKeybindSetting(binds []config.Keybind, label string) *KeybindSetting {
	result := KeybindSetting{
		rootElement: qt.NewQGroupBox3(label),
		content:     qt.NewQVBoxLayout2(),
		binds:       binds,
	}
	result.rootElement.SetSizePolicy2(qt.QSizePolicy__Expanding, qt.QSizePolicy__Minimum)

	result.populateBinds()

	result.rootElement.SetLayout(result.content.QLayout)
	return &result
}

func NewKeybindSetting2(binds []config.Keybind, label string, callback func(binds []config.Keybind)) *KeybindSetting {
	result := NewKeybindSetting(binds, label)
	result.OnConfigUpdate(callback)
	return result
}

func (k *KeybindSetting) OnConfigUpdate(handler func([]config.Keybind)) {
	k.onConfigUpdate = handler
}

func (k *KeybindSetting) createBindingWidget(bind config.Keybind) *qt.QWidget {
	box := qt.NewQHBoxLayout2()
	label := qt.NewQLabel3(bind.ToString())
	label.SetSizePolicy2(qt.QSizePolicy__Expanding, qt.QSizePolicy__Fixed)

	deleteButton := qt.NewQPushButton3("Remove")

	box.AddWidget(label.QWidget)
	box.AddWidget(deleteButton.QWidget)

	widget := qt.NewQWidget2()
	widget.SetLayout(box.Layout())

	deleteButton.OnClicked(func() {
		index := -1
		for i, currentBind := range k.binds {
			if currentBind == bind {
				index = i
				break
			}
		}
		if index == -1 {
			return
		}
		widget.Delete()
		k.binds = append(k.binds[:index], k.binds[index+1:]...)
		if k.onConfigUpdate != nil {
			k.onConfigUpdate(k.binds)
		}
	})
	return widget
}

func (k *KeybindSetting) populateBinds() {
	for _, bind := range k.binds {
		k.content.AddWidget(k.createBindingWidget(bind))
	}
	addButton := qt.NewQPushButton3("Add bind")
	addButton.OnClicked(func() {
		binding := k.addBinding()
		if binding == nil {
			return
		}

		k.binds = append(k.binds, *binding)
		k.content.InsertWidget(k.content.Count()-1, k.createBindingWidget(*binding))
		if k.onConfigUpdate != nil {
			k.onConfigUpdate(k.binds)
		}
	})
	k.content.AddWidget(addButton.QWidget)
	k.rootElement.SetLayout(k.content.QLayout)
}

func (k *KeybindSetting) addBinding() *config.Keybind {
	window := qt.NewQDialog2()

	var result *config.Keybind = nil

	root := qt.NewQVBoxLayout2()
	label := qt.NewQLabel3("Press any button")
	label.SetAlignment(qt.AlignCenter)
	label.SetSizePolicy2(qt.QSizePolicy__Expanding, qt.QSizePolicy__Expanding)
	label.GrabKeyboard()
	defer label.ReleaseKeyboard()

	actions := NewFormConfirmButtons()
	actions.SetConfirmDisabled(true)
	actions.OnCancel(func() {
		result = nil
		window.Close()
	})
	actions.OnConfirm(func() {
		window.Close()
	})

	root.AddWidget(label.QWidget)
	root.AddWidget(actions.QWidget())
	window.SetLayout(root.QLayout)
	window.SetFixedSize2(400, 200)
	window.Raise()
	window.ActivateWindow()

	var logger *inputlogger.InputLogger
	// Initialize the logger on a seperate thread since it can add a delay to the window appearing
	go func() {
		inputlogger.NewAllInputsLogger(func(deviceName string, button int) {
			mainthread.Start(func() {
				result = &config.Keybind{DeviceName: deviceName, Key: button}
				label.SetText(result.ToString())
				actions.SetConfirmDisabled(false)
			})
		})
	}()
	defer func() {
		if logger != nil {
			logger.Close()
		}
	}()
	qt.QCoreApplication_ProcessEvents()
	window.Exec()
	return result

}
