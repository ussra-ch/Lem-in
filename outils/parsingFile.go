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
		// fmt.Println(x)
		x = strings.TrimSpace(x)
		if (x != "" && x[0] == '#') || x == "" {
			continue
		} else {
			AntIndex = i
			break
		}
	}
	// fmt.Println(AntIndex)

	ant, err := strconv.Atoi(strings.TrimSpace(lines[AntIndex]))
	if err != nil || ant <= 0 || ant > 10000000 {
		fmt.Println("ERROR: invalid data format")
		return nil, ""
	}
	colony.NumAnts = ant
	// end means the End Room
	// start means the Start Room
	var end, start bool
	links := false

	for r := AntIndex + 1; r < len(lines); r++ {
		line := strings.TrimSpace(lines[r])
		if line == "" || (line[0] == '#' && line != "##start" && line != "##end") {
			continue
		}
		if line[0] == 'L' {
			fmt.Println("ERROR: invalid data format")
			return nil, ""
		}

		room := strings.Fields(line)
		if len(room) == 3 {
			if links {
				fmt.Println("ERROR: invalid data format")
				return nil, ""
			}
			if start {
				if colony.start != nil {
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
				if _, exists := colony.rooms[room[0]]; exists {
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
				if colony.end != nil {
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
				if _, exists := colony.rooms[room[0]]; exists {
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}
				colony.end = &Room{
					name: room[0],
					x:    RoomX,
					y:    RoomY,
				}
				end = false
				continue
			} else {
				if _, exists := colony.rooms[room[0]]; exists {
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}
				if (colony.start != nil && room[0] == colony.start.name) || (colony.end != nil && room[0] == colony.end.name) {
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
				if (colony.start != nil && RoomX == colony.start.x && RoomY == colony.start.y) || (colony.end != nil && RoomX == colony.end.x && RoomY == colony.end.y) {
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}
				colony.rooms[room[0]] = &Room{
					name: room[0],
					x:    RoomX,
					y:    RoomY,
				}
				continue
			}
		} else if line == "##start" {
			if start || end {
				fmt.Println("ERROR: invalid data format")
				return nil, ""
			}
			start = true
		} else if line == "##end" {
			if end || start {
				fmt.Println("ERROR: invalid data format")
				return nil, ""
			}
			end = true
		} else {
			links = true
			link := strings.Split(line, "-")
			if len(link) == 2 && link[0] != "" && link[1] != "" {
				if link[0] == link[1]{
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}
				ok1 := IsValidRoom(link[0], colony)
				ok2 := IsValidRoom(link[1], colony)
				if !ok1 || !ok2{
					fmt.Println("ERROR: invalid data format")
					return nil, ""
				}
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
	if colony.start.x == colony.end.x && colony.start.y == colony.end.y{
		fmt.Println("ERROR: invalid data format")
		return nil, ""
	}

	return colony, string(file)
}



func IsValidRoom( name string, colony *Colony) bool {
	if _, ok := colony.rooms[name]; ok {
		return true
	}
	if colony.start != nil && name == colony.start.name {
		return true
	}
	if colony.end != nil && name == colony.end.name {
		return true
	}
	return false
}