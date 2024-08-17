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

	//Criando um array sem declara
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
}
