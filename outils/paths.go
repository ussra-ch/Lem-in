package lem

import (
	"fmt"
	"sort"
)


/* This function returns all the possibele paths 
from the start to the end */
func Graph(colony *Colony) [][]string {
	if colony == nil || colony.start == nil || colony.end == nil {
		fmt.Println("Error: Invalid colony data.")
		return [][]string{}
	}
	//makes start & end forbidden 
	forbidden := []string{colony.start.name, colony.end.name}
	allpath := [][]string{}

	for _, StartNeighbor := range colony.links[colony.start.name] {
		if StartNeighbor == colony.start.name || StartNeighbor == colony.end.name {
			if StartNeighbor == colony.end.name {
				path := []string{colony.start.name, colony.end.name}
				allpath = append(allpath, path)
			}
			continue
		}
		for _, EndNeighbor := range colony.links[colony.end.name] {
			if EndNeighbor == colony.start.name || EndNeighbor == colony.end.name {
				continue
			}
			path := []string{colony.start.name}
			path = append(path, bfs(StartNeighbor, EndNeighbor, colony, forbidden)...)
			path = append(path, colony.end.name)
			allpath = append(allpath, path)
		}
	}
	sort.Slice(allpath, func(i, j int) bool {
		return len(allpath[i]) < len(allpath[j])
	})

	return allpath
}


/* This function traverses the graph using BFS algo */
func bfs(start, end string, colony *Colony, forbidden []string) []string {
	queue := [][]string{{start}}
	visited := make(map[string]bool)
	visited[start] = true

	for _, node := range forbidden {
		visited[node] = true
	}
	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		lastNode := path[len(path)-1]

		if lastNode == end {
			return path
		}

		for _, neighbor := range colony.links[lastNode] {
			if !visited[neighbor] {
				visited[neighbor] = true
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	return []string{}
}
