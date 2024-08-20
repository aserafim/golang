package main

import "fmt"

func funcaoUm() {
	fmt.Println("Executando a função 1...")
}

func funcaoDois() {
	fmt.Println("Executando a função 2...")
}

func alunoEstaAprovado(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado.")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}
	return false
}

func main() {

	fmt.Println("Defer")

	//A clausula defer adia a execução
	//de uma função ou statement
	defer funcaoUm()
	funcaoDois()

	//Note que a mensagem média calculada
	//foi impressa antes do retorno da
	//função alunoEstaAprovado
	alunoEstaAprovado(7, 5)
}
