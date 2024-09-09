package main

import "fmt"

/*
Crie uma função que declara uma variável e
tente acessar essa variável fora da função.
Veja o erro gerado e refatore o código para corrigir o problema.
*/

func variaveis() {
	a := 3
	fmt.Println(a)
}

func variaveisDois() {
	//fmt.Println(a)
}

func main() {

	fmt.Println("Escopo de variáveis")

	//fmt.Println(a)
}
