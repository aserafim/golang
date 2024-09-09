/*
Crie uma variável que armazene a idade de uma pessoa.
Em seguida, altere o valor dessa variável para simular
que a pessoa envelheceu um ano. Exiba o valor da variável
antes e depois da reatribuição.
*/

package main

import "fmt"

func incrementIdade(idade int) int {
	return idade + 1
}

func main() {
	var idade int

	fmt.Println("Reatribuição de Valores")

	fmt.Print("Entre com a sua idade: ")
	fmt.Scan(&idade)

	fmt.Printf("Sua idade atual é: %d. No seu próximo aniversário sua idade será: %d.\n", idade, incrementIdade(idade))
}
