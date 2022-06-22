// - Crie um map com key tipo string e value tipo []string.
//     - Key deve conter nomes no formato sobrenome_nome
//     - Value deve conter os hobbies favoritos da pessoa
// - Demonstre todos esses valores e seus índices.

package main

import "fmt"

func main() {
	x := map[string][]string{
		"abner_sousa": {
			"irado", "teimoso",
		},
		"alfred": {
			"differs", "pokemo",
		},
		"neto": {
			"lesado", "dormi muito",
		},
	}
	for i, va := range x {
		fmt.Println(i)
		for j, v := range va {
			fmt.Println("\t", j, " - ", v)
		}
	}
}