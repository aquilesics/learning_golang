// An array is a numbered sequence of elements of a single type with a fixed length. In
// Go, they look like this:

package main

import "fmt"

func main() {
	// := auto detect type, type inferece
	x := 10
	y := "olá Pessoa"
	w := true

	fmt.Printf("value:%v | type:%T\n", x, x)
	fmt.Printf("value:%v | type:%T\n", y, y)
	fmt.Printf("value:%v | type:%T\n", w, w)

	x = 20
	fmt.Printf("nem value for x => value:%v type:%T\n", x, x,)

	// you can do this too, need a new variable (z, in this case):

	x , z := 40, 50

	fmt.Printf("z value:%v type:%T\n", z, z,)

	const (//create sequencial number,without types,start=0
			a = iota * 10
			b 
			c 
			d 
			e 
	)
	
	fmt.Println(a,b,c,d,e)


}
