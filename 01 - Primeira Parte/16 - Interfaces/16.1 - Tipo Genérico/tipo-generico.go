package main

import "fmt"

func generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {

	//Note que a nossa funcao aceita qualquer
	//tipo

	generica("String")
	generica(1)
	generica(true)

	//Um exemplo disso é a funcao Println,
	//que aceita qualquer parâmetro

	//Note ainda que abusar disso
	//pode ser perigoso

	mapa := map[interface{}]interface{}{
		1:            "String",
		float32(100): true,
	}

	fmt.Println(mapa)

}
