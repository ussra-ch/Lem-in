package lem

import "fmt"

func Printing(antsOnPath [][]string, foundPaths [][]string) {
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
	for PathIndex, TheAnt, stack := 0, 0, 0; PathIndex < len(antsOnPath); PathIndex++ {
		for resIndex, room := range foundPaths[PathIndex] {
			if resIndex == 0 {
				continue
			}
			if TheAnt >= len(antsOnPath[PathIndex]) {
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

func Print(capacities []int, bestGroup [][]string){
	
}
