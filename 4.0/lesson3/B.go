package main

import (
	"container/heap"
	"fmt"
	"math"
)

//type Vertex struct {
//	index, distance int
//}
//
//type PriorityQueue []*Vertex
//
//func (pq PriorityQueue) Len() int { return len(pq) }
//func (pq PriorityQueue) Less(i, j int) bool {
//	return pq[i].distance < pq[j].distance
//}
//func (pq PriorityQueue) Swap(i, j int) {
//	pq[i], pq[j] = pq[j], pq[i]
//}
//
//func (pq *PriorityQueue) Push(x interface{}) {
//	vertex := x.(*Vertex)
//	*pq = append(*pq, vertex)
//}
//
//func (pq *PriorityQueue) Pop() interface{} {
//	old := *pq
//	n := len(old)
//	vertex := old[n-1]
//	*pq = old[0 : n-1]
//	return vertex
//}

func dijkstraB(graph [][]int, source, target int) ([]int, int) {
	n := len(graph)
	distances := make([]int, n)
	for i := 0; i < n; i++ {
		distances[i] = math.MaxInt32
	}
	distances[source] = 0

	visited := make([]bool, n)
	previous := make([]int, n)

	pq := make(PriorityQueue, 0)
	heap.Push(&pq, &Vertex{index: source, distance: 0})

	for len(pq) > 0 {
		current := heap.Pop(&pq).(*Vertex)

		if current.index == target {

			path := []int{target}
			for path[0] != source {
				path = append([]int{previous[path[0]]}, path...)
			}
			return path, current.distance
		}

		visited[current.index] = true

		for i := 0; i < n; i++ {
			if graph[current.index][i] >= 0 && !visited[i] {
				weight := graph[current.index][i]
				if current.distance+weight < distances[i] {
					distances[i] = current.distance + weight
					previous[i] = current.index
					heap.Push(&pq, &Vertex{index: i, distance: distances[i]})
				}
			}
		}
	}

	return nil, -1
}

func mainB() {
	var n, s, f int
	fmt.Scan(&n, &s, &f)

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			fmt.Scan(&graph[i][j])
		}
	}

	distance, _ := dijkstraB(graph, s-1, f-1)
	if len(distance) == 0 {
		fmt.Println(-1)
		return
	}
	for i := 0; i < len(distance); i++ {
		fmt.Print(distance[i]+1, " ")
	}
	fmt.Println()
}
