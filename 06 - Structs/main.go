/*
Structs é o mais próximo que temos de uma classe

	em Go sua definição é: Uma coleção de campos
*/
package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int8
}

func main() {
	fmt.Println("Structs")

	var u1 usuario
	u1.idade = 10
	u1.nome = "Lucas"

	fmt.Println(u1)
	end1 := endereco{"Rua dos Bobos", 0}

	u2 := usuario{"Pedro", 12, end1}
	fmt.Println(u2)

	u3 := usuario{nome: "Solange", endereco: end1}
	fmt.Println(u3)
}
