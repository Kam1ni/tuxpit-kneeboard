package widgets

import (
	"github.com/mappu/miqt/qt6"
)

type FileInput struct {
	input          *qt6.QLineEdit
	button         *qt6.QPushButton
	container      *qt6.QVBoxLayout
	label          *qt6.QLabel
	onInputHandler func(val string)
}

func NewFileInput() *FileInput {
	result := FileInput{}

	result.label = qt6.NewQLabel2()
	result.input = qt6.NewQLineEdit2()
	result.input.OnTextEdited(func(param1 string) {
		if result.onInputHandler == nil {
			return
		}
		result.onInputHandler(param1)
	})

	result.button = qt6.NewQPushButton3("")
	result.button.OnClicked(func() {
		result.openFileDialog()
	})

	inputContainer := qt6.NewQHBoxLayout2()
	inputContainer.AddWidget(result.input.QWidget)
	inputContainer.AddWidget(result.button.QWidget)
	inputContainer.SetContentsMargins(0, 0, 0, 0)
	inputContainerWidget := qt6.NewQWidget(nil)
	inputContainerWidget.SetLayout(inputContainer.QLayout)
	inputContainerWidget.SetFixedHeight(50)

	result.container = qt6.NewQVBoxLayout2()
	result.container.AddWidget(result.label.QWidget)
	result.container.AddWidget(inputContainerWidget)

	return &result
}

func NewFileInput2(value string) *FileInput {
	result := NewFileInput()
	result.input.SetText(value)
	return result
}

func NewFileInput3(value string, label string) *FileInput {
	result := NewFileInput2(value)
	result.label.SetText(label)
	return result
}

func (f *FileInput) OnInput(handler func(string)) {
	f.onInputHandler = handler
}

func (f *FileInput) QWidget() *qt6.QWidget {
	widget := qt6.NewQWidget(nil)
	widget.SetLayout(f.container.QLayout)
	widget.SetSizePolicy2(qt6.QSizePolicy__Expanding, qt6.QSizePolicy__Minimum)
	return widget
}

func (f *FileInput) openFileDialog() {
	dialog := qt6.NewQFileDialog3()
	dialog.SetDirectory(f.input.Text())
	dialog.SetFileMode(qt6.QFileDialog__Directory)
	dialog.OnFileSelected(func(file string) {
		f.input.SetText(file)
		if f.onInputHandler != nil {
			f.onInputHandler(file)
		}
		dialog.Close()
	})
	if f.input.Text() != "" {
		dialog.SetWindowTitle(f.label.Text())
	} else {
		dialog.SetWindowTitle("File chooser")
	}
	dialog.Exec()
}
