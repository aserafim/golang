package main

import "fmt"

func main() {
	var var1 int = 10

	//Ao fazer isso o Go
	//cria uma copia do valor de var1
	//e atribui a var2
	var var2 int = var1

	fmt.Println(var1, var2)

	//Por isso quando incrementamos
	//var1, somente o seu valor
	//é alterado
	var1++
	fmt.Println(var1, var2)

	//PONTEIRO É UMA REFERÊNCIA DE MEMÓRIA
	//ELE INDICA UM ENDEREÇO DE MEMÓRIA
	var var3 int
	var ponteiro *int

	var3 = 100
	ponteiro = &var3

	fmt.Println(var3, ponteiro)
	fmt.Println(var3, *ponteiro)

	var3 = 15

	fmt.Println(var3, ponteiro)
	fmt.Println(var3, *ponteiro)

}
