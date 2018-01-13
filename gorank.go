// GoRank library
// Created by: 	Thomas Shallenberger 12/12/2017
// Purpose:		Contains all problem solutions submitted to HackerRank.com  

// Easy challenges were thrown into the main gorank file. 
// Hard challenges were spread to other files. 


package gorank
import "fmt"
import _ "strconv"
import math "math"
import _ "strings"
import "sync"
import "errors"



// main running application. 
func GoRank() {
	MergeSortFunc()
}

// Balanced Brackets algorithm

// 		INPUT:
// 		1. Single integer n, denoting number of strings
// 		2. Each line i of the n subsequent lines consists of a single string s, denoting a sequence of brackets. 
// 		CONSTRAINTS
// 		1 <= n <= 10^3
// 		1 <= len(s) <= 10^3

func BalancedBrackets() {

	var n int 
	if _, err := fmt.Scanf("%d\n", &n); 	err != nil {
		fmt.Println("Err")
		return
	}

	lines := make([]string, n)
	
	for i := 0; i < n; i++ { 			// For each line of brackets to be entered
		fmt.Scanf("%s\n", &lines[i]) 	// Scan in the string of brackets
	}
	for _, s := range lines {
		if IsBalanced(s) {fmt.Println("YES")} else { fmt.Println("NO")}
	}
}

func IsBalanced(line string) bool {
	
	if len(line) % 2 != 0 {return false}
	s := NewStack()
	for _, c := range line {	// Iterate over the string 
		if c == '{' {				// If we have an opening bracket, push a closing one on
			s.Push('}')
		} else if c == '[' {
			s.Push(']')
		} else if c == '(' {
			s.Push(')')
		} else {							// Else, if the stack is empty or current bracket doesn't
			if s.Len() == 0 || c != s.Peek() {	// match what's on the stack, print NO
				return false
			}
			s.Pop()					// Pop the closing bracket off the stack
		}
	}
	return s.Len() == 0;
}

// LINKED LIST STACK

type (
	Stack struct {
		top *node
		length int
	}
	node struct {
		value interface{}
		prev *node
	}
)

func NewStack() *Stack {
	return &Stack{nil, 0}
}

func (s *Stack) Len() int {
	return s.length
}

func (s *Stack) Peek() interface{} {
	if (s.length == 0) {return nil}
	return s.top.value
}

func (s *Stack) Pop() interface{} {
	if (s.length == 0) {return nil}
	n := s.top
	s.top = n.prev
	s.length--
	return n.value
}

func (s *Stack) Push(value interface{}) {
	n := &node{value, s.top}
	s.top = n
	s.length++
}
// Good for most use cases, but slice stacks can be space intensive.

type sliceStack struct {
	lock sync.Mutex // you don't have to do this if you don't want thread safety
	s []rune	
	top *rune
	Empty bool
}

func NewSliceStack() *sliceStack {
   return &sliceStack {sync.Mutex{}, make([]rune,0), nil, true }
}

func (s *sliceStack) Push(v rune) {
   s.lock.Lock()
   defer s.lock.Unlock()

   s.s = append(s.s, v)
   s.Empty = false
   s.top = &s.s[len(s.s)-1]
}

func (s *sliceStack) Pop() (rune, error) {
   s.lock.Lock()
   defer s.lock.Unlock()


   l := len(s.s)
   if l == 0 {
	   return 0, errors.New("Empty Stack")
   }
   newLen := l-1
   res := s.s[newLen]
   s.s = s.s[:newLen]
   s.Empty = (newLen == 0)
   if !s.Empty {
	   s.top = &s.s[len(s.s)-1]
	} else {
	   s.top = nil
   }
   return res, nil
}

func (s *sliceStack) Peek() (rune) {
	var top rune
	if !s.Empty { 
		top = *s.top 
	} else {
		top = '\n'
	}
	return top
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
    if x > y {
        return y
    }
    return x
}

func Max(x, y int64) int64 {
    if x > y {
        return x
    }
    return y
}

func MinMax(x, y int) (int,int) {
	if x > y {
		return y, x
	} else {
		return x, y
	}
}

func Abs(x int64) uint32 {
	f_ret := float64(x)				// Convert to float64
	return uint32(math.Abs(f_ret))	// Get Absolute val, convert to uint32 and return
}