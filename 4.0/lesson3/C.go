package main

import (
	"container/heap"
	"fmt"
	"math"
)

type Edge struct {
	to     int
	weight int
}

func Dijkstra(N, A, B int, roads [][]int) int {
	graph := make([][]Edge, N+1)
	for _, road := range roads {
		a, b, weight := road[0], road[1], road[2]
		graph[a] = append(graph[a], Edge{to: b, weight: weight})
		graph[b] = append(graph[b], Edge{to: a, weight: weight})
	}

	distance := make([]int, N+1)
	for i := range distance {
		distance[i] = math.MaxInt32
	}
	distance[A] = 0

	pq := make(PriorityQueue, 0)
	heap.Init(&pq)
	heap.Push(&pq, &Item{city: A, priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		city := item.city
		dist := item.priority

		if city == B {
			break
		}

		if dist > distance[city] {
			continue
		}

		for _, edge := range graph[city] {
			newDist := dist + edge.weight
			if newDist < distance[edge.to] {
				distance[edge.to] = newDist
				heap.Push(&pq, &Item{city: edge.to, priority: newDist})
			}
		}
	}

	if distance[B] == math.MaxInt32 {
		return -1
	}

	return distance[B]
}

type Item struct {
	city     int
	priority int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) Update(item *Item, priority int) {
	item.priority = priority
	heap.Fix(pq, item.priority)
}
func mainC() {
	var N, K, A, B int
	fmt.Scan(&N, &K)

	roads := make([][]int, K)
	for i := 0; i < K; i++ {
		var ai, bi, li int
		fmt.Scan(&ai, &bi, &li)
		roads[i] = []int{ai, bi, li}
	}

	fmt.Scan(&A, &B)

	shortestDistance := Dijkstra(N, A, B, roads)

	fmt.Println(shortestDistance)
}
