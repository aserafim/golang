package main

import (
	"errors"
	"fmt"
)

func main() {

	//NÚMEROS INTEIROS

	fmt.Println("\nExistem 4 tipos de inteiros, cada um armazena uma detrminada quantidade de bits.")
	var (
		num1 int8
		num2 int16
		num3 int32
		num4 int64
	)

	fmt.Println(num1, num2, num3, num4)

	fmt.Println("A declaração usando apenas 'int' vai criar um inteiro com o número de bits da sua arquitetura (32 ou 64 bits)")

	var num5 uint = 100

	fmt.Println("Temos ainda os unsygned int, que são inteiros positivos 'sem sinal'. Eles também são divididos por qtd. de bits.", num5)

	fmt.Println("\nTemos ainda alguns alias, 'rune' para int32 e 'byte' para uint8.")

	var oito byte = 4
	var trintaedois rune = 36

	fmt.Println("\n", oito, trintaedois)

	//FIM NÚMEROS INTEIROS

	//NÚMEROS REAIS

	var real1 float32 = 12.4
	var real2 float64 = 45.8

	fmt.Println(real1, real2)

	//FIM NÚMEROS REAIS

	//STRINGS
	//Não temos char em GoLang
	//o char na verdade retorna o número
	//do índice do caractere na tabela ASCII
	char := 'A'
	fmt.Println(char)

	//ERROR
	var falha error = errors.New("Erro de execução.")
	fmt.Println(falha)

}
