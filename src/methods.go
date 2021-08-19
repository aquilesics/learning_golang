package main

import (
	"fmt"
)

type pessoa struct {
	nome string
	idade int
	tel string
}

func (p pessoa) show_details (){
	fmt.Printf("nome:%v \nidade:%v \ntelefone:%v",p.nome,p.idade,p.tel)
}


func main(){

	joao := pessoa{
		nome:"Joao",
		idade: 23,
		tel: "+55 11 99999-5050",
	}

	joao.show_details()

}