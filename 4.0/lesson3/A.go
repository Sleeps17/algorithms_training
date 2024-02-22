package main

import (
	"fmt"
	"math"
)

func dijkstra(graph [][]int, source, destination int) int {
	N := len(graph)

	dist := make([]int, N)
	for i := range dist {
		dist[i] = math.MaxInt32
	}

	dist[source] = 0

	visited := make([]bool, N)

	for count := 0; count < N-1; count++ {
		u := minDistance(dist, visited)

		visited[u] = true

		for v := 0; v < N; v++ {
			if !visited[v] && graph[u][v] != -1 && dist[u]+graph[u][v] < dist[v] {
				dist[v] = dist[u] + graph[u][v]
			}
		}
	}

	return dist[destination]
}

func minDistance(dist []int, visited []bool) int {
	min := math.MaxInt32
	minIndex := -1

	for v, d := range dist {
		if !visited[v] && d <= min {
			min = d
			minIndex = v
		}
	}

	return minIndex
}

func mainA() {
	var N, S, F int
	fmt.Scan(&N, &S, &F)

	graph := make([][]int, N)
	for i := range graph {
		graph[i] = make([]int, N)
		for j := range graph[i] {
			fmt.Scan(&graph[i][j])
		}
	}

	distance := dijkstra(graph, S-1, F-1)

	if distance == math.MaxInt32 {
		fmt.Println(-1)
	} else {
		fmt.Println(distance)
	}
}
