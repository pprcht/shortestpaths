package fibheap

import (
	"testing"
)

/*
  Some tests for the fibonacci heap
*/

func BenchmarkCreateHeap(b *testing.B) {
	Q := NewFibonacciHeap()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		H := NewHeapnode(1.33)
		Q.InsertHeapnode(H)
	}
}

//set up a fibonacci heap with 50 points.
//Put the Fibonacci sequence as node keys
func TestHeap(t *testing.T) {
	Q := NewFibonacciHeap()
	F1 := 1.0
	F2 := 0.0
	Fn := 0.0
	for i := 0; i < 50; i++ {
		if i > 0 {
			Fn = F1 + F2
			F2 = F1
			F1 = Fn
		}
		H := NewHeapnode(Fn)
		Q.InsertHeapnode(H)
	}
	//fmt.Println(Q.min.right.right, Q.min.left, Q.nodes)
	min := Q.min.key
	if min != 0.0 {
		t.Errorf("Key of the minimum node incorrect, got %f, want %f", min, 0.0)
	}
	min = Q.min.left.key
	if min != 1.2586269025e+10 {
		t.Errorf("Key of the last node incorrect, got %f, want %f", min, 1.2586269025e+10)
	}

}

func TestDetachHeapNode(t *testing.T) {
	Q := NewFibonacciHeap()
	F1 := 1.0
	F2 := 0.0
	Fn := 0.0
	for i := 0; i < 10; i++ {
		if i > 0 {
			Fn = F1 + F2
			F2 = F1
			F1 = Fn
		}
		H := NewHeapnode(Fn)
		H.index = i
		Q.InsertHeapnode(H)
	}
	node := Q.min.right
	// fmt.Println("detaching node", node.index)
	// fmt.Println("original neigbours", node.left.index, node.right.index)
	node.DetachHeapnode()
	// fmt.Println("New neigbours after detaching node", node.left.index, node.right.index)
	// fmt.Println("original position of the node was replaced by node", Q.min.right.index)
	min := node.left.index
	if min != node.index {
		t.Errorf("Node incorrectly detached, got %d as neighbour, want node itself", min)
	}
	if Q.min.right.index != 2 {
		t.Errorf("Node incorrectly detached, got %d as neighbour, want %d", Q.min.right.index, 2)
	}

}

func TestHeapCleanup(t *testing.T) {
	Q := NewFibonacciHeap()
	F1 := 1.0
	F2 := 0.0
	Fn := 0.0
	for i := 0; i < 10; i++ {
		if i > 0 {
			Fn = F1 + F2
			F2 = F1
			F1 = Fn
		}
		H := NewHeapnode(Fn)
		H.index = i
		Q.InsertHeapnode(H)
	}
	Q.Cleanup()
	min, _ := Q.Getmin()
	if min != 0.0 {
		t.Errorf("Key of the minimum node incorrect, got %f, want %f", min, 0.0)
	}
	if Q.rootnodes != 2 {
		t.Errorf("Number of remaining root nodes incorrect, got %d, want %d", Q.rootnodes, 2)
	}
}

func TestHeapPopmin(t *testing.T) {
	Q := NewFibonacciHeap()
	F1 := 1.0
	F2 := 0.0
	Fn := 0.0
	for i := 0; i < 10; i++ {
		if i > 0 {
			Fn = F1 + F2
			F2 = F1
			F1 = Fn
		}
		H := NewHeapnode(Fn)
		H.index = i
		Q.InsertHeapnode(H)
	}
	Q.Cleanup()

	Q.Popmin()
	min, _ := Q.Getmin()
	if min != 1.0 {
		t.Errorf("Key of the minimum node incorrect, got %f, want %f", min, 1.0)
	}
}
