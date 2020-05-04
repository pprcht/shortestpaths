/*
This file contains routines to run a shortest-path
search using Dijkstra's algorithm.
*/
package main

import (
	"fmt"
	"math"

	"fibheap"
)

// a wrapper for the example
func exampleDijkstraFibonacci(G *Graph, start, end int) {
	// run the algorithm. It will yield all the shortest distances
	// from the start node to all other vertices.
	dist, prev := DijkstraFibonacci(G, start)
	// recustruct the shortest path between the two points
	fmt.Println("shortest path from vertex", start, "to vertex", end, ":")
	if end >= len(prev) || math.IsInf(dist[end], 0) {
		err := fmt.Errorf("the selected vertex is not connected to the start point")
		fmt.Println(err)
	} else {
		path := make([]int, 0, G.V)
		path, _ = getPathD(&start, end, prev, path)
		fmt.Println(path)

		fmt.Println("with a total path length of", dist[end])
	}
}

/*
DijkstraFibonacci is the implementation of Dijkstra's algorithm using
a Fibonacci heap. Should show better asymptotic behaviour
than the regular implementation */
func DijkstraFibonacci(G *Graph, start int) ([]float64, []int) {
	// Initialize the distances.
	// I.e., this is the total distance from the source to any given point
	dist := make([]float64, G.V)
	// Initialize the predecessors.
	// I.e., this is the predecessor for any given point in the path from the source
	prev := make([]int, G.V)
	// removed := make([]bool, G.V)
	// Initialize neighbour lists for all vertices (for easy access later).
	neigh := G.Nlist() // get the neighbour lists for easy access
	// Initialize a list of heap nodes.
	HL := make([]*fibheap.Heapnode, G.V)
	Q := fibheap.NewFibonacciHeap()
	// Initialize data
	for i := 0; i < G.V; i++ {
		infty := math.Inf(0)
		if i == start {
			HL[i] = fibheap.NewHeapnode(0.0)
			prev[i] = start
		} else {
			HL[i] = fibheap.NewHeapnode(infty)
			prev[i] = -1
		}
		// removed[i] = false
		HL[i].SetIndex(i) //index MUST correspond to positon in HL
		Q.InsertHeapnode(HL[i])
	}
	/* 	fmt.Println(Q)
	   	fmt.Println(Q.GetNodes(), Q.GetRootNodes()) */

	var newdist float64
	//while there are still nodes in the heap run the algo
	for Q.GetNodes() > 0 {
		ukey, u := Q.Getmin()

		// removed[u] = true
		for _, v := range neigh[u] {
			// except vertex v to avoid walking back
			if prev[u] != v {
				// if removed[v] == false {
				newdist = ukey + float64(G.getWeight(u, v))
				// is the disteance from u to v smaller than the old distance saved for v?
				if newdist < HL[v].GetKey() {
					// if so, update distance and predecessors for v
					Q.UpdateKey(HL[v], newdist)
					prev[v] = u //keep track of the path
				}

			}
		}
		_, _ = Q.Popmin()
	}

	//finally (only for this implementation) write the distances to slice dist
	for i := 0; i < G.V; i++ {
		dist[i] = HL[i].GetKey()
	}

	return dist, prev
}
