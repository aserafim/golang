package main

import "fmt"

func main() {

	var variavelUm string = "Variável Um - Declaração explícita"
	variavelDois := "Variável Dois - Declaração implícita (inferência de tipo)"

	fmt.Println(variavelUm)
	fmt.Println(variavelDois)

	var (
		var3 string = "terceira"
		var4 string = "quarta"
	)

	fmt.Println(var3, var4)

	const const1 string = "Constante 1 - Valor imutável"

	fmt.Println(const1)

	var3, var4 = var4, var3

	fmt.Println(var3, var4)

}
