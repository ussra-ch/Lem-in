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

	AntIndex  := FindAntLine(lines)
	if AntIndex == -1 {
		fmt.Println("ERROR: invalid data format")
		return nil, ""
	}

	if !ParseAntCount(colony, lines[AntIndex]){
		return nil, ""
	}

	if !ParseLines(colony, lines, AntIndex){
		return nil, ""
	}

	if !ValidateColony(colony){
		return nil, ""
	}

	return colony, string(file)
}

func FindAntLine(lines []string) int {
	for i, line := range lines {
		line = strings.TrimSpace(line)
		if line != "" && line[0] != '#' {
			return i
		}
	}
	return -1
}

func ParseAntCount(colony *Colony, line string) bool {
	ant, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil || ant <= 0 || ant > 10000000 {
		fmt.Println("ERROR: invalid data format")
		return false
	}
	colony.NumAnts = ant
	return true
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

func ParseLines(colony *Colony, lines []string, AntIndex int) bool{
	// end means the End Room
	// start means the Start Room
	var end, start bool
	links := false

	for r := AntIndex + 1; r < len(lines); r++ {
		line := strings.TrimSpace(lines[r])
		if ShouldSkipLine(line) {
			continue
		}
		if line[0] == 'L' {
			return ErrorInvalidFormat()
		}

		room := strings.Fields(line)
		if len(room) == 3 {
			if links {
				return ErrorInvalidFormat()
			}
			if start {
				if colony.start != nil {
					return ErrorInvalidFormat()
				}
				RoomX, RoomY, ok := RoomCoordinatesAndCheck(colony, room)
				if !ok{
					return ErrorInvalidFormat()
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
					return ErrorInvalidFormat()
				}
				RoomX, RoomY, ok := RoomCoordinatesAndCheck(colony, room)
				if !ok{
					return ErrorInvalidFormat()
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
					return ErrorInvalidFormat()
				}
				if (colony.start != nil && room[0] == colony.start.name) || (colony.end != nil && room[0] == colony.end.name) {
					return ErrorInvalidFormat()
				}
				RoomX, RoomY, ok := RoomCoordinatesAndCheck(colony, room)
				if !ok{
					return ErrorInvalidFormat()
				}
				for _, existingRoom := range colony.rooms {
					if RoomX == existingRoom.x && RoomY == existingRoom.y {
						return ErrorInvalidFormat()
					}
				}
				if (colony.start != nil && RoomX == colony.start.x && RoomY == colony.start.y) || (colony.end != nil && RoomX == colony.end.x && RoomY == colony.end.y) {
					return ErrorInvalidFormat()
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
				return ErrorInvalidFormat()
			}
			start = true
		} else if line == "##end" {
			if end || start {
				return ErrorInvalidFormat()
			}
			end = true
		} else {
			links = true
			link := strings.Split(line, "-")
			if len(link) == 2 && link[0] != "" && link[1] != "" {
				if link[0] == link[1]{
					return ErrorInvalidFormat()
				}
				ok1 := IsValidRoom(link[0], colony)
				ok2 := IsValidRoom(link[1], colony)
				if !ok1 || !ok2{
					return ErrorInvalidFormat()
				}
				colony.links[link[0]] = append(colony.links[link[0]], link[1])
				colony.links[link[1]] = append(colony.links[link[1]], link[0])

			} else {
				return ErrorInvalidFormat()
			}
		}

	}
	return true
}

func ValidateColony(colony *Colony) bool {
	if colony.start == nil {
		fmt.Println("ERROR: invalid data format")
		return false
	}
	if colony.end == nil {
		fmt.Println("ERROR: invalid data format")
		return false
	}
	if colony.start.x == colony.end.x && colony.start.y == colony.end.y{
		fmt.Println("ERROR: invalid data format")
		return false
	}

	return true
}

func ShouldSkipLine(line string) bool {
	return line == "" || (line[0] == '#' && line != "##start" && line != "##end")
}

func RoomCoordinatesAndCheck(colony *Colony, room []string) (int, int, bool) {
	x, err := strconv.Atoi(room[1])
	if err != nil {
		return 0, 0, false
	}
	y, err := strconv.Atoi(room[2])
	if err != nil {
		return 0, 0, false
	}
	if _, exists := colony.rooms[room[0]]; exists {
		return 0, 0, false
	}
	return x, y, true
}


func ErrorInvalidFormat() bool {
	fmt.Println("ERROR: invalid data format")
	return false
}
