package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

// o número de parâmtro variático é limitado a um
// e ele deve ser o último parâmetro
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {

	result := soma(1, 2, 5, 4, 8, 69)
	fmt.Println(result)

	escrever("Índice de número: ", 1, 2, 3, 4, 5, 6)
}
