package main

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

func don(x []string) bool {
	xy := make(map[string][]string)

	for i := 1; i < len(x); i++ {
		lin := strings.Fields(x[i])
		if len(lin)==3 {
			for _, value := range xy {
				if reflect.DeepEqual(lin[1:], value) {
					return false
				}
			}
		}

		if _, exists := xy[lin[0]]; exists {
			return false
		}

		xy[lin[0]] = lin[1:]
	}

	return true
}

func main() {
	fileName := os.Args[1]
	fil, _ := os.ReadFile(fileName)
	flInfo := strings.Split(string(fil), "\n")
	number := flInfo[0]
	index := 0
	x, err := strconv.Atoi(number)
	if !don(flInfo) {
		return
	}
	fmt.Println("number of int==", x)
	pass := true
	for i := 0; i < len(flInfo); i++ {
		if !pass && flInfo[i] == "##start" {
			fmt.Println("error 2 start")
			return

		} else if flInfo[i] == "##start" {
			index = i + 1
			pass = false
		}
	}
	pass = true
	start := strings.Fields(flInfo[index])
	fmt.Println("strt==", start[0])
	if err != nil {
		fmt.Println("Error")
		return
	}

	for i := 0; i < len(flInfo); i++ {
		if !pass && flInfo[i] == "##end" {
			fmt.Println("error 2 end")
			return

		} else if flInfo[i] == "##end" {
			index = i + 1
			pass = false
		}
	}
	end := strings.Fields(flInfo[index])
	fmt.Println("end==", end[0])
	mapp := make(map[string][]string)
	for i := index + 1; i < len(flInfo); i++ {
		Ways := strings.Split(string(flInfo[i]), "-")

		if len(Ways) == 2 {
			mapp[Ways[0]] = append(mapp[Ways[0]], Ways[1])
		}

	}
	if len(mapp[start[0]]) == 0 {
		fmt.Println("Error:no way for the int")
		return
	}
	fmt.Println(mapp)
}
