package kneeboardview

import "github.com/mappu/miqt/qt"

func createTabs(v *View) *qt.QWidget {
	tabs := qt.NewQHBoxLayout2()

	for i, cat := range v.categories {
		btn := qt.NewQPushButton3(cat.name)
		tabs.AddWidget(btn.QWidget)
		btn.OnClicked(func() {
			v.SelectCategory(i)
		})
	}

	widget := qt.NewQWidget(nil)
	widget.SetLayout(tabs.QLayout)
	widget.SetSizePolicy(*qt.NewQSizePolicy2(qt.QSizePolicy__Expanding, qt.QSizePolicy__Minimum))
	return widget
}
