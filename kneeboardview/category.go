package kneeboardview

import (
	"fmt"
	"os"
	"path"
	"regexp"
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
		if ValidImageFileRegex.MatchString(item.Name()) {
			list = append(list, path.Join(i.dir, item.Name()))
		}
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

func NewImageViewCategory(name string, dir string, label *qt6.QLabel) *imageViewCategory {
	view := imageViewCategory{name: name, dir: dir, label: label}
	view.loadImages()
	return &view
}
