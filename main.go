package main

import (
	"fmt"

	"lem/lem"
)

func main() {
	colony, success := lem.Parsing()
	if success == ""{
		return
	}
	allgroup:=lem.Grouping(colony)
	// fmt.Println(allgroup)
	bestGroup:=lem.Found(allgroup, colony.NumAnts)
	// fmt.Println(bestGroup)
	capacities := lem.PutAnts(bestGroup, colony.NumAnts)
	
	count := 0
	for _, x := range capacities {
		if x == 0 {
			count++
		}
	}
	if count == len(capacities) {
		fmt.Println("No available paths")
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

	// for i := 0; i < len(bestGroup); i++{
	// 	if capacities[i] > 0 && antNumber<totalAnts{
	// 		fmt.Printf("L%d", antNumber)
	// 		capacities[i] --
	// 		antNumber++
	// 		totalAnts--
	// 	}
	// }
	lem.PrintAnts(result, bestGroup)
}
