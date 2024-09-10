package main

import "fmt"

func main() {

	fmt.Println("Map(dicionário)")

	nomesIdades := map[string]int{
		"Alefe":     30,
		"Natalia":   27,
		"Jefferson": 34,
		"Talita":    33,
		"Manoel":    66,
	}

	nomesIdades["Solange"] = 59

	fmt.Println(nomesIdades)

	map2 := map[string]int{}
	map2["Qualquer coisa"] = 2

	fmt.Println(nomesIdades["Alefe"])

	map3 := map[int]int{}

}
