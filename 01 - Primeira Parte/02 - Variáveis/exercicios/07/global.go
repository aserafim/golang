/*
Declare uma variável global e crie uma
função que modifique o valor dessa variável.
Verifique como o valor global é alterado mesmo
fora do escopo da função.
*/

package main

import "fmt"

var global int = 9

func alteraValor(numero int) int {
	return numero - 1
}

func main() {

	fmt.Println("Variáveis Globais")

	global = alteraValor(global)

	fmt.Println(global)
}
