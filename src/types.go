package main

import (
	"fmt"
)

var x int = 10

type hotdog int

var b hotdog = 929

func main() {
	// types in Go are Static, don't change
	// x = 34.6 dont work

	// primitivos: int, bool, string,float
	// compostos : array, slice, struct, map

	fmt.Println(x)

	// Create types
	fmt.Printf("value:%v | type:%T\n", b, b)

	// convert b (hotdog) in int
	x = int(b)
	fmt.Printf("x <= b:%v", x)

	// numeric types
	// byte == int8
	// rune == int32

	

}
