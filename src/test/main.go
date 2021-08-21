package main

import "fmt"

func main() {
	fmt.Println(Soma(1, 2, 3, 7))
}

func Soma(x ...int) int {
	total := 0

	for _, v := range x {
		total += v
	}

	return total
}
