package main

import (
	"fmt"
)

func main() {

	//i := 0

	/* 	for i < 10 {
	   		i++
	   		fmt.Println("Incrementando i:", i)
	   		time.Sleep(time.Second)
	   	}

	   	for j := 0; j < 10; j++ {
	   		fmt.Println("Incrementando j", j+1)
	   		time.Sleep(time.Second)
	   	} */
	nomes := [3]string{"João", "Pedro", "Priscila"}

	//imprimindo os indice e os valores
	for ind, nome := range nomes {
		fmt.Println(ind, nome)
	}

	//imprimindo apenas o conteúdo
	for _, nome := range nomes {
		fmt.Println(nome)
	}

	//iterando por um string
	for indice, letra := range "Palavra" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string{
		"nome":      "Paulo",
		"sobrenome": "Pereira",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

}
