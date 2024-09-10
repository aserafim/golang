package main

import "fmt"

func main() {

	fmt.Println("Declaração e uso de variáveis")

	var num1 int = 3
	var str1 string = "Nada por aqui"
	var float1 float64 = 3.1415

	fmt.Println(num1, str1, float1)

	//Convertendo tipos
	num2 := int(float1)

	fmt.Printf("float1 convertido para int é: %d\n", num2)

	fmt.Printf("num1 convertido para float é: %f\n", float64(num1))
}
