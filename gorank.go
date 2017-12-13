// GoRank library
// Created by: 	Thomas Shallenberger 12/12/2017
// Purpose:		Contains all problem solutions submitted to HackerRank.com  
package gorank
import "fmt"
import _ "strconv"
import math "math"
import _ "strings"


// main running application. 
func GoRank() {
	RansomNote()
}


// Ransom Note algorithm

// INPUT: 
// 	1. 2 space delimited ints: (m, n) -> # words in magazine, # words in ransom note
// 	2. m space delimited strings denoting magazine words (mword)
// 	3. n space delimited strings denoting words in ransom note. (nword) 

// CONSTRAINTS: 
// 	1 <= m, n <= 30,000
// 	1 <= mword.len, nword.len <= 5
// 	All words are subsets of all English alphabetic letters (a-z) AND (A-Z)
// 	All words are CASE-SENSITIVE

// func IntScanln(n int) ([]int, error) {
// 	x := make([]int, n)
// 	y := make([]interface{}, len(x))
// 	for i := range x {
// 		y[i] = &x[i]
// 	}
// 	n, err := fmt.Scanln(y...)
// 	x = x[:n]
// 	return x, err
// }

func RansomNote() {
	var mCount, nCount, n int
	var err error

	if _, err := fmt.Scanf("%d %d\n", &mCount, &nCount); 	err != nil {
		fmt.Println("Err")
		return
	}

	mInput := make([]interface{}, mCount)
	nInput := make([]interface{}, nCount)
	magazine := make([]string, mCount)
	note := make([]string, nCount)

	for i := range magazine { mInput[i] = &magazine[i]	}
	for i := range note { nInput[i] = &note[i]	}

	if n, err = fmt.Scan(mInput...); err != nil {fmt.Println(err)}
	magazine = magazine[:n]
	if n, err = fmt.Scan(nInput...); err != nil {fmt.Println(err)}
	note = note[:n]

	words := make(map[string]int)
	result := "Yes"
	for _, s := range magazine { words[s]++ }
	for _, s := range note { 
		if words[s] > 0 { words[s]-- } else { result = "No"; break } 
	}

	fmt.Println(result)









}


// ANAGRAMS ALGORITHM
// Completed: 12/12/2017

func Anagrams() {
	fmt.Println("Initializing ANAGRAMS runtime...")
	//Enter your code here. Read input from STDIN. Print output to STDOUT
	
	// First line contains single string A. Second line contains string B. 
	// Output should be an integer (total deletions to be made). 

	// A and B are sets of characters in the set of all lowercase English letters (a-z). 
	var first, second string
	var counter [26]int64
	var result uint32
	fmt.Scanf("%s", &first)
	fmt.Scanf("%s", &second)
	
	// Now we have first, second, counter, and result. 

	for _, c := range(first) { counter[c - 'a']++ }
	for _, c := range(second) { counter[c - 'a']-- }
	for _, count := range(counter) { result += Abs(count) }
	fmt.Printf("%d", result)


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

func Min(x, y int64) int64 {
    if x < y {
        return x
    }
    return y
}

func Max(x, y int64) int64 {
    if x > y {
        return x
    }
    return y
}

func Abs(x int64) uint32 {
	f_ret := float64(x)				// Convert to float64
	return uint32(math.Abs(f_ret))	// Get Absolute val, convert to uint32 and return
}