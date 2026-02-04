package kneeboardview

import (
	"fmt"
	"math/rand/v2"

	"github.com/mappu/miqt/qt"
	"github.com/mappu/miqt/qt/mainthread"
)

type bookmark struct {
	category  int
	imagePath string
	colors    colorCombo
}

func (b bookmark) isSameBookmark(b2 bookmark) bool {
	if b.imagePath != b2.imagePath {
		return false
	}
	if b.category != b2.category {
		return false
	}
	return true
}

func (v *View) getBookmarkIndex(b bookmark) int {
	for i, bm := range v.bookmarks {
		if b.isSameBookmark(bm) {
			return i
		}
	}
	return -1
}

func (v *View) addBookmark(bm bookmark) {
	bm.colors = colorCombos[rand.IntN(len(colorCombos))]
	v.bookmarks = append(v.bookmarks, bm)
	v.createBookmarksBar()
}

func (v *View) removeBookmark(index int) {
	v.bookmarks = append(v.bookmarks[:index], v.bookmarks[index+1:]...)
	v.createBookmarksBar()
}

func (v *View) toggleBookmark(bm bookmark) {
	mainthread.Wait(func() {
		index := v.getBookmarkIndex(bm)
		if index == -1 {
			v.addBookmark(bm)
		} else {
			v.removeBookmark(index)
		}
	})

}

func (v *View) removeAllBookmarks() {
	v.bookmarks = []bookmark{}
	v.createBookmarksBar()
}

func (v *View) selectBookmark(bm bookmark) {
	v.currentCategoryIndex = bm.category
	category := v.categories[bm.category]
	category.loadImage(bm.imagePath)
}

func (v *View) createBookmarksBar() {
	for v.bookmarksContainer.Count() > 0 {
		item := v.bookmarksContainer.TakeAt(0)
		if item != nil && item.Widget() != nil {
			item.Widget().DeleteLater()
		}
	}

	for i, bm := range v.bookmarks {
		btn := qt.NewQPushButton3(fmt.Sprintf("%d", i))
		btn.Thread().MoveToThread(v.bookmarksContainer.Thread())
		btn.OnClicked(func() {
			v.selectBookmark(bm)
		})
		btn.SetFixedHeight(25)
		btn.SetStyleSheet(fmt.Sprintf("background-color: %s; color: %s;", bm.colors.backgroundColor, bm.colors.foregroundColor))
		v.bookmarksContainer.AddWidget(btn.QWidget)
	}

	v.bookmarksContainer.AddStretch()
}

func (v *View) getCurrentBookmarkIndex() int {
	result := -1
	for i, bm := range v.bookmarks {
		if bm.category != v.currentCategoryIndex {
			continue
		}
		if bm.imagePath != v.getSelectedCategory().currentImage {
			continue
		}
		result = i
	}
	return result
}

func (v *View) NextBookmark() {
	if len(v.bookmarks) == 0 {
		return
	}
	index := v.getCurrentBookmarkIndex()
	index++
	if index >= len(v.bookmarks) {
		index = 0
	}

	v.currentCategoryIndex = v.bookmarks[index].category
	v.getSelectedCategory().loadImage(v.bookmarks[index].imagePath)
}

func (v *View) PreviousBookmark() {
	if len(v.bookmarks) == 0 {
		return
	}
	index := v.getCurrentBookmarkIndex()
	if index == -1 {
		index = len(v.bookmarks) - 1
	} else {
		index--
		if index < 0 {
			index = len(v.bookmarks) - 1
		}
	}

	v.currentCategoryIndex = v.bookmarks[index].category
	v.getSelectedCategory().loadImage(v.bookmarks[index].imagePath)
}
