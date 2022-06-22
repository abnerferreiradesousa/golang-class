// - Crie uma slice contendo slices de strings ([][]string). Atribua valores a este
// slice multi-dimensional da seguinte maneira:
//     - "Nome", "Sobrenome", "Hobby favorito"
// - Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.

package main

import "fmt"

func main() {

	x := [][]string{
		{"Nome", "Sobrenome", "Hobby favorito"},
		{"Abner", "Smith", "ouvir musica"},
		{"Gustavo", "Wender", "jogar videogame"},
		{"Álife", "Sousa", "passar tempo com os irmãos"},
	}

	for _, v := range x {
		fmt.Println(v[0])
		for j := 0; j < len(v); j++ {
			fmt.Println("\t", v[j])
		}

	}

}