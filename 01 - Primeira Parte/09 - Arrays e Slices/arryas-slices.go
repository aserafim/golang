package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	fmt.Println(array1)

	array1[0] = 10
	fmt.Println(array1)

	array2 := [2]string{"Talita", "Felip"}
	fmt.Println(array2)

	//Criando um array sem declarar
	//o tamanho
	arr3 := [...]int{1, 2, 3, 4}
	fmt.Println(arr3)

	//o tamanho do slice é dinâmico
	slc := []int{1, 2, 3}
	fmt.Println(slc)
	fmt.Println(slc)

	fmt.Println(reflect.TypeOf(slc))
	fmt.Println(reflect.TypeOf(arr3))

	slc = append(slc, 5)
	fmt.Println(slc)

	//os elementos dele funcionam como
	//ponteiros para o array de origem
	slc2 := arr3[1:3]
	fmt.Println(slc2)

	//Arrays internos
	//a função make
	//cria um array e
	//retorna pra nós apenas a fatia(slice)
	//que criamos
	fmt.Println("------------------")
	slc3 := make([]float32, 10, 11)
	fmt.Println(slc3)
	fmt.Println(len(slc3))
	fmt.Println(cap(slc3))
	slc3 = append(slc3, 18)
	slc3 = append(slc3, 8)
	fmt.Println(cap(slc3))
}
