package main 
// Go program must start with it.
// Packages are Go’s way of organizing and reusing code. There are two types of Go pro‐
// grams: executables and libraries.

import "fmt" 

//printa o classico hello world :)
func main(){
	fmt.Println("Ola mundo")

	_,erro := fmt.Println("demonstrando", "retorno de funcoes")
	fmt.Println(erro)
}
