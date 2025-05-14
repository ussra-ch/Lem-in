package lem

import (
	"fmt"
	"math"
	"sort"
)

func Graph(colony *Colony) [][]string {
	if colony == nil || colony.start == nil || colony.end == nil {
		fmt.Println("Error: Invalid colony data.")
		return [][]string{}
	}
	forbidden := []string{colony.start.name, colony.end.name}
	allpath := [][]string{}
	for _, linkstart := range colony.links[colony.start.name] {
		if linkstart == colony.start.name || linkstart == colony.end.name {
			if linkstart == colony.end.name {
				path := []string{colony.start.name, colony.end.name}
				allpath = append(allpath, path)
			}
			continue
		}
		for _, linkend := range colony.links[colony.end.name] {
			if linkend == colony.start.name || linkend == colony.end.name {
				continue
			}
			path := []string{colony.start.name}
			path = append(path, bfs(linkstart, linkend, colony, forbidden)...)
			path = append(path, colony.end.name)
			allpath = append(allpath, path)
		}
	}
	sort.Slice(allpath, func(i, j int) bool {
		return len(allpath[i]) < len(allpath[j])
	})

	return allpath
}

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

func notval(b []string, groups [][]string) bool {
	if len(b) == 2{
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

func Grouping(colony *Colony) []Paths {
	allpath := Graph(colony)

	allGroups := []Paths{}

	for _, pathA := range allpath {
		var groups [][]string
		group := pathA
		groups = append(groups, group)

		for _, pathB := range allpath {
			if notval(pathB, groups) {
				groups = append(groups, pathB)
			}
		}
		allGroups = append(allGroups, Paths{Path: groups})

	}

	return allGroups
}

func Found(all []Paths,  NumberOfAnts int) [][]string{
	BestTurn := math.MaxInt64
	BestGroup := map[int][][]string{}
	for i := 0; i < len(all); i++{
		group := all[i].Path
		GroupCapacity := PutAnts(group, NumberOfAnts)
		BigTurn := 0
		for j, x := range group{
			s := len(x) + GroupCapacity[j]-1
			if s > BigTurn{
				BigTurn = s
			}
		}
		BestGroup[BigTurn] = group
		if BigTurn < BestTurn{
			BestTurn = BigTurn
		}
	} 

	return BestGroup[BestTurn]

}
