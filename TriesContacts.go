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
	var counts []int
	trie := NewTrie()

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
		if op == "add" {
			// Add the string to the Trie. 
			trie.ExistsOrAdd(val)
		} else {
			_, count := trie.FindPartial(val)
			counts = append(counts, count)
		}
	}
	for _, count := range counts {
		fmt.Println(count)
	}
	
}

// Design a Trie. 

type TrieLink struct {
	value rune
	count int
	link *Trie
}

type Trie struct {
	childNodes []*TrieLink
	stringEnd bool
}

func NewTrie() *Trie {
	return &Trie{stringEnd: true, childNodes: make([]*TrieLink, 0)}
}

func (r *Trie) PrintTrie() {
	var runeVals []rune
	for _, link := range r.childNodes {
		runeVals = append(runeVals, link.value)
		link.link.PrintTrie()
	}
	if len(runeVals) > 0 {
		fmt.Println(string([]byte(string(runeVals))))
	}
}

func findLink(links []*TrieLink, val rune, count int) (*Trie, bool, int) {
	for _, link := range links {
		if link.value == val {
			link.count = link.count + count			
			return link.link, true, link.count
		}
	}
	return nil, false, 0
}

func (r *Trie) ExistsOrAdd(s string) (*Trie, bool) {
	check := true 	// What is this?9
	i := r 			// Why are we making a copy of this
	for _, runeVal := range s {							// For each rune in string
		trie, ok, _ := findLink(i.childNodes, runeVal, 1)		// Find the runeVal in current Trie's childNodes
		if !ok {										// if runeVal doesn't exist in the current Tries childNodes
			trie = new(Trie)							// Make a new trie
			trie.childNodes = make([]*TrieLink, 0)		// make new childNodes
			i.childNodes = append(i.childNodes, &TrieLink{value: runeVal, link:trie, count: 1}) // add runeval to current Trie's childnodes
		}
		i = trie										// If runeVal does exist, set i = trie, loop
	}
	// After Looping over entire string and adding appropriate Trie nodes
	if !i.stringEnd {			// if i.stringEnd is false
		i.stringEnd = true		// we looped over entire string, so set stringEnd to true
		check = false			// No idea what the check means
	}
	return i, check				// Return the last? trie and check
}

func (r *Trie) FindPartial(s string) (*Trie, int) {

	retTrie := r
	fullPartialExists := true
	count := 0
	for _, runeVal := range s {
		retTrie, fullPartialExists, count = findLink(retTrie.childNodes, runeVal, 0)
		
		if !fullPartialExists {
			fullPartialExists = false
			count = 0
			break
		}
	}

	return retTrie, count
}
