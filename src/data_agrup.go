package main

import "fmt"

var x [5]int

func main() {
	// arrays
	// or you can declare this way:
	array_7 := [7]int{1, 2, 3, 4, 5, 6, 7}

	for i := 0; i < 5; i++ {
		x[i] = (i + 1) * 10
	}
	fmt.Println(x)
	fmt.Println(array_7)

	// slices
	slice := []int{1, 2, 3, 4, 5, 6, 7}
	slice = append(slice, 8)

	fmt.Print(slice[3], "\n\n")

	frutas := []string{"banana", "uva", "maça", "jaca"}

	for indice, valor := range frutas {
		fmt.Println(indice, valor)
	}

	// slicing slices
	fmt.Println(frutas[0:2])

	// append function

	carrinho := []string{"abacaxi"}

	carrinho = append(carrinho, frutas...) //... is like "unpacking" list in python

	fmt.Print("carrinho:", carrinho, "\n")

	// make func. create a slice with a max capacitie

	//numeros := make([]int,5,10)

	// mult-dimensional slice

	grid := [][]string{
		[]string{"      0", "x", "0\n"},
		[]string{"     x", "x", "x\n"},
		[]string{"     0", "x", "0\n"},
	}

	fmt.Println(grid)

	// map ( like dicts in python)
	amigos := map[string]int{
		"joao":  12,
		"vitor": 23,
		"lucas": 45,
	}

	fmt.Println(amigos["joao"])

	// comma ok. check if the key exits
	if age, ok := amigos["fantasma"]; !ok {

		fmt.Printf("Fantasma nao existe")

	} else {

		fmt.Println("idade do fantasma:", age)

	}

	// structs

	type cliente struct { //atempt to uppercase keys, with you want export it to json!!s
		nome      string
		sobrenome string
		fumante   bool
	}

	client1 := cliente{
		nome:      "Joao",
		sobrenome: "Silva",
		fumante:   true,
	}

	fmt.Println("\n", client1)

}
