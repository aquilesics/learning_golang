package main

import (
	"fmt"
)

var x int = 10

func main(){
	// interpreted string literals vs Raw string literal

	inter := "ola!\n como vai voce?\t espero q bem\n\n"

	raw:= `ola!\n como vai voce?\t espero q bem`

	fmt.Println(inter,raw)

	// each character is 'Rune' 
	// group #1 show content on terminal
		// fmt.Print()
		// fmt.Println()
		// fmt.Printf()
		
	// group #2 return String
		// fmt.Sprint()
		// fmt.Sprintln()
		// fmt.Spintf()

	// group #3 write in file, or server
		// fmt.Fprint()
		// fmt.Fprintln()
		// fmt.Fprintf()

}					
