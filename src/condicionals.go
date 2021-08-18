package main

import "fmt"

func main(){
	

	if x:=false; x == true {
		fmt.Println("verdade")
	}else {
		fmt.Println("falso")
	}

	// if condition{

	// }else if{

	// }else ..

	// switch case
	x:=29
 
	switch{
		case x == 2: 
		case x == 4:
		case x > 5: fmt.Println("x>5")

	}

	// fallthrough
	switch true{
		case x == 2: 
		case x == 4:
		case x > 5: fmt.Println("x>5")
			fallthrough
		case x == 2342: fmt.Println("passa atraves do case e mostra o resultado")	
		case x == 442: fmt.Println("passa atraves do case e mostra o resultado")
		default:
			fmt.Println("valor padrao")

	}

	switch x{
		case 1, 2: fmt.Println("1 ou 2")
		case 3, 4:fmt.Println("3 ou 4")
		case 5, 29:fmt.Println(" 4 ou 29")
		default:fmt.Println("padrao")

	}

	// example with switch types
	
	var t interface{}
	t = true

	switch t.(type){
		case bool:
			fmt.Println("tipo bool")
		case int : 
			fmt.Println("tipo integer")
		case string:
			fmt.Println("tipo String")
		

	}
		


	
}