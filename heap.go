// HEAP implementation

package gorank

import _ "fmt"
import "container/heap"
import "fmt"

// int capacity = 10
// int size = 10

// int[] items = new int[capacity]
// getLeftChildIndex(int parent)
// getRightChildIndex(int parent)
// getParentIndex(int parent)

// hasLeftChild(int parent)
// hasRightChild(int parent)
// hasParent(int parent)

// leftChild(int parent)
// rightChild(int parent)
// parent(int parent)

// swap
// ensureExtraCapacity
// peek
// poll

type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Top() int { return h[0]}

func (h *MinHeap) Push(x interface{}) {
    // Push and Pop use pointer receivers because they modify the slice's length,
    // not just its contents.
    *h = append(*h, x.(int))
}

func (h *MinHeap) Pop() interface{} {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}

type MaxHeap struct {
	MinHeap
}

// So, I just learned that h MaxHeap is a "receiver". 
// It should be a pointer if changes made to the receiver are to be persistent. 
// It acts very similarly to class Methods in OOP, which Go is NOT. 
func (h MaxHeap) Less(i, j int) bool { return h.MinHeap[i] > h.MinHeap[j]}

// func (h MaxHeap) Len() int {return len(h.MinHeap)}
// func (h MaxHeap) Swap(i, j int) {h.MinHeap[i], h.MinHeap[j] = h.MinHeap[j], h.MinHeap[i]}

type RunningMedian struct {
	minh MinHeap 
	maxh MaxHeap 
	allminh []MinHeap
	allmaxh []MaxHeap
}

// Init RunningMedian struct

func (r *RunningMedian) Init() {
	h1 := &MinHeap{}
	heap.Init(h1)

	h2 := &MaxHeap{}
	heap.Init(h2)

	r.minh = *h1
	r.maxh = *h2
}

// Insert Value: 
// Take a pointer to a RunningMedian struct, 
// Insert val into it
// Rebalance if necessary
func (r *RunningMedian) InsertValue(val int) {

	minhLen, maxhLen := r.minh.Len(), r.maxh.Len()
	// If max empty OR val <= maxh.top: maxh.push(val) ELSE: minh.push
	if r.maxh.Len() == 0 || val <= r.maxh.Top() {
		heap.Push(&r.maxh, val)
	} else {
		heap.Push(&r.minh, val)
	}
	
	fmt.Println("PUSHED ONTO HEAP")
	r.PrintMaxHeap()
	r.PrintMinHeap()
	fmt.Println("###############")

	// Find bigger heap, after getting the sizes again
	minhLen, maxhLen = r.minh.Len(), r.maxh.Len()
	min, max := MinMax(minhLen, maxhLen)

	// If length difference is GREATER than 1, rebalance (since if it's 1 or 0, we leave it be)
	if max - min > 1 { 
		if maxhLen > minhLen {
			// r.minh.Push(r.maxh.Top())
			// r.maxh.Pop()
			heap.Push(&r.minh, r.maxh.Top())
			heap.Pop(&r.maxh)
		} else {
			// r.maxh.Push(r.minh.Top())
			// r.minh.Pop()
			heap.Push(&r.maxh, r.minh.Top())
			heap.Pop(&r.minh)
		}
	}

	fmt.Println("REBALANCED HEAPS")
	r.PrintMaxHeap()
	r.PrintMinHeap()
	fmt.Println("###############")
}

// Take copy receiver of RunningMedian, get median. 
func (r RunningMedian) GetMedian() int {
	// get min, max lengths
	minhLen, maxhLen := r.minh.Len(), r.maxh.Len()

	// If min and max heaps are equal in size and max heap size is GREATER than 0, return the sum of both roots DIVIDED in HALF
	if (maxhLen == minhLen && maxhLen > 0) {
		return (r.maxh.Top() + r.minh.Top())/2
	} else {				
		// Find 	
		if (maxhLen > minhLen) {
			return r.maxh.Top()
		} else if minhLen > maxhLen {
			return r.minh.Top()
		} else {
			return 0
		}
	}
}

// Print MAX heap
func (r RunningMedian) PrintMaxHeap() {
	fmt.Println("### MAXHEAP ###")
	fmt.Println(r.maxh)
	fmt.Println("###############")
}
// Print MIN heap
func (r RunningMedian) PrintMinHeap() {
	fmt.Println("### MINHEAP ###")
	fmt.Println(r.minh)
	fmt.Println("###############")
}