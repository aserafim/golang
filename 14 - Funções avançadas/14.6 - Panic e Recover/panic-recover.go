package main

import "fmt"

func recuperaExecucao() {
	fmt.Println("Tentando recuperar o programa...")
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func verificaAprovacao(n1, n2 float32) bool {
	defer recuperaExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}
	//antes de panic ser executado
	//ele executa todas as funções
	//que tem a cláusula
	//defer
	panic("A MÉDIA É 6!")
}

func main() {

	fmt.Println(verificaAprovacao(6, 6))
	fmt.Println("Pós execução...")

}
