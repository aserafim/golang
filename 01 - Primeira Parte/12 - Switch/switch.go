package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-Feira"
	case 3:
		return "Terça-Feira"
	case 4:
		return "Quarta-Feira"
	case 5:
		return "Quinta-Feira"
	case 6:
		return "Sexta-Feira"
	case 7:
		return "Sábado"
	default:
		return "Número inválido"
	}

}

func diaDaSemanaAlt(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-Feira"
	default:
		return "Número inválido"
	}
}

func main() {

	fmt.Println("Switch")

	dia := diaDaSemana(1)
	fmt.Println(dia)

	dia2 := diaDaSemanaAlt(2)
	fmt.Println(dia2)
}
