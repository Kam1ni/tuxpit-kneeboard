package widgets

import "github.com/mappu/miqt/qt"

type FormConfirmButtons struct {
	confirmButton *qt.QPushButton
	cancelButton  *qt.QPushButton
	container     *qt.QHBoxLayout
}

func (f *FormConfirmButtons) QWidget() *qt.QWidget {
	widget := qt.NewQWidget(nil)
	widget.SetLayout(f.container.QLayout)
	widget.SetSizePolicy2(qt.QSizePolicy__Expanding, qt.QSizePolicy__Minimum)
	return widget
}

func NewFormConfirmButtons() *FormConfirmButtons {
	result := FormConfirmButtons{}
	result.confirmButton = qt.NewQPushButton3(" Confirm")
	result.cancelButton = qt.NewQPushButton3(" Cancel")
	result.container = qt.NewQHBoxLayout2()
	result.container.AddWidget(result.cancelButton.QWidget)
	result.container.AddStretch()
	result.container.AddWidget(result.confirmButton.QWidget)
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
