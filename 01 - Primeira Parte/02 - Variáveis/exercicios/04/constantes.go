/*
Defina uma constante com o valor de PI e use-a
para calcular a área de um círculo com raio
fornecido pelo usuário.
*/
package main

import "fmt"

const pi float32 = 3.14

func main() {

	fmt.Println("Variáveis Constantes")
	var raio float32
	fmt.Print("Entre com o valor do raio do círculo: ")
	fmt.Scan(&raio)

	fmt.Println("A área do círculo é: ", pi*(raio*raio))

}
