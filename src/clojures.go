package main

import "fmt"

func main(){

	a := cloj()

	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

}

func cloj() func() int{
	x := 10 //scope of clojure
	return func() int {
		x++
		return x
	}
}