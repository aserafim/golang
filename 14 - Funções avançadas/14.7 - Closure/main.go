package main

import "fmt"

func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {

	//Funções que referenciam variáveis que são externas
	//ao seu escopo

	texto := "Dentro da função main"
	fmt.Println(texto)

	funcao := closure()
	funcao()
}
