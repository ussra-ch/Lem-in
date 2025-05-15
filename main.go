package main

import (
	"fmt"

	"lem/outils"
)

func main() {
	colony, success := lem.Parsing()
	if success != ""{
		fmt.Println(success)
		fmt.Println("")
	}else{
		return
	}
	allgroup:=lem.GroupingPaths(colony)

	bestGroup:=lem.ChooseAGroup(allgroup, colony.NumAnts)
	capacities := lem.PathsCapacity(bestGroup, colony.NumAnts)
	
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
	lem.Printing(result, bestGroup)
}
