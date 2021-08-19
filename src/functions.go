package main

import "fmt"

func main() {
	a := soma(2, 2)
	xs := []int{1, 2, 3}

	for _,v := range xs {
		fmt.Print(v)
	}

	defer fmt.Println("executada em ultimo caso, por fim")

	fmt.Println(a) //functions as pass by value

	sli := []int{1, 2, 3, 4}

	double(sli...)
	

	// anonymous functions
	func(x int) { fmt.Println("\n anonima", x*3) }(4)

}

func soma(x int, y int) int {

	return x + y
}

func double(x ...int) []int {
	re := []int{}
	for _, v := range x {
		re = append(re, v*2)

	}
	fmt.Println("\ndouble numbers\n", re)
	return re
	

	
}
