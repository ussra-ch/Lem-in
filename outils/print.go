package lem

import "fmt"

func Printing(antsOnPath [][]string, foundPaths [][]string) {
	maxLen := len(antsOnPath[0])
	for _, v := range antsOnPath {
		if len(v) > maxLen {
			maxLen = len(v)
		}
	}
	res := make([][]string, 1)
	for PathIndex, TheAnt, stack := 0, 0, 0; PathIndex < len(antsOnPath); PathIndex++ {
		for resIndex, room := range foundPaths[PathIndex] {
			//check if i'm in the start room, if yes i skip it
			if resIndex == 0 {
				continue
			}
			if TheAnt == len(antsOnPath[PathIndex]) {
				break
			}
			if (resIndex-1)+stack >= len(res) {
				res = append(res, []string{})
			}
			res[(resIndex-1)+stack] = append(res[(resIndex-1)+stack], antsOnPath[PathIndex][TheAnt]+"-"+room)
		}
		if PathIndex+1 >= len(antsOnPath) {
			PathIndex = -1
			TheAnt++
			stack++
		}
		if TheAnt >= maxLen {
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

func AntsName(colony *Colony, capacities []int)[][]string{
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
	return result
}