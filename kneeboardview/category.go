package kneeboardview

import (
	"fmt"
	"os"
	"path"
	"regexp"
	"slices"
	"sort"
	"strings"
	"tuxpit-kneeboard/config"

	"github.com/mappu/miqt/qt6"
)

var ValidImageFileRegex = regexp.MustCompile(`(\.png|\.jpe?g)`)

func GetDcsAircraftDir(conf config.Config, aircraft string) string {
	return path.Join(conf.DcsSavedGamesPath, "Kneeboard", aircraft)
}

func GetDcsTerrainDir(conf config.Config, terrain string) string {
	baseDir := path.Join(conf.DcsInstallPath, "Mods/terrains", terrain)
	entries, err := os.ReadDir(baseDir)
	if err != nil {
		fmt.Printf("Failed to read base terrain dir\n%s", err.Error())
	}
	for _, entry := range entries {
		if strings.ToLower(entry.Name()) == "kneeboard" && entry.IsDir() {
			return path.Join(conf.DcsInstallPath, "Mods/terrains", terrain, entry.Name())
		}
	}

	return path.Join(conf.DcsInstallPath, "Mods/terrains", terrain, "Kneeboard")
}

func createMissionTmpDir() string {
	tempDirPath := path.Join(os.TempDir(), "tuxpit-kneeboard-mission")
	err := os.Mkdir(tempDirPath, os.ModePerm)
	if err != nil {
		if !strings.Contains(err.Error(), "file exists") {
			panic("Failed to create tmp dir\n" + err.Error())
		}
	}
	return tempDirPath
}

type imageViewCategory struct {
	name         string
	currentImage string
	sortedImages []string
	dir          string
	label        *qt6.QLabel
	view         *View
}

func (i *imageViewCategory) reloadImages() {
	currentImageName := path.Base(i.currentImage)
	i.loadImages()
	if i.view.config.DayNightMode == config.DAY_NIGHT_MODE_DISABLED {
		return
	}
	if i.view.config.DayNightMode == config.DAY_NIGHT_MODE_DAY && strings.Contains(currentImageName, "_Day_") {
		return
	}
	if i.view.config.DayNightMode == config.DAY_NIGHT_MODE_NIGHT && strings.Contains(currentImageName, "_Night_") {
		return
	}
	switch i.view.config.DayNightMode {
	case config.DAY_NIGHT_MODE_DAY:
		currentImageName = strings.ReplaceAll(currentImageName, "_Night_", "_Day_")
	case config.DAY_NIGHT_MODE_NIGHT:
		currentImageName = strings.ReplaceAll(currentImageName, "_Day_", "_Night_")
	}
	fullPath := path.Join(i.dir, currentImageName)
	if !slices.Contains(i.sortedImages, fullPath) {
		return
	}
	i.currentImage = fullPath
	if i.view.getSelectedCategory() != i {
		return
	}
	i.loadImage(i.currentImage)
}

func (i *imageViewCategory) loadImages() {
	result, err := os.ReadDir(i.dir)
	if err != nil {
		fmt.Println("Could not load images", err.Error())
		i.sortedImages = []string{}
		return
	}

	list := []string{}
	for _, item := range result {
		if item.IsDir() {
			continue
		}
		if !ValidImageFileRegex.MatchString(item.Name()) {
			continue
		}
		if i.view.config.DayNightMode == config.DAY_NIGHT_MODE_DAY && strings.Contains(item.Name(), "_Night_") {
			continue
		}
		if i.view.config.DayNightMode == config.DAY_NIGHT_MODE_NIGHT && strings.Contains(item.Name(), "_Day_") {
			continue
		}
		list = append(list, path.Join(i.dir, item.Name()))
	}

	sort.Strings(list)
	i.sortedImages = list
}

func (v *imageViewCategory) loadImage(imgPath string) {
	pxmap := qt6.NewQPixmap4(imgPath)
	v.label.SetPixmap(pxmap)
	v.currentImage = imgPath
}

func (v *imageViewCategory) nextPage() bool {
	if len(v.sortedImages) == 0 {
		return false
	}
	index := -1
	for i := range v.sortedImages {
		if v.sortedImages[i] == v.currentImage {
			index = i
			break
		}
	}

	if index >= len(v.sortedImages)-1 {
		return true
	}
	v.loadImage(v.sortedImages[index+1])
	return false
}

func (v *imageViewCategory) previousPage() bool {
	if len(v.sortedImages) == 0 {
		return false
	}
	index := len(v.sortedImages)
	for i := range v.sortedImages {
		if v.sortedImages[i] == v.currentImage {
			index = i
			break
		}
	}

	if index == 0 {
		return true
	}
	v.loadImage(v.sortedImages[index-1])
	return false
}

func NewImageViewCategory(name string, dir string, label *qt6.QLabel, view *View) *imageViewCategory {
	cat := imageViewCategory{name: name, dir: dir, label: label, view: view}
	cat.loadImages()
	return &cat
}
