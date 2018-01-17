// Merge Sort solution
// Created by: 	Thomas Shallenberger
// Date: 		1/2/2018

// TODO:

// Create DataSet struct
// Add ScanInts functionality for scanning in space delimited integers.
//

package gorank

// import _ "container/heap"
import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// import _ "bufio"
// import _ "os"

// MergeSortFunc should perform MergeSort alg for hackerrank
func MergeSortFunc() {
	var d int // dataset count
	// each 2d subsequent lines describe each respective dataset over 2 lines:
	// 		int n, dataset size
	// 		n space seperated ints

	if _, err := fmt.Scanf("%d\n", &d); err != nil {
		fmt.Println("Err")
		return
	}
	invCounts := make([]uint64, d)
	datasets := make([]MSDataSet, d)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	for i := 0; i < d; i++ {
		var ds MSDataSet
		ds.ScanInts(scanner, 0)
		datasets[i] = ds
	}
	for i, ds := range datasets {
		inv := ds.MergeSort(0, ds.size-1)
		invCounts[i] = inv
	}

	for _, val := range invCounts {
		fmt.Println(val)
	}
}

// MSDataSet data structure stores merge sort data set
// Might actually begin to use this for all int sets.
type MSDataSet struct {
	data    []int
	size    int
	visited bool
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

// ScanInts scans in a space delimited set of ints
func (ds *MSDataSet) ScanInts(scanner *bufio.Scanner, size int) {
	// Get dataset size

	if size == 0 {
		scanner.Scan()
		ds.size = toInt(scanner.Bytes())
	} else {
		ds.size = size
	}
	// Init scanner

	// Init dataset to be size d
	ds.data = make([]int, ds.size)

	for i := 0; i < ds.size; i++ {
		scanner.Scan()
		ds.data[i] = toInt(scanner.Bytes())
	}

}

// MergeSort performs merge sort sorting algorithm
// Takes starting and ending points, ran on original dataset.
func (ds *MSDataSet) MergeSort(start, end int) (inversions uint64) {
	if start == end {
		return 0
	}

	mid := (start + end) / 2

	inversions += ds.MergeSort(start, mid) // Left
	inversions += ds.MergeSort(mid+1, end) // Right
	inversions += ds.Merge(start, end, mid)

	return inversions
}

//Merge sort
func (ds *MSDataSet) Merge(start, end, mid int) (inversions uint64) {

	newArray := make([]int, (end - start + 1))

	curr, i, j := 0, start, mid+1

	for i <= mid && j <= end {
		if ds.data[i] > ds.data[j] {
			newArray[curr] = ds.data[j]
			curr++
			j++
			inversions += uint64(mid - i + 1)
		} else {
			newArray[curr] = ds.data[i]
			curr++
			i++
		}
	}
	for i <= mid {
		newArray[curr] = ds.data[i]
		curr++
		i++
	}
	for j <= end {
		newArray[curr] = ds.data[j]
		curr++
		j++
	}

	dsa := ds.data[start:(start + (end - start) + 1)]

	copy(dsa, newArray)
	return inversions
}

// Custom split func I found
func crunchSplitFunc(data []byte, atEOF bool) (advance int, token []byte, err error) {

	// Return nothing if at end of file and no data passed
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Find the index of the input of a newline followed by a
	// pound sign.
	if i := strings.Index(string(data), "\n#"); i >= 0 {
		return i + 1, data[0:i], nil
	}

	// If at end of file with data return the data
	if atEOF {
		return len(data), data, nil
	}

	return
}
