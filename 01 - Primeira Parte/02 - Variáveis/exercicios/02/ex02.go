package main

import "fmt"

/*
Declare variáveis de diferentes tipos
(int, float, string, bool)
sem inicializá-las explicitamente.
Exiba os valores para ver o comportamento
padrão de Go em relação a valores zero.
*/

var novoInteiro int
var real float32
var texto string
var verdadeiro bool

func main() {

	fmt.Println(novoInteiro)
	fmt.Print(real)
	fmt.Println(texto)
	fmt.Println(verdadeiro)

}
