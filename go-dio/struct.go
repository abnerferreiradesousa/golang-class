package main

import "fmt"

type Pessoa struct {
	nome  string
	idade int
}

func main() {

	x := Pessoa{"Ana", 54}
	fmt.Printf(x.nome)
	fmt.Println(Pessoa{"Ana", 54})

}
