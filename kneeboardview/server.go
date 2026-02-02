package kneeboardview

import (
	"archive/zip"
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"regexp"
	"strings"

	"github.com/mappu/miqt/qt/mainthread"
)

var savedGamesPath = regexp.MustCompile(`^C:\\users\\steamusers\\Saved Games\\DCS`)
var validMissionKneeboardPath = regexp.MustCompile(`^kneeboard\/images\/.*\.(png|jpe?g)$`)

func runServer(view *View) *net.UDPConn {

	udpAddr, err := net.ResolveUDPAddr("udp", ":8912")
	if err != nil {
		panic("Failed to resolve udp address for server\n" + err.Error())
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		panic("Failed to listen on udp\n" + err.Error())
	}

	go func() {
		for {
			var buf [512]byte
			n, _, err := conn.ReadFromUDP(buf[0:])
			if err != nil {
				fmt.Println("Failed to read udp message\n", err.Error())
				return
			}

			message := string(buf[:n])

			fmt.Printf("Received via udp %s \n", message)

			parts := strings.Split(message, ":")
			fmt.Println(parts)
			switch parts[0] {
			case "aircraft":
				view.onAircraftReceived(parts[1])
			case "terrain":
				view.onTerrainReceived(parts[1])
			case "mission":
				mainthread.Wait(func() { view.onMissionReceived(parts[1]) })
			}
		}
	}()

	return conn
}

func (view *View) onAircraftReceived(aircraftName string) {
	newDir := GetDcsAircraftDir(view.config, aircraftName)
	fmt.Println("Loading aircraft kneeboard from", newDir)
	view.aircraftCategory.dir = GetDcsAircraftDir(view.config, aircraftName)
	view.aircraftCategory.loadImages()
	if view.getSelectedCategory() == view.aircraftCategory {
		view.aircraftCategory.nextImage()
	}
}

func (view *View) onTerrainReceived(terrainName string) {
	newDir := GetDcsTerrainDir(view.config, terrainName)
	fmt.Println("Loading terrain kneeboard from", newDir)
	view.terrainCategory.dir = newDir
	view.terrainCategory.loadImages()
	if view.getSelectedCategory() == view.terrainCategory {
		view.terrainCategory.nextImage()
	}
}

func (view *View) onMissionReceived(pth string) {
	err := os.RemoveAll(view.missionTmpDir)
	if err != nil {
		fmt.Println("Failed to remove old mission tmp dir")
	}
	view.missionCategory.sortedImages = []string{}
	view.missionTmpDir = createMissionTmpDir()

	if pth == "--from-last-mission-track-file--" {
		pth = path.Join(view.config.DcsSavedGamesPath, "Tracks/LastMissionTrack.trk")
	} else if savedGamesPath.MatchString(pth) {
		pth = strings.ReplaceAll(savedGamesPath.ReplaceAllString(pth, view.config.DcsSavedGamesPath), "\\", "/")
	} else {
		pth = path.Join(view.config.DcsInstallPath, strings.ReplaceAll(pth, "\\", "/"))
	}
	stats, err := os.Stat(pth)
	if err != nil {
		fmt.Printf("Failed to access mission: %s\n%s\n", pth, err.Error())
		return
	}
	if stats.IsDir() {
		fmt.Printf("Failed to access mission: %s\n%s\n", pth, "Mission is not a zip file")
		return
	}

	zipReader, err := zip.OpenReader(pth)
	if err != nil {
		fmt.Printf("Failed to read mission archive: %s\n%s\n", pth, err.Error())
		return
	}
	defer zipReader.Close()

	fmt.Println("Mission path: ", pth)
	fmt.Println("Writing kneeboards to ", view.missionTmpDir)
	for _, file := range zipReader.File {
		if !validMissionKneeboardPath.MatchString(strings.ToLower(file.Name)) {
			continue
		}
		fileName := path.Join(view.missionTmpDir, path.Base(file.Name))
		err = writeTempFileFromZipFile(file, fileName)
		if err != nil {
			fmt.Printf("Failed to write mission file %s\n%s", file.Name, err.Error())
		}
	}

	view.missionCategory.dir = view.missionTmpDir
	view.missionCategory.loadImages()
	if view.getSelectedCategory() == view.missionCategory {
		view.missionCategory.nextImage()
	}
}

func writeTempFileFromZipFile(file *zip.File, targetPath string) error {
	reader, err := file.Open()
	if err != nil {
		return err
	}
	defer reader.Close()

	b, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	return os.WriteFile(targetPath, b, os.ModePerm)
}
