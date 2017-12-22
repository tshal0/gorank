// Running Median solution
// Created by: 	Thomas Shallenberger 
// Date: 		12/22/2017

package gorank

import "fmt"
import "container/heap"



// Running Median algorithm
// INPUT: 
//		1. Single int (n) denoting the number of integers in the datastream
// 		2. Each subsequent line (i) of the n lines contains int a_i, to be added to list. 
// OUTPUT:
// 		After each subsequent integer a_i, print the lists updated median on a newline as a single double precision number 
// 		scaled to 1 decimal place. 


func RunningMedianFunc() {
	var n int
	var runningMedian RunningMedian
	var medians []float64
	// Actually surprised this works. Init expects a pointer receiver. 
	runningMedian.Init()
	
	if _, err := fmt.Scanf("%d\n", &n); 	err != nil {
		fmt.Println("Err")
		return
	}

	for i := 0; i < n; i++ {
		var newVal int
		if _, err := fmt.Scanf("%d\n", &newVal); 	err != nil {
			fmt.Println("Err")
			return
		}
		runningMedian.InsertValue(newVal)
		median := runningMedian.GetMedian()
		medians = append(medians, median)
		// fmt.Println("### MEDIAN ###")
		// fmt.Println("##############")
	}

	for i := 0; i < len(medians); i++ {
		fmt.Printf("%.1f\n", medians[i])
	}
}


type MinHeap []int

func (h MinHeap) Len() int           { return len(h) }
func (h MinHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h MinHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h MinHeap) Top() int { 
	if len(h) > 0 {
		return h[0]
	} else { return 0 }
}

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
}

// Take copy receiver of RunningMedian, get median. 
func (r RunningMedian) GetMedian() float64 {
	// get min, max lengths
	minhLen, maxhLen := r.minh.Len(), r.maxh.Len()
	// get minh, maxh tops as floats

	minhTop, maxhTop := float64(r.minh.Top()), float64(r.maxh.Top())

	// If min and max heaps are equal in size and max heap size is GREATER than 0, return the sum of both roots DIVIDED in HALF
	if (maxhLen == minhLen && maxhLen > 0) {
		return (maxhTop + minhTop)/2
	} else {				
		// Find 	
		if (maxhLen > minhLen) {
			return maxhTop
		} else if minhLen > maxhLen {
			return minhTop
		} else {
			return 0
		}
	}
}
