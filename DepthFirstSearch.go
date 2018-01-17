// DFS algorithm solution
// Created by: 	Thomas Shallenberger
// Date: 		1/17/2018

// TODO:

// Input n, m grid sizes
// Input m rows of n space delimited chars

package gorank

import (
	"bufio"
	"fmt"
	"os"
)

// DepthFirstSearchSolution function
func DepthFirstSearchSolution() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	rows := toInt(scanner.Bytes())
	scanner.Scan()
	cols := toInt(scanner.Bytes())
	grid := make([]MSDataSet, rows)

	for i := 0; i < rows; i++ {
		var row MSDataSet
		row.ScanInts(scanner, cols)
		grid[i] = row
	}
	fmt.Print(grid)
}
