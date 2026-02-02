package widgets

import "github.com/mappu/miqt/qt"

type FileInput struct {
	input          *qt.QLineEdit
	button         *qt.QPushButton
	container      *qt.QVBoxLayout
	label          *qt.QLabel
	onInputHandler func(val string)
}

func NewFileInput() *FileInput {
	result := FileInput{}

	result.label = qt.NewQLabel2()

	result.input = qt.NewQLineEdit2()
	result.input.OnTextEdited(func(param1 string) {
		if result.onInputHandler == nil {
			return
		}
		result.onInputHandler(param1)
	})

	result.button = qt.NewQPushButton3("")
	result.button.OnClicked(func() {
		result.openFileDialog()
	})

	inputContainer := qt.NewQHBoxLayout2()
	inputContainer.AddWidget(result.input.QWidget)
	inputContainer.AddWidget(result.button.QWidget)
	inputContainerWidget := qt.NewQWidget(nil)
	inputContainerWidget.SetLayout(inputContainer.QLayout)
	inputContainerWidget.SetFixedHeight(50)

	result.container = qt.NewQVBoxLayout2()
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

func (f *FileInput) QWidget() *qt.QWidget {
	widget := qt.NewQWidget(nil)
	widget.SetLayout(f.container.QLayout)
	widget.SetFixedHeight(75)
	return widget
}

func (f *FileInput) openFileDialog() {
	dialog := qt.NewQFileDialog3()
	dialog.SetDirectory(f.input.Text())
	dialog.SetFileMode(qt.QFileDialog__DirectoryOnly)
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
