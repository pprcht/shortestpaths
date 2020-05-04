// Dijkstra's Algorithm written in Go
package main

import (
	"fmt"
)

func main() {
	// set up a sample graph
	G := newGraph()

	G.example1() // all edges have the same weight
	// three different paths from node 0 to 13 are possible

	// G.example2() // edges have different weights
	//there is just one best path from 0 to 13

	fmt.Println("Vertices", G.V)
	fmt.Println("Edges", G.E)
	fmt.Println()

	// start & end point
	start := 0
	end := 13

	//serach the shortest path using Dijkstra's algorithm
	fmt.Println("Shortest path from", start, "to", end, "using Dijkstra's algorithm:")
	exampleDijkstra(G, start, end)
	fmt.Println()

	//serach the shortest path using the Bellman-Ford algorithm
	fmt.Println("Shortest path from", start, "to", end, "using the Bellman-Ford algorithm:")
	exampleBellmanFord(G, start, end)
	fmt.Println()

	//serach the shortest path using the Floyd-Warshall algorithm
	fmt.Println("Shortest path from", start, "to", end, "using the Floyd-Warshall algorithm:")
	exampleFloydWarshall(G, start, end)
	fmt.Println()

	//serach the shortest path using Dijkstra's algorithm with a Fibonacci heap implementation
	fmt.Println("Shortest path from", start, "to", end, "using Dijkstra's algorithm (Fibonacci heap):")
	exampleDijkstraFibonacci(G, start, end)
	fmt.Println()

}
