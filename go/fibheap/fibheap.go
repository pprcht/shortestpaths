/*
Package fibheap includes an implementation of a fibonacci heap priority queue.
A fibonacci heap stores a list of trees. Each tree is a heap.
In general the heap condition is: the key of a node is <= the key of its children.
The fibonacci heap allows the trees to be any shape (e.g. opposed to the
binomial heap) */
package fibheap

import "fmt"

//Heapnode this is the struct for a single node in the heap
type Heapnode struct {
	key      float64 // the key of the node
	index    int
	tag      bool      // is the node tagged/marked?
	children int       // number of children, also called the degree of the node
	parent   *Heapnode // pointer to the parent node (if there is any)
	child    *Heapnode // pointer to the first child node
	left     *Heapnode // pointer to the left silbling node
	right    *Heapnode // pointer to the right silbling node
}

// NewHeapnode is used to initialize a new node for the heap
func NewHeapnode(k float64) *Heapnode {
	H := new(Heapnode)
	H.key = k
	H.children = 0
	H.parent = nil
	H.child = nil
	H.left = H
	H.right = H
	H.tag = false
	return H
}

// Fheap this is the struct containing the heap
type Fheap struct {
	nodes     int //total number of nodes in the heap
	rootnodes int
	min       *Heapnode //pointer to the current node with the smallest key
}

// NewFibonacciHeap is used to set up an empty heap
func NewFibonacciHeap() *Fheap {
	Q := new(Fheap)
	Q.nodes = 0     //at the beginning there are no nodes in the heap
	Q.rootnodes = 0 //naturally then there are also no rootnodes
	Q.min = nil     //and the pointer to the min node is nil
	return Q
}

// InsertHeapnode a fresh new heap node into the heap. It is put into the root
func (Q *Fheap) InsertHeapnode(H *Heapnode) {
	min := Q.min
	// if the Heap is not empty, add the new node next to the min node
	if min != nil {
		tmp := min.left
		min.left = H
		H.right = min
		H.left = tmp
		tmp.right = H
	}
	//depending on the key of the new heap node (or if no min node is set) update min
	if min == nil || H.key <= min.key {
		Q.min = H
	}
	// don't forget to update the number of nodes
	Q.nodes++
	Q.rootnodes++
}

// DetachHeapnode removes a heap node from its silblings and parent.
// The node will keep its children, however.
func (H *Heapnode) DetachHeapnode() {
	// if the node has silblings, connect them
	if H.right != H {
		H.left.right = H.right
		H.right.left = H.left
	}
	// if there is a parent to node H, we have to update it accordingly
	if H.parent != nil {
		if H.parent.children > 1 {
			H.parent.child = H.left
		} else {
			H.parent.child = nil
		}
		H.parent.children--
	}
	H.right = H
	H.left = H
	H.parent = nil
}

// AddChild adds a child heap node C to parent node H
func (H *Heapnode) AddChild(C *Heapnode) {
	//if the parent node H has no children this is easy
	if H.children < 1 {
		H.child = C
		C.right = C
		C.left = C
	} else { //otherwise update silblings of C
		tmp := H.child.left
		H.child.left = C
		C.right = H.child
		C.left = tmp
		tmp.right = C
	}
	//include the parent pointer C->H
	C.parent = H
	//update number of children of H
	H.children++
}

// LinkHeaps is used to disconnect the fist given
// Heapnode and add it as child to the second
func (H *Heapnode) LinkHeaps(H1 *Heapnode) {
	//First detach H1 from its (former) silblings
	H1.DetachHeapnode()
	//then add H1 as a new child to H
	H.AddChild(H1)
	H1.tag = false
}

// AddRootNode inserts a heap node as into the root nodes.
// This is the same as the InsertHeapnode function, except that the total
// number of roots does not change
func (Q *Fheap) AddRootNode(H *Heapnode) {
	min := Q.min
	// if the Heap is not empty, add the new node next to the min node
	if min != nil {
		tmp := min.left
		min.left = H
		H.right = min
		H.left = tmp
		tmp.right = H
		if H.key < min.key {
			Q.min = H
		}
	} else {
		//depending on the key of the new heap node (or if no min node is set) update min
		// if min == nil || H.key <= min.key {
		Q.min = H
		H.left = H
		H.right = H
	}
	H.tag = false
	H.parent = nil
	Q.rootnodes++
}

// LinkAndDetach is used to disconnect the
// Heapnodes and add one as child to the other
func (H *Heapnode) LinkAndDetach(H1 *Heapnode) {
	//BOTH nodes are detached from their (former) silblings
	H.DetachHeapnode()
	H1.DetachHeapnode()
	//then add H1 as a new child to H
	H.AddChild(H1)
	H1.tag = false
}

/*Cleanup is the most important function of the heap.
There must not be two trees in the Fheap heap that
are silblings and have the same degree (number of children)
This function is used to sort the heap accordingly */
func (Q *Fheap) Cleanup() {
	// Q.nodes == 1 is the trivial case
	if Q.nodes > 1 {
		/* 	initialize the root slice with a single element
		rootslice is a list of Heapnode pointers */
		rootslice := []*Heapnode{nil}
		/* if we call the cleanup Q.min must not necessarily
		be the min node, it is updated after the tree merges */
		tmp := Q.min
		tmpnext := Q.min.left
		b := Q.rootnodes

		//clear Q.min so that during the cleanup process we get the correct list of root nodes
		Q.min = nil
		Q.rootnodes = 0

		var k int
		for i := 0; i < b; i++ {
			tmpnext = tmp.left
			k = tmp.children
			for {
				// if the root slice is not large enough append nil pointers until it is
				for len(rootslice) <= k {
					rootslice = append(rootslice, nil)
				}
				// if we can just put the node into rootslice, or if the node is
				// already at the correct position, go to the next outer loop iteration
				if rootslice[k] == nil || rootslice[k] == tmp {
					// tmp.DetachHeapnode()
					rootslice[k] = tmp
					break
				} else {
					/* if there is already another node with degree k in the slice
					add them together, but maintain the heap condition! */
					if tmp.key > rootslice[k].key {
						rootslice[k].LinkAndDetach(tmp)
						tmp = rootslice[k]
					} else {
						tmp.LinkAndDetach(rootslice[k])
					}
					rootslice[k] = nil //don't forget to clean up the rootslice
					k++                //by linking the nodes, k of tmp grew by 1
				}
			}
			// finally update tmp
			tmp = tmpnext
		}
		//afterwards update Q.min
		for i := 0; i < len(rootslice); i++ {
			if rootslice[i] != nil {
				Q.AddRootNode(rootslice[i])
			}
		}
	}
}

//Promote a heap node to a root heap node
func (Q *Fheap) Promote(H *Heapnode) {
	//clear the silblings
	H.DetachHeapnode()
	//clear the parent
	H.parent = nil
	if Q.min != nil {
		//inset next to min node
		tmp := Q.min.left
		Q.min.left = H
		H.right = Q.min
		H.left = tmp
		tmp.right = H
		//update the min pointer if necessary
		if H.key < Q.min.key {
			Q.min = H
		}
	} else {
		//or if root is currently empt, make it the min node
		H.right = H
		H.left = H
		Q.min = H
	}
	//H is now a root node
	H.tag = false
	Q.rootnodes++
}

/*RecursiveCut must be called if a child node is cut
from its parent. If the parent already has been tagged
it has to be removed also. If the parent's parent
has been taggen, also remove this node, and so forth... */
func (Q *Fheap) RecursiveCut(H *Heapnode) {
	//if the heap node has no parent (is a root node), nothing happens
	if H.parent != nil {
		//untagged nodes get tagged, if they get tagged again they are moved to the root
		if H.tag == false {
			H.tag = true
		} else {
			formerparent := H.parent
			Q.Promote(H)
			//after moving the node to root we have to check its former parent
			Q.RecursiveCut(formerparent)
		}
	}
}

//UpdateKey is used to change the key value of a heap node
//depending on the value it might has to be moved to the root
func (Q *Fheap) UpdateKey(H *Heapnode, newkey float64) {
	//if the new key is larger than the old key, it will just be updated and nothing else happens
	if newkey > H.key {
		H.key = newkey
	} else {
		//if it is less than that, we have to check is the tree still statisfies heap condition
		H.key = newkey
		if H.parent != nil {
			if H.parent.key > H.key {
				//if not, move it to the root
				formerparent := H.parent
				Q.Promote(H)
				//after moving the node to root we have to check its former parent
				Q.RecursiveCut(formerparent)
			}
		}
	}
}

/*Popmin function. Removes the current min node from the heap
and triggers the Cleanup. */
func (Q *Fheap) Popmin() (float64, int) {
	minval := Q.min.key
	minindex := Q.min.index
	/* remove the min node from the root.
	the min pointer is put to some other root node
	(i.e. some silbling tmp3 of the former min) */
	tmp3 := Q.min.left
	if tmp3 == Q.min { //this happens if Q.min is the only root node
		tmp3 = nil
	}
	// Q.PrintRootnodes()
	//if the min node has no children we just detach it
	//and make its silbling to the new min node
	if Q.min.child == nil {
		Q.min.DetachHeapnode()
		Q.rootnodes-- //
		Q.nodes--
		Q.min = tmp3
	} else {
		//otherwise we have to decide how to handle the former children of Q.min
		//get the first child node of the current min node
		tmp := Q.min.child
		cc := Q.min.children
		//and its silbling (tmp2 == tmp if there are no silblings)
		tmp2 := tmp.left

		Q.min.DetachHeapnode()
		Q.rootnodes--
		Q.nodes--
		Q.min = tmp3 //the old Q.min is gone after this point.

		/* it does not matter if tmp3 is actually the new min
		as this will be handled by the Cleanup() function later */

		/* iterate through all the children of the
		former min node and promote them to root nodes*/
		tmp2 = tmp.left
		for k := 0; k < cc; k++ {
			tmp3 = tmp2.left
			Q.Promote(tmp2)
			tmp2 = tmp3
		}
	}
	//and at the end trigger the cleanup
	/* 	fmt.Println("after poppin min")
	   	Q.PrintRootnodes()
	*/
	Q.Cleanup()

	/* 	fmt.Println("after cleanup")
	   	Q.PrintRootnodes() */

	return minval, minindex
}

/*Getmin gets the key and the index of the current min node */
func (Q *Fheap) Getmin() (float64, int) {
	return Q.min.key, Q.min.index
}

/*GetKey returns the key of node H*/
func (H *Heapnode) GetKey() float64 {
	return H.key
}

/*GetIndex returns the key of node H*/
func (H *Heapnode) GetIndex() int {
	return H.index
}

// SetIndex is used to set the index of a node
func (H *Heapnode) SetIndex(i int) {
	H.index = i
}

// GetNodes returns the current number of nodes in the heap Q
func (Q *Fheap) GetNodes() int {
	return Q.nodes
}

// GetRootNodes returns the current number of nodes in the heap Q
func (Q *Fheap) GetRootNodes() int {
	return Q.rootnodes
}

// PrintRootnodes iterates through the silbling list at root level and prints them to the console
func (Q *Fheap) PrintRootnodes() {
	if Q.min != nil {
		tmp := Q.min.left
		fmt.Println("  Roots in total", Q.rootnodes)
		fmt.Println("  Root", 1, Q.min, "(current min)")
		c := 1
		for tmp != Q.min {
			c++
			fmt.Println("  Root", c, tmp)
			tmp = tmp.left
		}
	}
}
