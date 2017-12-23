// Tries: Contacts solution
// Created by: 	Thomas Shallenberger 
// Date: 		12/22/2017

package gorank

import "fmt"
import _ "strings"


////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Tries: Contacts function
// INPUT: 
//		1. Single int (n) denoting the number of ops to perform
// 		2. Each subsequent line (i) of the n lines contains an op, in one of two forms:
// 			- add name : name denotes a contact name. Op should store name in application. 
// 			- find partial : partial denotes a partial name to search the application for. 
// 			  It counts the number of contacts starting with partial and print the count on a new line. 
// OUTPUT:
// 		1. For each find partial op, print the number of contact names starting with partial on a newline. 
////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Sample Input
// 4
// add hack
// add hackerrank
// find hac
// find hak

// Sample Output
// 2
// 0
////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TriesContactsFunc(){
	// number of ops
	var n int

	if _, err := fmt.Scanf("%d\n", &n); 	err != nil {
		fmt.Println("Err")
		return
	}

	// Grab all ops, sort em into add and find ops. 
	for i:= 0; i < n; i++ {
		var op, val string
		if _, err := fmt.Scanf("%s %s\n", &op, &val); 	err != nil {
			fmt.Println("Err")
			return
		}

		// Read in the op and val. Time for Tries. 
		fmt.Printf("%s\n", op + " - " + val)

	}
}

// Design a Trie. 

type TrieLink struct {
	value rune
	link *Trie
}

type Trie struct {
	childNodes []TrieLink
	stringEnd bool
}

func findLink(links []TrieLink, val rune) (*Trie, bool) {
	return nil, false
}

func (r *Trie) FetchAndUpdateOrCreate(s string) (*Trie, bool) {
	return nil, false
}

func NewTrie() *Trie {
	return &Trie{stringEnd: true, childNodes: make([]TrieLink, 0)}
}