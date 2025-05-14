package lem

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Parsing() (*Colony, string) {
	if len(os.Args) != 2 {
		fmt.Println("Error: Please provide exactly one file name.")
		return nil, ""
	}

	colony := &Colony{
		rooms: make(map[string]*Room),
		links: make(map[string][]string),
		start: nil,
		end:   nil,
	}

	file, err := os.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error: Unable to read file:", err)
		return nil, ""
	}

	lines := strings.Split(string(file), "\n")

	AntIndex := 0
	for i, x := range lines {
		if x != "" && x[0] == '#' {
			continue
		} else {
			AntIndex = i
			break
		}
	}

	ant, err := strconv.Atoi(strings.TrimSpace(lines[AntIndex]))
	if err != nil || ant <= 0 || ant > 10000000 {
		fmt.Println("ERROR: invalid data format")
		return nil, ""
	}
	colony.NumAnts = ant
	var end, start bool
	linekes:=false
	for r := 1; r < len(lines); r++ {
		line := strings.TrimSpace(lines[r])

		if line == "" || line[0] == 'L' {
			continue
		}

		if line[0] == '#' && line != "##start" && line != "##end" {
			continue
		}

		room := strings.Fields(line)
		if len(room) == 3 {
			if linekes {
				fmt.Println("Error")
				return nil,""
			}
			if start {

				if len(room) != 3 {
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}
				if room[1] == "" || room[2] == "" {
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}
				RoomX, err := strconv.Atoi(room[1])
				if err != nil {
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}
				RoomY, err := strconv.Atoi(room[2])
				if err != nil {
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}
				colony.start = &Room{
					name: room[0],
					x:    RoomX,
					y:    RoomY,
				}
				start = false
				continue
			} else if end {
				
				if len(room) != 3 {
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}
				if room[1] == "" || room[2] == "" {
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}
				RoomX, err := strconv.Atoi(room[1])
				if err != nil {
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}
				RoomY, err := strconv.Atoi(room[2])
				if err != nil {
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}
				colony.end = &Room{
					name: room[0],
					x:    RoomX,
					y:    RoomY,
				}
				end=false
				continue
			} else {
				if _, exists := colony.rooms[room[0]]; exists {
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}

				if room[1] == "" || room[2] == "" {
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}

				RoomX, err := strconv.Atoi(room[1])
				if err != nil {
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}
				RoomY, err := strconv.Atoi(room[2])
				if err != nil {
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}
				for _, existingRoom := range colony.rooms {
					if RoomX == existingRoom.x && RoomY == existingRoom.y {
						fmt.Println("ERROR: invalid data format")
						return nil, ""
					}
				}

				colony.rooms[room[0]] = &Room{
					name: room[0],
					x:    RoomX,
					y:    RoomY,
				}
				continue
			}
		} else if line == "##start" {
			if r+1 >= len(lines) {
				fmt.Println("ERROR: invalid data format")
				return nil, ""
			}
			if start {
				fmt.Println("Error")
				return nil , ""
			}
			start = true

		} else if line == "##end" {
			if r+1 >= len(lines) {
				fmt.Println("ERROR: invalid data format")
				return nil, ""
			}
			if end {
				fmt.Println("Error")
				return nil , ""
			}
			end = true

		} else {
			linekes=true
			link := strings.Split(line, "-")
			if len(link) == 2 {
				colony.links[link[0]] = append(colony.links[link[0]], link[1])
				colony.links[link[1]] = append(colony.links[link[1]], link[0])
			} else {
				fmt.Println("ERROR: invalid data format")
				return nil, ""
			}
		}

	}

	if colony.start == nil {
		fmt.Println("ERROR: invalid data format")
		return nil, ""
	}
	if colony.end == nil {
		fmt.Println("ERROR: invalid data format")
		return nil, ""
	}

	return colony, string(file)
}
