package lem

import (
	"fmt"
	"sort"
)


/* This function returns all the possibele paths 
from the start to the end */
func GetAllPathsSorted(colony *Colony) [][]string {
	// Validate input: colony and its start/end must be non-nil
	if colony == nil || colony.start == nil || colony.end == nil {
		fmt.Println("Error: Invalid colony data.")
		return [][]string{}
	}
	// "forbidden" rooms should not appear in the middle of any path
	forbidden := []string{colony.start.name, colony.end.name}
	allpath := [][]string{}

	// Iterate over neighbors of the start room
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
	// Sort all collected paths by increasing length (shortest first)
	sort.Slice(allpath, func(i, j int) bool {
		return len(allpath[i]) < len(allpath[j])
	})

	return allpath
}


/* This function traverses the graph using BFS algo */
func bfs(startRoom, endRoom string, colony *Colony, forbidden []string) []string {
	// Initialize the queue with a single-element path starting at startRoom
	queue := [][]string{{startRoom}}
	// Track visited rooms to avoid revisiting and infinite loops
	visited := make(map[string]bool)
	visited[startRoom] = true

	for _, node := range forbidden {
		visited[node] = true
	}
	for len(queue) > 0 {
		// Dequeue the first path in the queue
		path := queue[0]
		queue = queue[1:]
		lastNode := path[len(path)-1]

		// If we've reached the end room, return this path as the result
		if lastNode == endRoom {
			return path
		}

		// Otherwise, explore all unvisited neighbors of lastNode
		for _, neighbor := range colony.links[lastNode] {
			if !visited[neighbor] {
				visited[neighbor] = true
				// Create a new path by copying the current path and appending the neighbor
				newPath := append([]string{}, path...)
				newPath = append(newPath, neighbor)
				queue = append(queue, newPath)
			}
		}
	}
	// If the queue is exhausted without finding endRoom, return empty slice (no path)
	return []string{}
}
