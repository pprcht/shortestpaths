/* This file contains some routines to set up a simple graph.
As they only serve as an example there is nothing fancy
about them and could probably be written more efficiently.
Also the graph dimension must not be too large. */

package main

import "math/rand"

//Graph is an object containing a set of vertices and edges
type Graph struct {
	V        int         // number of vertices (order of the graph)
	E        int         // number of edges (size of the graph)
	Nmat     [][]int     // neighbour matrix (adjacency matrix)
	Emat     [][]float32 // edge matrix (edge weights)
	directed bool        // is the graph a directed graph?
}

func newGraph() *Graph {
	G := new(Graph)
	G.V = 0
	G.E = 0
	G.directed = false
	return G
}

// set total number of vertices
func (G *Graph) setOrder(i int) {
	G.V = i
	a := make([][]int, G.V)
	for i := range a {
		a[i] = make([]int, G.V)
	}
	G.Nmat = a
	b := make([][]float32, G.V)
	for i := range b {
		b[i] = make([]float32, G.V)
	}
	G.Emat = b
}

// add a new edge between two vertices,
// but only if the vertices are within the order of G
// Each edge requires a value (edge length)
// Setting up parallel edges is avoided.
func (G *Graph) addEdge(v1, v2 int, l float64) {
	var e []int
	if v1 < G.V && v2 < G.V {
		if G.directed {
			e = append(e, v1, v2)
		} else { // if the graph is not directed, sort the edge
			e = edge(v1, v2)
		}
		if !G.checkEdges(e) {
			G.Nmat[e[0]][e[1]] = 1
			G.Emat[e[0]][e[1]] = float32(l)
			if !G.directed { // for undirected graphs the matrices are symmetric
				G.Nmat[e[1]][e[0]] = 1
				G.Emat[e[1]][e[0]] = float32(l)
			}
			G.E++ //update size of G
		}
	} else {
		// err := fmt.Errorf("Vertex does not exist")
		// return err
	}
}

// run a check on a graph: is a given edge part of the graph?
func (G *Graph) checkEdges(edge []int) bool {
	v1 := edge[0]
	v2 := edge[1]
	if G.Nmat[v1][v2] == 1 {
		return true
	}
	return false
}

// delete an edge between two vertices (if it is present) from the graph
func (G *Graph) delEdge(v1, v2 int) {
	G.Nmat[v1][v2] = 0
	G.Emat[v1][v2] = 0
	if !G.directed {
		G.Nmat[v2][v1] = 0
		G.Emat[v2][v1] = 0
	}
}

// quickly convert a pair of two vertices into
// an edge (in the correct order)
func edge(v1, v2 int) []int {
	var e []int
	if v1 <= v2 {
		e = append(e, v1, v2)
	} else {
		e = append(e, v2, v1)
	}
	return e
}

// Get the weight (length) of an edge
func (G *Graph) getWeight(v1, v2 int) float32 {
	return G.Emat[v1][v2]
}

// Get the degree of a vertex (i.e, the number of its connected neighbours)
// it is equal to the sum of row v of the adjacency matrix
func (G *Graph) degree(v int) int {
	deg := 0
	for _, k := range G.Nmat[v] {
		if k == 1 {
			deg++
		}
	}
	return deg
}

// Get a list of the connected neighbours of a vertex
// also outputs the degree of the vertex
func (G *Graph) neighbours(v int) ([]int, int) {
	var nei []int
	deg := 0
	/* this should work for directed and undirected graphs
	since G.Nmat[v] is the part of the adjacency matrix
	belonging to a single vertex v */
	for i, k := range G.Nmat[v] {
		if k == 1 {
			nei = append(nei, i)
			deg++
		}
	}
	return nei, deg
}

// Nlist is a function that returns a list of all neighbours for each vertex
func (G *Graph) Nlist() [][]int {
	neigh := make([][]int, G.V)
	for i := 0; i < G.V; i++ {
		neigh[i], _ = G.neighbours(i) // get the neighbour lists for easy access
	}
	return neigh
}

// disconnect a vertex from the graph
// (i.e., remove all its edges)
func (G *Graph) disconnectVert(v int) {
	nei, _ := G.neighbours(v)
	for _, k := range nei {
		G.delEdge(v, k)
	}
}

// set up the example graph depicted in assets/graph_1
// with the same edge weight for all edges
func (G *Graph) example1() {
	G.setOrder(16)
	l := 1.0 // default edge weight for all edges
	G.addEdge(0, 1, l)
	G.addEdge(0, 4, l)
	G.addEdge(1, 2, l)
	G.addEdge(2, 3, l)
	G.addEdge(2, 7, l)
	G.addEdge(3, 4, l)
	G.addEdge(3, 5, l)
	G.addEdge(4, 8, l)
	G.addEdge(5, 6, l)
	G.addEdge(5, 9, l)
	G.addEdge(5, 10, l)
	G.addEdge(6, 7, l)
	G.addEdge(6, 14, l)
	G.addEdge(8, 9, l)
	G.addEdge(10, 11, l)
	G.addEdge(10, 12, l)
	G.addEdge(11, 12, l)
	G.addEdge(11, 13, l)
	G.addEdge(12, 14, l)
	G.addEdge(13, 14, l)
	G.addEdge(14, 15, l)
}

// set up the example graph depicted in assets/graph_2
// with different weights for all edges
func (G *Graph) example2() {
	G.setOrder(16)
	G.addEdge(0, 1, 1.85)
	G.addEdge(0, 4, 1.36)
	G.addEdge(1, 2, 1.51)
	G.addEdge(2, 3, 2.14)
	G.addEdge(2, 7, 1.59)
	G.addEdge(3, 4, 0.55)
	G.addEdge(3, 5, 0.80)
	G.addEdge(4, 8, 0.91)
	G.addEdge(5, 6, 1.12)
	G.addEdge(5, 9, 1.05)
	G.addEdge(5, 10, 1.12)
	G.addEdge(6, 7, 0.92)
	G.addEdge(6, 14, 1.76)
	G.addEdge(8, 9, 0.78)
	G.addEdge(10, 11, 1.00)
	G.addEdge(10, 12, 0.50)
	G.addEdge(11, 12, 0.45)
	G.addEdge(11, 13, 1.87)
	G.addEdge(12, 14, 1.27)
	G.addEdge(13, 14, 1.64)
	G.addEdge(14, 15, 1.12)
}

//RandomGraph is used to generate a random graph with NV vertices
//each vertex will randomly get NE edges to other vertices
//For simplicity all edges will get the same weight (= 1.0)
func RandomGraph(NV, ne int) *Graph {
	G := newGraph()
	G.setOrder(NV)

	var k int
	var r int
	for i := 0; i < NV; i++ {
		k = 0
		for k < ne {
			r = rand.Intn(NV)
			if r != i {
				G.addEdge(i, r, 1.00)
				k++
			}
		}

	}

	return G
}
