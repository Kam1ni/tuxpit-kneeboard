package widgets

import "github.com/mappu/miqt/qt6"

type LabeledInput struct {
	container *qt6.QVBoxLayout
	label     *qt6.QLabel
	content   *qt6.QWidget
}

func (l LabeledInput) GetLabel() *qt6.QLabel {
	return l.label
}

func (l LabeledInput) QWidget() *qt6.QWidget {
	widget := qt6.NewQWidget(nil)
	widget.SetLayout(l.container.QLayout)
	widget.SetFixedHeight(75)
	return widget
}

func NewLabeledInput(label string, content *qt6.QWidget) *LabeledInput {
	result := LabeledInput{}

	result.label = qt6.NewQLabel3(label)

	result.container = qt6.NewQVBoxLayout2()
	result.container.AddWidget(result.label.QWidget)
	result.container.AddWidget(content)
	result.container.SetContentsMargins(0, 0, 0, 0)

	return &result
}
