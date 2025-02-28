package lem

import "fmt"

func Graph(colony *Colony) [][]string {
	if colony == nil || colony.start == nil || colony.end == nil {
		fmt.Println("Error: Invalid colony data.")
		return [][]string{}
	}
	allpath := [][]string{}
	for _, linkstart := range colony.links[colony.start.name] {
		for _, linkend := range colony.links[colony.end.name] {
			path := []string{colony.start.name}
			path = append(path, bfs(linkstart, linkend, colony)...)
			path = append(path, colony.end.name)

			allpath = append(allpath, path)
		}
	}
	return allpath
}

func bfs(start, end string, colony *Colony) []string {
	queue := [][]string{{start}}
	visited := make(map[string]bool)
	visited[start] = true

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
