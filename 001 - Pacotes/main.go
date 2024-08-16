package main

import (
	"fmt"
	"pacotes/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Testando funcionamento do módulo.")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("teste1@teste2.com")
	fmt.Println(erro)

}
