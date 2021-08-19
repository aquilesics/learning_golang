package main

import "fmt"

type animais interface {
	barulho()	
}

type gato struct {
	nome string		
}

type cachorro struct {
	nome string	
}

func (c cachorro ) barulho (){
	fmt.Println("AUAUAUAU")
}

func (g gato ) barulho (){
	fmt.Println("Miau, Miau")
}

func expressar(a animais){
	a.barulho()
}





func main() {

	dog := cachorro{
		nome:"Rex",
	}

	cat := gato{
		nome: "felix",
	}

	expressar(dog)
	fmt.Println("\n")
	expressar(cat)
}
