package main

import (
	"fmt"
	"lem/outils"
)

func main() {
	colony, success := lem.Parsing()
	if success == ""{
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
		fmt.Println("ERROR: invalid data format")
		return
	}
	result := lem.AntsName(colony, capacities)

	if success != ""{
		fmt.Println(success)
		fmt.Println("")
	}else{
		return 
	}

	lem.Printing(result, bestGroup)
}
