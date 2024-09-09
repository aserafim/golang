/*
Dentro de um loop for, declare uma variável local.
Imprima o valor da variável a cada iteração e veja
como ela é redefinida em cada ciclo do loop.
*/

package main

import "fmt"

func main() {

	fmt.Println("Variáveis de Bloco")

	i := 0
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}
}
