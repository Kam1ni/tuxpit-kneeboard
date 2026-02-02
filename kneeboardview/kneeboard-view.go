package kneeboardview

import (
	"fmt"
	"net"
	"os"
	"tuxpit-kneeboard/config"
	"tuxpit-kneeboard/inputlogger"

	"github.com/mappu/miqt/qt"
)

type View struct {
	currentCategoryIndex int
	categories           []*imageViewCategory
	aircraftCategory     *imageViewCategory
	terrainCategory      *imageViewCategory
	missionCategory      *imageViewCategory
	widget               *qt.QWidget
	inputLogger          *inputlogger.InputLogger
	bookmarks            []bookmark
	bookmarksContainer   *qt.QVBoxLayout
	server               *net.UDPConn
	config               config.Config
	missionTmpDir        string
	closed               bool
}

func (v View) Widget() *qt.QWidget {
	return v.widget
}

func (v *View) NextImage() {
	if v.categories[v.currentCategoryIndex].nextImage() {
		v.NextCategory()
	}
}

func (v *View) PreviousImage() {
	if v.categories[v.currentCategoryIndex].previousImage() {
		v.PreviousCategory()
	}
}

func (v *View) NextCategory() {
	for i := 0; i < len(v.categories); i++ {
		v.currentCategoryIndex++
		if v.currentCategoryIndex >= len(v.categories) {
			v.currentCategoryIndex = 0
		}
		if len(v.getSelectedCategory().sortedImages) != 0 {
			break
		}
	}
	v.categories[v.currentCategoryIndex].currentImage = ""
	fmt.Println(v.categories[v.currentCategoryIndex].currentImage)
	v.NextImage()
}

func (v *View) PreviousCategory() {
	for i := 0; i < len(v.categories); i++ {
		v.currentCategoryIndex--
		if v.currentCategoryIndex < 0 {
			v.currentCategoryIndex = len(v.categories) - 1
		}
		if len(v.getSelectedCategory().sortedImages) != 0 {
			break
		}
	}
	v.categories[v.currentCategoryIndex].currentImage = ""
	v.PreviousImage()
}

func (v *View) SelectCategory(catIndex int) {
	if len(v.categories[catIndex].sortedImages) == 0 {
		return
	}
	v.currentCategoryIndex = catIndex
	v.categories[v.currentCategoryIndex].currentImage = ""
	v.categories[v.currentCategoryIndex].nextImage()
}

func (v *View) getSelectedCategory() *imageViewCategory {
	return v.categories[v.currentCategoryIndex]
}

func CreateKneeboardView(conf config.Config) *View {
	v := View{config: conf}
	label := qt.NewQLabel3("")
	label.SetScaledContents(true)

	v.missionTmpDir = createMissionTmpDir()
	v.aircraftCategory = NewImageViewCategory("Aircraft", GetDcsAircraftDir(conf, ""), label)
	v.terrainCategory = NewImageViewCategory("Terrain", GetDcsTerrainDir(conf, ""), label)
	v.missionCategory = NewImageViewCategory("Mission", v.missionTmpDir, label)

	v.categories = []*imageViewCategory{
		v.aircraftCategory,
		v.terrainCategory,
		v.missionCategory,
	}

	root := qt.NewQVBoxLayout2()
	root.SetSpacing(0)

	tabs := qt.NewQHBoxLayout2()

	for i, cat := range v.categories {
		btn := qt.NewQPushButton3(cat.name)
		tabs.AddWidget(btn.QWidget)
		btn.OnClicked(func() {
			v.SelectCategory(i)
		})
	}

	body := qt.NewQHBoxLayout2()
	body.SetSpacing(0)
	v.bookmarksContainer = qt.NewQVBoxLayout2()

	bookmarksWidget := qt.NewQWidget2()
	bookmarksWidget.SetLayout(v.bookmarksContainer.QLayout)
	bookmarksWidget.SetFixedWidth(50)

	v.createBookmarksBar()

	body.AddWidget(bookmarksWidget)
	body.AddWidget(label.QWidget)

	topMenuWidget := qt.NewQWidget(nil)
	topMenuWidget.SetLayout(tabs.QLayout)
	topMenuWidget.SetFixedHeight(50)

	bodyWidget := qt.NewQWidget(nil)
	bodyWidget.SetLayout(body.QLayout)

	bottomToolBar := createBottomToolbar(&v)
	if bottomToolBar == nil {
		v.Close()
		return nil
	}
	bottomMenuWidget := qt.NewQWidget(nil)
	bottomMenuWidget.SetLayout(bottomToolBar.QLayout)
	bottomMenuWidget.SetFixedHeight(50)

	root.AddWidget(topMenuWidget)
	root.AddWidget(bodyWidget)
	root.AddWidget(bottomMenuWidget)

	initInputLoggert(&v)

	widget := qt.NewQWidget(nil)
	widget.SetLayout(root.QLayout)
	v.widget = widget

	v.server = runServer(&v)
	//	image := qt.NewQPixmap4("/home/kamil/Pictures/AI/Robokini.png")
	//	label.SetPixmap(image)

	v.NextImage()
	return &v
}

func (v *View) Close() {
	if v.closed {
		return
	}
	v.closed = true
	if v.server != nil {
		v.server.Close()
	}
	fmt.Println("Removing", v.missionTmpDir)
	os.RemoveAll(v.missionTmpDir)
	if v.inputLogger != nil {
		v.inputLogger.Close()
	}
}
