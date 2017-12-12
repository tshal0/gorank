// GoRank library
// Created by: 	Thomas Shallenberger 12/12/2017
// Purpose:		Contains all problem solutions submitted to HackerRank.com  
package gorank
import "fmt"

// main running application. 
func GoRank() {
	Anagrams()
}

// ANAGRAMS ALGORITHM


func Anagrams() {
	fmt.Println("Initializing ANAGRAMS runtime...")
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	
	// First line contains single string A. Second line contains string B. 
	// Output should be an integer (total deletions to be made). 

	// A and B are sets of characters in the set of all lowercase English letters (a-z). 



}


// LEFT SHIFT ALGORITHM 
// Completed: 12/12/2017 by Thomas Shallenberger

func Leftshift() {
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	 
	// TODO: Read in array size, # rotations, and array vals. 
	
	var size, rotation int

	if _, err := fmt.Scanf("%d %d", &size, &rotation); 	err != nil {
		fmt.Println("Err")
		return
	}
	var array []int
	for i := 0; i < size; i++ { array = append(array, 0) }
	for i := 0; i < size; i++ {
		newLoc := (i + (size - rotation)) % size
		fmt.Scanf("%d", &array[newLoc])
	}
	// Now we have both of our params read in. 
	for i := 0; i < size; i++ {fmt.Printf("%d ", array[i])}

}

func IntScanln(n int) ([]int, error) {
	x := make([]int, n)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, err := fmt.Scanln(y...)
	x = x[:n]
	return x, err
}
