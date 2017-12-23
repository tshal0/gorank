// Algorithm solutions for HackerRank
// created by: 		Thomas Shallenberger
// date: 			12/22/2017

package gorank 

import "fmt"

func BubbleSortFunc() {
	// Take array size int n
	// Take n space seperated ints for A

	var n, swapCount int
	
	if _, err := fmt.Scanf("%d\n", &n); 	err != nil {
		fmt.Println(err)
		return
	}

	input := make([]interface{}, n)
	unsorted := make([]int, n)
	for i := range unsorted { input[i] = &unsorted[i] }
	

	if _, err := fmt.Scan(input...); err != nil {fmt.Println(err)}
	unsorted = unsorted[:n]

	for i := 0; i < n; i++ {
		for j := 0; j < n-1; j++ {
			if unsorted[j] > unsorted[j+1] {
				tmp := unsorted[j]
				unsorted[j] = unsorted[j+1]
				unsorted[j+1] = tmp
				swapCount++
			}
		}
	}

	fmt.Printf("Array is sorted in %d swaps.\n", swapCount)
	fmt.Printf("First Element: %d\n", unsorted[0])
	fmt.Printf("Last Element: %d\n", unsorted[n-1])



}