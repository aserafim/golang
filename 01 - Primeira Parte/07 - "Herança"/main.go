package main

import "fmt"

// Dessa forma, estudante está
// basicamente tendo uma cópia
// dos campos de pessoa dentro dele
type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	cpf       string
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {

	fmt.Println("Herança")

	p1 := pessoa{"Jeffeson", "Serafim", 34, "33445688891"}
	fmt.Println(p1)

	estudante1 := estudante{p1, "Análise de Dados", "PUC"}

	fmt.Println(estudante1.nome, estudante1.sobrenome, estudante1.faculdade)

}
