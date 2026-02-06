package kneeboardview

import (
	"fmt"
	"net"
	"os"
	"tuxpit-kneeboard/config"
	"tuxpit-kneeboard/inputlogger"

	_ "embed"

	"github.com/mappu/miqt/qt6"
)

//go:embed style.qss
var windowStyle string

type View struct {
	currentCategoryIndex int
	categories           []*imageViewCategory
	aircraftCategory     *imageViewCategory
	terrainCategory      *imageViewCategory
	missionCategory      *imageViewCategory
	widget               *qt6.QWidget
	inputLogger          *inputlogger.InputLogger
	bookmarks            []bookmark
	bookmarksContainer   *qt6.QVBoxLayout
	server               *net.UDPConn
	config               config.Config
	missionTmpDir        string
	closed               bool
	mainWindow           *qt6.QMainWindow
	bottomToolbar        *bottomToolbar
}

func (v View) Widget() *qt6.QWidget {
	return v.widget
}

func (v View) ReloadImages() {
	for _, cat := range v.categories {
		cat.reloadImages()
	}
}

func (v *View) NextPage() {
	if v.categories[v.currentCategoryIndex].nextPage() {
		v.NextCategory()
	}
}

func (v *View) PreviousPage() {
	if v.categories[v.currentCategoryIndex].previousPage() {
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
	v.NextPage()
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
	v.PreviousPage()
}

func (v *View) SelectCategory(catIndex int) {
	if len(v.categories[catIndex].sortedImages) == 0 {
		return
	}
	v.currentCategoryIndex = catIndex
	v.categories[v.currentCategoryIndex].currentImage = ""
	v.categories[v.currentCategoryIndex].nextPage()
}

func (v *View) SetDayNightMode(mode config.DayNightModaType) {
	if v.bottomToolbar == nil {
		return
	}
	if v.bottomToolbar.toggleDayNightModeButton == nil {
		return
	}
	v.config.DayNightMode = mode
	switch v.config.DayNightMode {
	case config.DAY_NIGHT_MODE_DISABLED:
		v.bottomToolbar.toggleDayNightModeButton.SetVisible(false)
	case config.DAY_NIGHT_MODE_DAY:
		v.bottomToolbar.toggleDayNightModeButton.SetVisible(true)
		v.bottomToolbar.toggleDayNightModeButton.SetText("Day")
	case config.DAY_NIGHT_MODE_NIGHT:
		v.bottomToolbar.toggleDayNightModeButton.SetVisible(true)
		v.bottomToolbar.toggleDayNightModeButton.SetText("Night")
	}
	v.ReloadImages()
	_ = config.WriteConfig(v.config)
}

func (v *View) getSelectedCategory() *imageViewCategory {
	return v.categories[v.currentCategoryIndex]
}

func CreateKneeboardView(conf config.Config, mainWindow *qt6.QMainWindow) *View {
	v := View{config: conf, mainWindow: mainWindow}
	label := qt6.NewQLabel3("")
	label.SetScaledContents(true)

	v.missionTmpDir = createMissionTmpDir()
	v.aircraftCategory = NewImageViewCategory("Aircraft", GetDcsAircraftDir(conf, ""), label, &v)
	v.terrainCategory = NewImageViewCategory("Terrain", GetDcsTerrainDir(conf, ""), label, &v)
	v.missionCategory = NewImageViewCategory("Mission", v.missionTmpDir, label, &v)

	v.categories = []*imageViewCategory{
		v.aircraftCategory,
		v.terrainCategory,
		v.missionCategory,
	}

	root := qt6.NewQVBoxLayout2()
	root.SetContentsMargins(0, 0, 0, 0)
	root.SetSpacing(0)

	body := qt6.NewQHBoxLayout2()
	body.SetSpacing(0)
	body.SetContentsMargins(0, 0, 0, 0)

	v.bookmarksContainer = qt6.NewQVBoxLayout2()
	v.bookmarksContainer.SetSpacing(0)
	v.bookmarksContainer.SetContentsMargins(0, 0, 0, 0)
	bookmarksWidget := qt6.NewQWidget2()
	bookmarksWidget.SetLayout(v.bookmarksContainer.QLayout)
	bookmarksWidget.SetFixedWidth(80)

	v.createBookmarksBar()

	body.AddWidget(bookmarksWidget)
	body.AddWidget(label.QWidget)

	bodyWidget := qt6.NewQWidget(nil)
	bodyWidget.SetLayout(body.QLayout)
	bodyWidget.SetSizePolicy2(qt6.QSizePolicy__Expanding, qt6.QSizePolicy__Expanding)

	v.bottomToolbar = createBottomToolbar(&v)
	if v.bottomToolbar == nil {
		v.Close()
		return nil
	}

	root.AddWidget(createTabs(&v))
	root.AddWidget(bodyWidget)
	root.AddWidget(v.bottomToolbar.Widget())

	initInputLoggert(&v)

	widget := qt6.NewQWidget(nil)
	widget.SetLayout(root.QLayout)
	widget.SetStyleSheet(windowStyle)
	widget.SetObjectName(*qt6.NewQAnyStringView3("mainTuxpitWindowContianer"))
	v.widget = widget

	v.server = createServer(&v)
	//	image := qt6.NewQPixmap4("/home/kamil/Pictures/AI/Robokini.png")
	//	label.SetPixmap(image)

	v.NextPage()
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
