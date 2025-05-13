package lem

import "fmt"

func PutAnts(paths [][]string, NumberOfAnts int) []int {
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
