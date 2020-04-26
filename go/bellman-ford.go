/*
This file contains routines to run a shortest-path
search using the Bellman-Ford algorithm.
It is very similar to Dijkstra's algorithm but
can also handle graphs with negative edge weights
(opposed to Dijkstra's algo), making it more versatile.
It will only terminate incorrectly if there is a
negative cycle, i.e., an cycle with all negative
edge weights.
*/
package main

import (
	"fmt"
	"math"
)

// a wrapper for the example
func exampleBellmanFord(G *graph, start, end int) {
	// run the algorithm. It will yield all the shortest distances
	// from the start node to all other vertices.
	dist, prev := BellmanFord(G, start)
	// recustruct the shortest path between the two points
	fmt.Println("shortest path from vertex", start, "to vertex", end, ":")
	if math.IsInf(dist[end], 0) {
		err := fmt.Errorf("the selected vertex is not connected to the start point")
		fmt.Println(err)
	} else {
		path := make([]int, 0, G.V)
		// here we can use the same path reconstruction routine as for Dijkstra's algorithm
		path, _ = getPathD(&start, end, prev, path)
		fmt.Println(path)

		fmt.Println("with a total path length of", dist[end])
	}
}

// BellmanFord is the routine containing the setup and the algorithm
// for finding the shortest path to ALL vertices from a given
// starting point.
func BellmanFord(G *graph, start int) ([]float64, []int) {
	// Initialize the distances.
	// I.e., this is the total distance from the source to any given point
	dist := make([]float64, G.V)
	// Initialize the predecessors.
	// I.e., this is the predecessor for any given point in the path from the source
	prev := make([]int, G.V)
	// Initialize explicit list of edges (obtained from adjacency matrix).
	var edges [][]int
	// Initialize data
	for i := 0; i < G.V; i++ {
		dist[i] = math.Inf(0)      // set distance to vertex i to "infinity"
		prev[i] = -1               // set predecessor of vertex i to "undefined"
		for j := 0; j < G.V; j++ { // it has to be all edge combinations, i.e., both (u,v) and (v,u)
			if G.Nmat[i][j] == 1 {
				edges = append(edges, []int{i, j})
			}
		}
	}
	dist[start] = 0.0   // set distance of the source vertex to 0
	prev[start] = start // set the predecessor of the source to itself

	// repeated relaxation of edges (i => n-1 times)
	for i := 1; i < G.V; i++ {
		for _, e := range edges {
			u := e[0]
			v := e[1]
			newdist := dist[u] + float64(G.Emat[u][v])
			if newdist < dist[v] {
				dist[v] = newdist
				prev[v] = u
			}
		}
	}

	// check for negative cycles (would be the n-th iteration of the for-loop above)
	for _, e := range edges {
		u := e[0]
		v := e[1]
		newdist := dist[u] + float64(G.Emat[u][v])
		if newdist < dist[v] {
			err := fmt.Errorf("warning: the graph contains a negativ-weight cycle")
			fmt.Println(err)
		}
	}

	return dist, prev
}
