package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

// Funções em GoLang podem ter múltiplos
// retornos
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {

	soma := somar(10, 20)

	fmt.Println(soma)

	//Como função é um tipo
	//podemos retornar uma função
	//e guardar numa var

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Texto da função 1")

	res1, res2 := calculosMatematicos(3, 4)
	fmt.Println(res1, res2)

	//Se quisermos ignorar algum dos retornos
	//basta utilizar um _ (underline ou underscore)
	res3, _ := calculosMatematicos(4, 9)
	fmt.Println(res3)

}
