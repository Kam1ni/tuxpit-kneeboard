package widgets

import "github.com/mappu/miqt/qt6"

type FormConfirmButtons struct {
	confirmButton *qt6.QPushButton
	cancelButton  *qt6.QPushButton
	container     *qt6.QHBoxLayout
}

func (f *FormConfirmButtons) QWidget() *qt6.QWidget {
	widget := qt6.NewQWidget(nil)
	widget.SetLayout(f.container.QLayout)
	widget.SetSizePolicy2(qt6.QSizePolicy__Expanding, qt6.QSizePolicy__Minimum)
	return widget
}

func NewFormConfirmButtons() *FormConfirmButtons {
	result := FormConfirmButtons{}
	result.confirmButton = qt6.NewQPushButton3(" Confirm")
	result.cancelButton = qt6.NewQPushButton3(" Cancel")
	result.container = qt6.NewQHBoxLayout2()
	result.container.AddWidget(result.cancelButton.QWidget)
	result.container.AddStretch()
	result.container.AddWidget(result.confirmButton.QWidget)
	result.container.SetContentsMargins(0, 0, 0, 0)
	return &result
}

func (f *FormConfirmButtons) OnConfirm(handler func()) {
	f.confirmButton.OnClicked(handler)
}

func (f *FormConfirmButtons) OnCancel(handler func()) {
	f.cancelButton.OnClicked(handler)
}

func (f *FormConfirmButtons) SetCancelDisabled(disabled bool) {
	f.cancelButton.SetDisabled(disabled)
}

func (f *FormConfirmButtons) SetConfirmDisabled(disabled bool) {
	f.confirmButton.SetDisabled(disabled)
}
