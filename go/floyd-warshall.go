/*
This file contains routines for performing the
shortest-path search using a Floyd-Warshall algorithm.
The Floyd-Warshall algorithm yields all-pair shortest-paths
*/
package main

import (
	"fmt"
	"math"
)

// a wrapper for the actual routine call
func exampleFloydWarshall(G *graph, start, end int) {

	dist, prev := FloydWarshall(G)
	fmt.Println("shortest path from vertex", start, "to vertex", end, ":")
	if math.IsInf(dist[start][end], 0) {
		err := fmt.Errorf("the selected vertex is not connected to the start point")
		fmt.Println(err)
	} else {
		path := getPathFW(prev, start, end)
		fmt.Println(path)

		fmt.Println("with a total path length of", dist[start][end])
	}

}

// FloydWarshall is the implementation of the Floyd-Warshall algorithm
func FloydWarshall(G *graph) ([][]float64, [][]int) {

	// The algorithm requires a V x V distance matrix
	// with all the edge weights.
	// This matrix is set up from the edge weight matrix G.Emat
	// but elements not belonging to an edge get initialized with +Inf
	// Furthermore, for later path reconstruction we need a
	// neighbour matrix prev. We can construct it from the
	// adjacency matrix G.Nmat as such:
	dist := make([][]float64, G.V)
	prev := make([][]int, G.V)
	for i := range dist {
		dist[i] = make([]float64, G.V)
		prev[i] = make([]int, G.V)
		for j := range dist[i] {
			if G.Nmat[i][j] == 1 {
				dist[i][j] = float64(G.Emat[i][j])
				prev[i][j] = j
			} else {
				dist[i][j] = math.Inf(0)
				prev[i][j] = -1
			}
		}
	}

	// The algorithm is based on the following assumption:
	// If a shortest path from vertex u to vertex v runns through a thrid
	// vertex w, then the paths u-to-w and w-to-v are already minimal.
	// Hence, the shorest paths are constructed by searching all path
	// that run over an additional intermediate point k
	var kdist float64
	for k := 0; k < G.V; k++ {
		for i := 0; i < G.V; i++ {
			for j := 0; j < G.V; j++ {
				kdist = dist[i][k] + dist[k][j]
				if dist[i][j] > kdist { // if the path from i to j runs over k, update
					dist[i][j] = kdist
					prev[i][j] = prev[i][k]
				}
			}
		}
	}

	return dist, prev
}

// reconstruct the path
func getPathFW(prev [][]int, u, v int) []int {
	var path []int
	// if the vertices u and v are not connected prev[u][v] will be 0
	if prev[u][v] == -1 {
		return path
	}
	path = append(path, u)
	for u != v {
		u = prev[u][v]
		path = append(path, u)
	}
	return path
}
