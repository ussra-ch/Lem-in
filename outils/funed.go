package lem

import (
	"math"
)

/* This function checks the presence of a slice in a matrix */
func IsValidPath(b []string, groups [][]string) bool {
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
func ChooseAGroup(all []Paths,  NumberOfAnts int) [][]string{
	BestTurn := math.MaxInt64
	BestGroup := map[int][][]string{}
	for i := 0; i < len(all); i++{
		group := all[i].Path
		GroupCapacity := PathsCapacity(group, NumberOfAnts)
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