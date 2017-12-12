package main
import (
	"fmt"
)
func main() {
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

func intScanln(n int) ([]int, error) {
	x := make([]int, n)
	y := make([]interface{}, len(x))
	for i := range x {
		y[i] = &x[i]
	}
	n, err := fmt.Scanln(y...)
	x = x[:n]
	return x, err
}


