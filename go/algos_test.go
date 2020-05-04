package main

import (
	"math/rand"
	"testing"
)

/*
  Some tests for the different shortest path algorithms
*/

func BenchmarkDijkstra(b *testing.B) {
	// set up a large random sample graph
	// 10000 vertices, 3+ edges per vertex
	L := RandomGraph(10000, 3)
	start := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = Dijkstra(L, start)
	}
}

func BenchmarkBellmanFord(b *testing.B) {
	// set up a large random sample graph
	// 1000 vertices, 3+ edges per vertex
	L := RandomGraph(1000, 3)
	start := 0

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = BellmanFord(L, start)
	}
}

func BenchmarkFloydWarshall(b *testing.B) {
	// set up a large random sample graph
	// 1000 vertices, 3+ edges per vertex
	L := RandomGraph(1000, 3)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = FloydWarshall(L)
	}
}

func TestDijkstraFibonacciHeap(t *testing.T) {
	rand.Seed(1992)
	G := RandomGraph(500, 2)
	start := 0
	dist, prev := DijkstraFibonacci(G, start)
	path := make([]int, 0, G.V)
	path, _ = getPathD(&start, 499, prev, path)
	// fmt.Println(path)
	if dist[269] != 7.0 {
		t.Errorf("Distance to test node 269 incorrect, got %f, want %f", dist[269], 7.0)
	}
	if dist[499] != 4.0 {
		t.Errorf("Distance to test node 499 incorrect, got %f, want %f", dist[499], 4.0)
	}

}

func BenchmarkDijkstraFibonacciHeap(b *testing.B) {
	// set up a large random sample graph
	// 10000 vertices, 3+ edges per vertex
	G := RandomGraph(10000, 3)
	start := 0

	b.ResetTimer()
	for j := 0; j < b.N; j++ {
		_, _ = DijkstraFibonacci(G, start)
	}
}
