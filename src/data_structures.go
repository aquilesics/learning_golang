// An array is a numbered sequence of elements of a single type with a fixed length. In
// Go, they look like this:

package main

import "fmt"

func main() {

	var x [5]int
	x[4] = 100

	y := [5]float64{1, 2, 3, 4, 5}

	fmt.Println("Arrays:", x, y)

	// Slices

	// w := make([]float64, 5, 10)

}
