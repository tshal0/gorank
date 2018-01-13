// Merge Sort solution
// Created by: 	Thomas Shallenberger 
// Date: 		1/2/2018

// TODO:

// implement merge sort, count the number of inversions made. 
package gorank

import _ "container/heap"
import "fmt"
import _ "bufio"
import _ "os"
import "bufio"
import "os"
import "strings"


func MergeSortFunc() {
	var d int // dataset count
	// each 2d subsequent lines describe each respective dataset over 2 lines: 
	// 		int n, dataset size
	// 		n space seperated ints 
	
	
	
	
	
	if _, err := fmt.Scanf("%d\n", &d); 	err != nil {
		fmt.Println("Err")
		return
	}

	for i := 0; i < d; i++ {
		var ds MSDataSet
		if _, err := fmt.Scanf("%d\n", &ds.size); 	err != nil {
			fmt.Println("Err")
			return
		}
		fmt.Println(ds)

		// var n, k, c int
		
		fmt.Println(ds.data)
		

	}
	
}

func MergeSort() {

}

type MSDataSet struct {
	data []int
	size int
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return 
}

// My fast implementation of inputting a space delimited set of ints

func ScanInts(size int) (values []int) {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	
	for i:= 0; i< size; i++ {
		scanner.Scan()
		
		 values := append(values, toInt(scanner.Bytes()))
		
		fmt.Println(values)
	}

	return values
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