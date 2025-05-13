package main

import (
	"fmt"

	"lem/lem"
)

func main() {
	colony, success := lem.Parsing()
	if success !=""{
		fmt.Println(success)
	}
	allgroup:=lem.Grouping(colony)
	bestGroup:=lem.Fined(allgroup)
	fmt.Println(bestGroup)
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
		for i := 0; i < len(capacities) && antNumber <= totalAnts; i++ {
			if assigned[i] < capacities[i] {
				result[i] = append(result[i], fmt.Sprintf("L%d", antNumber))
				assigned[i]++
				antNumber++
			}
		}
	}
	lem.PrintAnts(result, lem.Graph(colony))
}
