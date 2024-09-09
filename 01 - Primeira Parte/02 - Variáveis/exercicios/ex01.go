package main

import "fmt"

/*
01 - Declare três variáveis usando as diferentes formas
de declaração que Go oferece (var, atribuição curta :=,
e múltipla declaração). Exiba os valores das variáveis.
*/

// a := 1; não é possível declaração desse tipo aqui
var inteiro, terceiroInteiro int //declarando var de escopo global, sem atribuir valor
var segundoInteiro int = 35

func main() {

	fmt.Println(inteiro)
	fmt.Println(terceiroInteiro)
	fmt.Println(segundoInteiro)

	comValor := 34.9

	fmt.Println(comValor)

	texto := "Texto"
	fmt.Println(texto)

}
