package lem

import (
	"math"
)

/* This function checks the presence of a slice in a matrix */
func IsValidPath(b []string, groups [][]string) bool {
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
func GroupingPaths(colony *Colony) []Paths {
	allpath := GetAllPathsSorted(colony)
	allGroups := []Paths{}
	for r, pathA := range allpath {
		var groups [][]string
		group := pathA
		groups = append(groups, group)
		for j, pathB := range allpath {
			if r != j && IsValidPath(pathB, groups) {
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

	//This loop is to choose the best path to add the new ant
	for NumberOfAnts > 0 {
		BestPathIndex := 0
		for j := 0; j < len(paths); j++{
			if len(paths[j]) + paths_capacity[j] < len(paths[BestPathIndex]) + paths_capacity[BestPathIndex] {
				BestPathIndex = j
			}
		}
		paths_capacity[BestPathIndex] += 1
		NumberOfAnts--
	}
	return paths_capacity
}
