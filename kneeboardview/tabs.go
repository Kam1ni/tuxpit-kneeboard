package kneeboardview

import "github.com/mappu/miqt/qt6"

func createTabs(v *View) *qt6.QWidget {
	tabs := qt6.NewQHBoxLayout2()
	tabs.SetContentsMargins(0, 0, 0, 0)
	tabs.SetSpacing(0)

	for i, cat := range v.categories {
		btn := qt6.NewQPushButton3(cat.name)
		tabs.AddWidget(btn.QWidget)
		btn.OnClicked(func() {
			v.SelectCategory(i)
		})
	}

	widget := qt6.NewQWidget(nil)
	widget.SetLayout(tabs.QLayout)
	widget.SetSizePolicy(*qt6.NewQSizePolicy2(qt6.QSizePolicy__Expanding, qt6.QSizePolicy__Minimum))
	return widget
}
