/*
This file contains routines to run a shortest-path
search using Dijkstra's algorithm.
*/
package main

import (
	"fmt"
	"math"
)

// a wrapper for the example
func exampleDijkstra(G *Graph, start, end int) {
	// run the algorithm. It will yield all the shortest distances
	// from the start node to all other vertices.
	dist, prev := Dijkstra(G, start)
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

// Dijkstra is the routine containing the setup and the algorithm
// for finding the shortest path to ALL vertices from a given
// starting point.
func Dijkstra(G *Graph, start int) ([]float64, []int) {
	// Initialize the distances.
	// I.e., this is the total distance from the source to any given point
	dist := make([]float64, G.V)
	// Initialize the predecessors.
	// I.e., this is the predecessor for any given point in the path from the source
	prev := make([]int, G.V)
	// Initialize the list of unvisited vertices.
	Q := make([]int, G.V)
	// Initialize neighbour lists for all vertices (for easy access later).
	neigh := G.Nlist() // get the neighbour lists for easy access
	// Initialize data
	for i := 0; i < G.V; i++ {
		dist[i] = math.Inf(0) // set distance to vertex i to "infinity"
		prev[i] = -1          // set predecessor of vertex i to "undefined"
		Q[i] = i              // add the vertex to the queue
	}
	dist[start] = 0.0   // set distance of the source vertex to 0
	prev[start] = start // set the predecessor of the source to itself

	// As long as there are unvisited vertices in Q, we continue
	for len(Q) >= 1 {
		// get the vertex u with the smallest distance of vertices in the set Q
		u, i, err := minQ(Q, dist)
		// resize Q, important: order in Q is *not* maintained, i is overwritten
		Q[i] = Q[len(Q)-1]
		Q[len(Q)-1] = 0
		Q = Q[:len(Q)-1]

		if err == nil {
			// loop over all neighbours of u
			for _, v := range neigh[u] {
				// except vertex v to avoid walking back
				if prev[u] != v {
					newdist := dist[u] + float64(G.getWeight(u, v))
					// is the disteance from u to v smaller than the old distance saved for v?
					if newdist < dist[v] {
						// if so, update distance and predecessors for v
						dist[v] = newdist
						prev[v] = u
					}
				}
			}
		}
		// if we are only interested in the path between the start and end nodes
		// we could already exit the loop after we have "visited" the end vertex
	}

	return dist, prev
}

// minQ gets the vertex with the smallest distance to the source from Q
func minQ(Q []int, dist []float64) (int, int, error) {
	var vertex int
	var pos int
	var err error
	dref := math.Inf(0)
	for i, k := range Q {
		d := dist[k]
		if d < dref {
			vertex = k
			pos = i
			dref = d
		}
	}
	if math.IsInf(dref, 0) {
		err = fmt.Errorf("Warning: disconnected graph")
	}
	return vertex, pos, err
}

// analyze the predecessors (recursion) and construct the shortest possible path
func getPathD(start *int, pos int, prev []int, path []int) ([]int, error) {
	var err error
	if pos != *start {
		path = append([]int{pos}, path...) //prepend slice
		path, err = getPathD(start, prev[pos], prev, path)
	} else {
		path = append([]int{pos}, path...)
	}
	return path, err
}
