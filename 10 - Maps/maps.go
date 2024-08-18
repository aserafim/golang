package main

import "fmt"

func main() {

	fmt.Println("Maps")

	usuario := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	//o acesso aos dados do mapa
	//é diferente da struct
	fmt.Println(usuario["nome"])

	//Maps aninhados
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Arnoldo",
			"segundo":  "Thiago",
		},
	}

	fmt.Println(usuario2)

	//Deletar chaves
	delete(usuario2, "nome")
	fmt.Println(usuario2)
}
