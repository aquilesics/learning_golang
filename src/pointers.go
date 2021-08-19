package main

import "fmt"

func main() {
	x := 5
	y := &x
	fmt.Printf("%T", y) //show the  address of the x varaible
	fmt.Println(*y)     //show the content of the address

}
