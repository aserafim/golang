package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {

	fmt.Println("Tipos compostos(array e slice)")

	var vetor1 [10]int
	for indice := range vetor1 {
		vetor1[indice] = rand.IntN(100)
	}

	fmt.Println(vetor1)

	var slc1 []int = vetor1[2:5]
	for indice2 := range slc1 {
		slc1[indice2] = rand.IntN(5)
	}

	//	fmt.Println(slc1)

	slc1 = append(slc1, 5)
	fmt.Println(slc1)

	//arr3 := [...]int{1, 2, 3, 4}

}
