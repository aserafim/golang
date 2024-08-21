package main

import "fmt"

func inverteSinal(n1 int) int {
	return n1 * -1
}

func inverteSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	fmt.Println("Ponteiros em funções")

	//Note que a função inverteSinal não está alterando
	//o valor armazenado em numero, mas sim uma cópia dele

	numero := 20

	novoNumero := inverteSinal(numero)
	fmt.Println(novoNumero)
	fmt.Println(numero)

	numeroDois := -5
	fmt.Println("Antes da função: ", numeroDois)
	inverteSinalComPonteiro(&numeroDois)
	fmt.Println("Após a função: ", numeroDois)

}
