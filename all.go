package main

import (
	"fmt"
	// lem "lem/outils"
	// lem "lem/outils"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main1() {
	colony, success := Parsing()
	if success == "" {
		return
	}
	allgroup := PathsGrouping(colony)
	bestGroup := ChooseAGroup(allgroup, colony.NumAnts)
	capacities := PathsCapacity(bestGroup, colony.NumAnts)

	count := 0
	for _, x := range capacities {
		if x == 0 {
			count++
		}
	}
	if count == len(capacities) {
		fmt.Println("Invalid input")
		return
	}
	result := make([][]string, len(capacities))
	totalAnts := colony.NumAnts
	antNumber := 1
	assigned := make([]int, len(capacities))
	for antNumber <= totalAnts {
		for i := 0; i < len(capacities); i++ {
			if assigned[i] < capacities[i] {
				result[i] = append(result[i], fmt.Sprintf("L%d", antNumber))
				assigned[i]++
				antNumber++
			}
		}
	}
	PrintAnts(result, bestGroup)
}

/* This function checks the presence of a slice in a matrix */
func IsValidPath(b []string, groups [][]string) bool {
	if len(b) == 2 {
		return false
	}
	set := make(map[string]bool)
	for _, a := range groups {
		for r, val := range a {
			if r != 0 && r != len(a)-1 {
				set[val] = true
			}
		}
	}
	for _, val := range b {
		if set[val] {
			return false
		}
	}
	return true
}

/* This function groups the paths by non-crossing path */
func PathsGrouping(colony *Colony) []Paths {
	allpath := Graph(colony)
	allGroups := []Paths{}

	for _, pathA := range allpath {
		var groups [][]string
		group := pathA
		groups = append(groups, group)

		for _, pathB := range allpath {
			if IsValidPath(pathB, groups) {
				groups = append(groups, pathB)
			}
		}
		allGroups = append(allGroups, Paths{Path: groups})

	}

	return allGroups
}

/* This function finds the perfect group to send the ants through */
func ChooseAGroup(all []Paths, NumberOfAnts int) [][]string {
	BestTurn := math.MaxInt64
	BestGroup := map[int][][]string{}
	for i := 0; i < len(all); i++ {
		group := all[i].Path
		GroupCapacity := PathsCapacity(group, NumberOfAnts)
		BigTurn := 0
		for j, x := range group {
			s := len(x) + GroupCapacity[j] - 1
			if s > BigTurn {
				BigTurn = s
			}
		}
		BestGroup[BigTurn] = group
		if BigTurn < BestTurn {
			BestTurn = BigTurn
		}
	}

	return BestGroup[BestTurn]

}

func PathsCapacity(paths [][]string, NumberOfAnts int) []int {
	paths_capacity := make([]int, len(paths))
	if len(paths) == 0 {
		return paths_capacity
	}
	if len(paths) == 1 {
		paths_capacity[0] = NumberOfAnts
		return paths_capacity
	}

	for i := 0; NumberOfAnts > 0; i++ {
		if i == len(paths)-1 {
			i = 0
		}
		if len(paths[i])+paths_capacity[i] > len(paths[i+1])+paths_capacity[i+1] {
			paths_capacity[i+1] += 1
		} else {
			paths_capacity[i] += 1
		}
		NumberOfAnts--
	}
	return paths_capacity
}

type Room struct {
	name string
	x, y int
}
type Colony struct {
	NumAnts int
	rooms   map[string]*Room
	links   map[string][]string
	start   *Room
	end     *Room
}
type Paths struct {
	Path [][]string
}

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
	linekes := false
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
				return nil, ""
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
				end = false
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
				return nil, ""
			}
			start = true

		} else if line == "##end" {
			if r+1 >= len(lines) {
				fmt.Println("ERROR: invalid data format")
				return nil, ""
			}
			if end {
				fmt.Println("Error")
				return nil, ""
			}
			end = true

		} else {
			linekes = true
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

/*
	This function returns all the possibele paths

from the start to the end
*/
func Graph(colony *Colony) [][]string {
	if colony == nil || colony.start == nil || colony.end == nil {
		fmt.Println("Error: Invalid colony data.")
		return [][]string{}
	}
	//makes start & end forbidden
	forbidden := []string{colony.start.name, colony.end.name}
	allpath := [][]string{}

	for _, StartNeighbor := range colony.links[colony.start.name] {
		if StartNeighbor == colony.start.name || StartNeighbor == colony.end.name {
			if StartNeighbor == colony.end.name {
				path := []string{colony.start.name, colony.end.name}
				allpath = append(allpath, path)
			}
			continue
		}
		for _, EndNeighbor := range colony.links[colony.end.name] {
			if EndNeighbor == colony.start.name || EndNeighbor == colony.end.name {
				continue
			}
			path := []string{colony.start.name}
			path = append(path, bfs(StartNeighbor, EndNeighbor, colony, forbidden)...)
			path = append(path, colony.end.name)
			allpath = append(allpath, path)
		}
	}
	sort.Slice(allpath, func(i, j int) bool {
		return len(allpath[i]) < len(allpath[j])
	})

	return allpath
}

/* This function traverses the graph using BFS algo */
func bfs(start, end string, colony *Colony, forbidden []string) []string {
	queue := [][]string{{start}}
	visited := make(map[string]bool)
	visited[start] = true

	for _, node := range forbidden {
		visited[node] = true
	}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		lastNode := path[len(path)-1]

		if lastNode == end {
			return path
		}

		for _, neighbor := range colony.links[lastNode] {
			if !visited[neighbor] {
				visited[neighbor] = true
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return []string{}
}

func PrintAnts(antsOnPath [][]string, foundPaths [][]string) {
	// fmt.Println(antsOnPath)
	// fmt.Println(foundPaths)
	if len(foundPaths[0]) == 1 {
		count := len(antsOnPath[0])
		for i := 1; count > 0; i++ {
			fmt.Printf("L%d-%s ", i, foundPaths[0][0])
			count--
		}
		fmt.Println()
		return
	}
	maxLen := len(antsOnPath[0])
	for _, v := range antsOnPath {
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}
	res := make([][]string, 1)
	for index, element, stack := 0, 0, 0; index < len(antsOnPath); index++ {
		for resIndex, room := range foundPaths[index] {
			if resIndex == 0 {
				continue
			}
			if element >= len(antsOnPath[index]) {
				break
			}
			if (resIndex-1)+stack >= len(res) {
				res = append(res, []string{})
			}
			res[(resIndex-1)+stack] = append(res[(resIndex-1)+stack], antsOnPath[index][element]+"-"+room)
		}
		if index+1 >= len(antsOnPath) {
			index = -1
			element++
			stack++
		}
		if element >= maxLen {
			break
		}
	}
	for _, line := range res {
		for _, move := range line {
			fmt.Printf("%s ", move)
		}
		fmt.Println()
	}
}











// for i := 0; NumberOfAnts > 0; i++ {
// 	if i == len(paths)-1 {
// 		i = 0
// 	}
// 	if len(paths[i])+paths_capacity[i] > len(paths[i+1])+paths_capacity[i+1] {
// 		paths_capacity[i+1] += 1
// 	} else {
// 		paths_capacity[i] += 1

// 	}
// 	if paths_capacity[i] > paths_capacity[0] {
// 		i = 0
// 	}
// 	NumberOfAnts--
// }