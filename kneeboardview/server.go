package kneeboardview

import (
	"fmt"
	"net"
	"path"
	"regexp"
	"strings"
)

var savedGamesPath = regexp.MustCompile(`^C:\\users\\steamusers\\Saved Games\\DCS`)

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

			fmt.Printf("Received %s via udp\n", message)

			parts := strings.Split(message, ":")
			fmt.Println(parts)
			if parts[0] == "aircraft" {
				fmt.Println(len(parts[1]))
				newDir := GetDcsAircraftDir(view.config, parts[1])
				fmt.Println(newDir)
				if newDir == view.aircraftCategory.dir {
					fmt.Println("Same dir")
					continue
				}
				view.aircraftCategory.dir = GetDcsAircraftDir(view.config, parts[1])
				view.aircraftCategory.loadImages()
				if view.getSelectedCategory() == view.aircraftCategory {
					view.aircraftCategory.nextImage()
				}
			}
			if parts[0] == "terrain" {
				newDir := GetDcsTerrainDir(view.config, parts[1])
				if newDir == view.terrainCategory.dir {
					continue
				}
				view.terrainCategory.dir = newDir
				view.terrainCategory.loadImages()
				if view.getSelectedCategory() == view.terrainCategory {
					view.terrainCategory.nextImage()
				}
			}
			if parts[0] == "mission" {
				pth := parts[1]
				if savedGamesPath.MatchString(pth) {
					pth = strings.ReplaceAll(savedGamesPath.ReplaceAllString(pth, view.config.DcsSavedGamesPath), "\\", "/")
				} else {
					pth = path.Join(view.config.DcsInstallPath, strings.ReplaceAll(pth, "\\", "/"))
				}
				fmt.Println("Mission path: ", pth)
				// TODO: Get kneeboard files from mission
			}
		}
	}()

	return conn
}
