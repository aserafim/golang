package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Alterações salvas. Obrigado, %s.\n", u.nome)
}

func (u usuario) verificaMaioridade() {
	if u.idade >= 18 {
		fmt.Printf("%s é maior de idade.\n", u.nome)
	} else {
		fmt.Printf("É necessário ser maior de idade para fazer isso.\n")
	}

}

func (u *usuario) fazerAniversario() {
	u.idade++
}

func main() {

	fmt.Println("Métodos")

	userUm := usuario{"Alefe", 30}
	userUm.salvar()

	userDois := usuario{"Pedro", 12}
	userDois.salvar()

	userDois.verificaMaioridade()
	userUm.verificaMaioridade()

	fmt.Println(userDois.idade)
	userDois.fazerAniversario()
	fmt.Println(userDois.idade)
}
