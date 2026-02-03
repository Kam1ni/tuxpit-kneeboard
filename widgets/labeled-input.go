package widgets

import "github.com/mappu/miqt/qt"

type LabeledInput struct {
	container *qt.QVBoxLayout
	label     *qt.QLabel
	content   *qt.QWidget
}

func (l LabeledInput) GetLabel() *qt.QLabel {
	return l.label
}

func (l LabeledInput) QWidget() *qt.QWidget {
	widget := qt.NewQWidget(nil)
	widget.SetLayout(l.container.QLayout)
	widget.SetFixedHeight(75)
	return widget
}

func NewLabeledInput(label string, content *qt.QWidget) *LabeledInput {
	result := LabeledInput{}

	result.label = qt.NewQLabel3(label)

	result.container = qt.NewQVBoxLayout2()
	result.container.AddWidget(result.label.QWidget)
	result.container.AddWidget(content)

	return &result
}
