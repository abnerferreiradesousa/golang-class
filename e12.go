// - Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map,
// utilizando os sobrenomes como key.
// - Demonstre os valores do map utilizando range.
// - Os diferentes sabores devem ser demonstrados utilizando outro range, dentro
// do range anterior.

package main

import (
	"fmt"
)

type pessoa struct {
	nome      string
	sobrenome string
	sabores   []string
}

func main() {

	mymap := make(map[string]pessoa)

	mymap["abn"] = pessoa{
		nome:      "Renata",
		sobrenome: "Pimentão",
		sabores:   []string{"pistache", "morango", "baunilha"}}

	mymap["trat"] = pessoa{"Frederico", "da Prússia",
		[]string{"sabão em pó", "pé de coelho", "feijão"}}

	for _, valor := range mymap {
		fmt.Println("Meu nome é", valor.nome, "e o sabor de sorvete que eu gosto são")
		for _, valor := range valor.sabores {
			fmt.Println(" - ", valor)
		}
	}

}